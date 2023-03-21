package appmesh

import (
	"github.com/layer5io/meshkit/errors"
)

var (
	// ErrCustomOperationCode should really have an error code defined by now.
	ErrCustomOperationCode = "appmesh_test_code"
	// ErrInstallAppMeshCode provisioning failure
	ErrInstallAppMeshCode = "appmesh_test_code"
	// ErrMeshConfigCode   service mesh configuration failure
	ErrMeshConfigCode = "appmesh_test_code"
	// ErrClientConfigCode adapter configuration failure
	ErrClientConfigCode = "appmesh_test_code"
	// ErrStreamEventCode  failure
	ErrStreamEventCode = "appmesh_test_code"
	// ErrSampleAppCode    failure
	ErrSampleAppCode = "appmesh_test_code"
	// ErrLoadNamespaceToMeshCode represents the error
	// which is generated when the namespace could not be labeled and updated
	ErrLoadNamespaceToMeshCode = "appmesh_test_code"
	// ErrOpInvalidCode failure
	ErrOpInvalidCode = "appmesh_test_code"
	// ErrNilClientCode represents the error code which is
	// generated when kubernetes client is nil
	ErrNilClientCode = "replace"
	// ErrApplyHelmChartCode represents the error generated
	// during the process of applying helm chart
	ErrApplyHelmChartCode = "replace"

	// ErrParseAppMeshCoreComponentCode represents the error code
	// when app-mesh core components can't be parsed
	ErrParseAppMeshCoreComponentCode = "replace"

	// ErrAppMeshCoreComponentFailCode represents error code when
	// there is an error parsing components
	ErrAppMeshCoreComponentFailCode = "replace"

	// ErrInvalidOAMComponentTypeCode represents error code when
	// invalid OAM components are registered
	ErrInvalidOAMComponentTypeCode = "replace"

	// ErrProcessOAMCode represents error code while parsing OAM
	// components
	ErrProcessOAMCode = "replace"

	// ErrAddonFromTemplateCode represents the errors which are generated
	// during addon deployment process
	ErrAddonFromTemplateCode = "replace"

	//ErrAddonFromHelmCode represents the error while installing addons through helm charts
	ErrAddonFromHelmCode = "replace"

	// ErrParseOAMComponentCode represents the error which is
	// generated during the OAM component parsing
	ErrParseOAMComponentCode = "replace"

	// ErrParseOAMConfigCode represents the error which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfigCode = "replace"
	// ErrOpInvalid is an error when an invalid operation is requested
	ErrOpInvalid = errors.New(ErrOpInvalidCode, errors.Alert, []string{"Invalid operation"}, []string{}, []string{}, []string{})

	// ErrNilClient represents the error generated when kubernetes client is nil
	ErrNilClient = errors.New(ErrNilClientCode, errors.Alert, []string{"kubernetes client not initialized"}, []string{"Kubernetes client is nil"}, []string{"kubernetes client not initialized"}, []string{"Reconnect the adaptor to Meshery server"})

	// ErrParseOAMComponent represents the error which is
	// generated during the OAM component parsing
	ErrParseOAMComponent = errors.New(ErrParseOAMComponentCode, errors.Alert, []string{"error parsing the component"}, []string{"Error occurred while prasing application component in the OAM request made"}, []string{"Invalid OAM component passed in OAM request"}, []string{"Check if your request has vaild OAM components"})

	// ErrParseOAMConfig represents the error which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfig = errors.New(ErrParseOAMConfigCode, errors.Alert, []string{"error parsing the configuration"}, []string{"Error occurred while prasing component config in the OAM request made"}, []string{"Invalid OAM config passed in OAM request"}, []string{"Check if your request has vaild OAM config"})
)

// ErrInstallAppMesh is the error for install mesh
func ErrInstallAppMesh(err error) error {
	return errors.New(ErrInstallAppMeshCode, errors.Alert, []string{"Error with App Mesh installation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrMeshConfig is the error for mesh config
func ErrMeshConfig(err error) error {
	return errors.New(ErrMeshConfigCode, errors.Alert, []string{"Error configuration mesh"}, []string{err.Error()}, []string{}, []string{})
}

// ErrClientConfig is the error for setting client config
func ErrClientConfig(err error) error {
	return errors.New(ErrClientConfigCode, errors.Alert, []string{"Error setting client config"}, []string{err.Error()}, []string{}, []string{})
}

// ErrStreamEvent is the error for streaming event
func ErrStreamEvent(err error) error {
	return errors.New(ErrStreamEventCode, errors.Alert, []string{"Error streaming events"}, []string{err.Error()}, []string{}, []string{})
}

// ErrSampleApp is the error for operations on the sample apps
func ErrSampleApp(err error, status string) error {
	return errors.New(ErrSampleAppCode, errors.Alert, []string{"Error with sample app operation"}, []string{err.Error(), "Error occurred while trying to install a sample application using manifests"}, []string{"Invalid kubeclient config", "Invalid manifest"}, []string{"Reconnect your adapter to meshery server to refresh the kubeclient"})
}

// ErrCustomOperation is the error for custom operations
func ErrCustomOperation(err error) error {
	return errors.New(ErrCustomOperationCode, errors.Alert, []string{"Error with applying custom operation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrApplyHelmChart is the occurend while applying helm chart
func ErrApplyHelmChart(err error) error {
	return errors.New(ErrApplyHelmChartCode, errors.Alert, []string{"Error occurred while applying Helm Chart"}, []string{err.Error()}, []string{}, []string{})
}

// ErrParseAppMeshCoreComponent is the error when app-mesh core component manifest parsing fails
func ErrParseAppMeshCoreComponent(err error) error {
	return errors.New(ErrParseAppMeshCoreComponentCode, errors.Alert, []string{"app-mesh core component manifest parsing failing"}, []string{err.Error()}, []string{}, []string{})
}

// ErrInvalidOAMComponentType is the error when the OAM component name is not valid
func ErrInvalidOAMComponentType(compName string) error {
	return errors.New(ErrInvalidOAMComponentTypeCode, errors.Alert, []string{"invalid OAM component name: ", compName}, []string{}, []string{}, []string{})
}

// ErrAppMeshCoreComponentFail is the error when core appmesh component processing fails
func ErrAppMeshCoreComponentFail(err error) error {
	return errors.New(ErrAppMeshCoreComponentFailCode, errors.Alert, []string{"error in app-mesh core component"}, []string{err.Error()}, []string{}, []string{})
}

// ErrProcessOAM is a generic error which is thrown when an OAM operations fails
func ErrProcessOAM(err error) error {
	return errors.New(ErrProcessOAMCode, errors.Alert, []string{"error performing OAM operations"}, []string{err.Error()}, []string{}, []string{})
}

// ErrLoadNamespaceToMesh identifies the inability to label the appropropriate namespace
func ErrLoadNamespaceToMesh(err error) error {
	return errors.New(ErrLoadNamespaceToMeshCode, errors.Alert, []string{"Could not label the appropriate namespace"}, []string{err.Error()}, []string{}, []string{})
}

// ErrAddonFromTemplate is the error for streaming event
func ErrAddonFromTemplate(err error) error {
	return errors.New(ErrAddonFromTemplateCode, errors.Alert, []string{"Error with addon install operation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrAddonFromHelm is the error for installing addons through helm chart
func ErrAddonFromHelm(err error) error {
	return errors.New(ErrAddonFromHelmCode, errors.Alert, []string{"Error with addon install operation by helm chart"}, []string{err.Error()}, []string{"The helm chart URL in additional properties of addon Operation might be incorrect", "The helm installation failed due to any other reason"}, []string{})
}
