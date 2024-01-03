package mycontroller

import (
	"context"

	v1 "github.com/orangedeng/test-extension-operator/pkg/apis/test.pandaria.io/v1"
	resources "github.com/orangedeng/test-extension-operator/pkg/client/generated/controllers/test.pandaria.io"
	v1core "github.com/rancher/wrangler/v2/pkg/generated/controllers/core/v1"
	"github.com/sirupsen/logrus"
	k8scorev1 "k8s.io/api/core/v1"
)

type handler struct {
	ctx        context.Context
	testClient resources.Interface
}

func Register(ctx context.Context, factory *resources.Factory, nsController v1core.NamespaceController) {
	h := &handler{
		ctx:        ctx,
		testClient: factory.Test(),
	}
	fooController := factory.Test().V1().Foo()
	fooController.OnChange(ctx, "foo-change", h.onFooChange)
	nsController.OnChange(ctx, "test-ns-change", h.onNamespaceChange)
}

func (h *handler) onFooChange(key string, obj *v1.Foo) (*v1.Foo, error) {
	logrus.Infof("get change from key %s", key)
	return obj, nil
}

func (h *handler) onNamespaceChange(key string, obj *k8scorev1.Namespace) (*k8scorev1.Namespace, error) {
	logrus.Infof("get change from ns key %s", key)
	return obj, nil
}
