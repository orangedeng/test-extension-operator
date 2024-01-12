import { IPlugin } from '@shell/core/types';

export function init($plugin: IPlugin, store: any) {
  const YOUR_PRODUCT_NAME = 'top-level-plugin';
  const BLANK_CLUSTER = '_';
  const YOUR_K8S_RESOURCE_NAME = 'test.pandaria.io.foo';

  const { product, configureType, basicType } = $plugin.DSL(store, YOUR_PRODUCT_NAME);

  // registering a top-level product
  product({
    icon:    'gear',
    inStore: 'management',
    weight:  100,
    to:      {
      name:   `${ YOUR_PRODUCT_NAME }-c-cluster-resource`,
      params: {
        product:  YOUR_PRODUCT_NAME,
        cluster:  BLANK_CLUSTER,
        resource: YOUR_K8S_RESOURCE_NAME
      }
    }
  });

  // defining a k8s resource as page
  configureType(YOUR_K8S_RESOURCE_NAME, {
    displayName: 'some-custom-name-you-wish-to-assign-to-this-resource',
    isCreatable: true,
    isEditable:  true,
    isRemovable: true,
    showAge:     true,
    showState:   true,
    canYaml:     true,
    customRoute: {
      name:   `${ YOUR_PRODUCT_NAME }-c-cluster-resource`,
      params: {
        product:  YOUR_PRODUCT_NAME,
        cluster:  BLANK_CLUSTER,
        resource: YOUR_K8S_RESOURCE_NAME
      }
    }
  });

  basicType([YOUR_K8S_RESOURCE_NAME]);
}
