package controllers

import (
	"context"
	"fmt"

	"github.com/orangedeng/test-extension-operator/internal/controllers/mycontroller"
	resources "github.com/orangedeng/test-extension-operator/pkg/client/generated/controllers/test.pandaria.io"

	v1core "github.com/rancher/wrangler/v2/pkg/generated/controllers/core"
	"github.com/rancher/wrangler/v2/pkg/start"
	"k8s.io/client-go/rest"
)

func Start(ctx context.Context, kubeconfig *rest.Config) error {
	tests, err := resources.NewFactoryFromConfig(kubeconfig)
	if err != nil {
		return fmt.Errorf("Error building test controllers: %s", err.Error())
	}
	core, err := v1core.NewFactoryFromConfig(kubeconfig)
	if err != nil {
		return fmt.Errorf("Error building core sample controllers: %s", err.Error())
	}
	mycontroller.Register(ctx, tests, core.Core().V1().Namespace())

	if err := start.All(ctx, 2, tests, core); err != nil {
		return fmt.Errorf("Error starting: %s", err.Error())
	}
	<-ctx.Done()
	return nil
}
