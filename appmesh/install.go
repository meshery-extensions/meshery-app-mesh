package appmesh

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"

	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

const (
	repo              = "https://aws.github.io/eks-charts"
	appMeshController = "appmesh-controller"
	appMeshInject     = "appmesh-inject"
	appMeshGateway    = "appmesh-gateway"
)

// Installs APP-MESH service mesh using helm charts.
func (appMesh *AppMesh) installAppMesh(del bool, version, namespace string, kubeconfigs []string) (string, error) {
	appMesh.Log.Debug(fmt.Sprintf("Requested install of version: %s", version))
	appMesh.Log.Debug(fmt.Sprintf("Requested action is delete: %v", del))
	appMesh.Log.Debug(fmt.Sprintf("Requested action is in namespace: %s", namespace))

	appMesh.Log.Info(fmt.Sprintf("Requested install of version: %s", version))
	st := status.Installing
	if del {
		st = status.Removing
	}

	err := appMesh.Config.GetObject(adapter.MeshSpecKey, appMesh)
	if err != nil {
		return st, ErrMeshConfig(err)
	}

	err = appMesh.applyHelmChart(del, version, namespace, kubeconfigs)
	if err != nil {
		appMesh.Log.Error(ErrInstallAppMesh(err))
		return st, ErrInstallAppMesh(err)
	}

	if del {
		return status.Removed, nil
	}
	return status.Installed, nil
}

func (appMesh *AppMesh) applyHelmChart(del bool, version, namespace string, kubeconfigs []string) error {
	version = strings.TrimPrefix(version, "v")
	appMesh.Log.Info("Installing using helm charts...")
	var act mesherykube.HelmChartAction
	if del {
		act = mesherykube.UNINSTALL
	} else {
		act = mesherykube.INSTALL
	}
	cv, err := mesherykube.HelmAppVersionToChartVersion(repo, appMeshController, version)
	if err != nil {
		return ErrApplyHelmChart(err)
	}

	var wg sync.WaitGroup
	var errs []error
	var errMx sync.Mutex

	for _, config := range kubeconfigs {
		wg.Add(1)
		go func(config string, act mesherykube.HelmChartAction) {
			defer wg.Done()
			kClient, err := mesherykube.New([]byte(config))
			if err != nil {
				errMx.Lock()
				errs = append(errs, err)
				errMx.Unlock()
				return
			}

			// Install the controller
			err = kClient.ApplyHelmChart(mesherykube.ApplyHelmChartConfig{
				ChartLocation: mesherykube.HelmChartLocation{
					Repository: repo,
					Chart:      appMeshController,
					Version:    cv,
				},
				Namespace:       namespace,
				Action:          act,
				CreateNamespace: true,
			})
			if err != nil {
				errMx.Lock()
				errs = append(errs, err)
				errMx.Unlock()
				return
			}

			// Install appmesh-injector. Only needed for controller versions older
			// than 1.0.0
			if controlPlaneVersion, err := strconv.Atoi(strings.TrimPrefix(version, "v")[:1]); controlPlaneVersion < 1 && err != nil {
				err = kClient.ApplyHelmChart(mesherykube.ApplyHelmChartConfig{
					ChartLocation: mesherykube.HelmChartLocation{
						Repository: repo,
						Chart:      appMeshInject,
						AppVersion: "", //defaults to latest
					},
					Namespace:       namespace,
					Action:          act,
					CreateNamespace: true,
				})
				if err != nil {
					errMx.Lock()
					errs = append(errs, err)
					errMx.Unlock()
					return
				}
			}
		}(config, act)
	}

	wg.Wait()
	if len(errs) == 0 {
		return nil
	}

	mergedErrors := mergeErrors(errs)
	return ErrApplyHelmChart(mergedErrors)
}

func (appMesh *AppMesh) applyManifest(manifest []byte, isDel bool, namespace string, kubeconfigs []string) error {
	var wg sync.WaitGroup
	var errs []error
	var errMx sync.Mutex

	for _, config := range kubeconfigs {
		wg.Add(1)
		go func(config string) {
			defer wg.Done()
			kClient, err := mesherykube.New([]byte(config))
			if err != nil {
				errMx.Lock()
				errs = append(errs, err)
				errMx.Unlock()
				return
			}

			err = kClient.ApplyManifest(manifest, mesherykube.ApplyOptions{
				Namespace: namespace,
				Update:    true,
				Delete:    isDel,
			})
			if err != nil {
				errMx.Lock()
				errs = append(errs, err)
				errMx.Unlock()
				return
			}
		}(config)
	}

	wg.Wait()
	if len(errs) == 0 {
		return nil
	}

	return mergeErrors(errs)
}
