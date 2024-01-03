package main

import (
	"os"

	"github.com/orangedeng/test-extension-operator/internal/codegen/crds"
	v1 "github.com/orangedeng/test-extension-operator/pkg/apis/test.pandaria.io/v1"
	controllergen "github.com/rancher/wrangler/v2/pkg/controller-gen"
	"github.com/rancher/wrangler/v2/pkg/controller-gen/args"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/orangedeng/test-extension-operator/pkg/client/generated",
		Boilerplate:   "scripts/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"test.pandaria.io": {
				Types: []interface{}{
					v1.Foo{},
				},
				GenerateTypes: true,
			},
		},
	})
	err := crds.WriteCRD()
	if err != nil {
		panic(err)
	}
}
