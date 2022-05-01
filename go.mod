module github.com/layer5io/meshery-app-mesh

go 1.16

replace github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200723152044-916f10574334

require (
	github.com/layer5io/meshery-adapter-library v0.5.3
	github.com/layer5io/meshkit v0.5.8
	github.com/layer5io/service-mesh-performance v0.3.4
	google.golang.org/genproto v0.0.0-20210903162649-d08c68adba83 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.23.6
)
