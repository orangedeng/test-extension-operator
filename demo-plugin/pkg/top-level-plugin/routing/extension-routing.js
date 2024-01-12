// Don't forget to create a VueJS page called index.vue in the /pages folder!!!
import Dashboard from '../pages/index.vue';

const BLANK_CLUSTER = '_';
const YOUR_PRODUCT_NAME = 'top-level-plugin';

const routes = [
  {
    name:      `${ YOUR_PRODUCT_NAME }-c-cluster-resource`,
    path:      `/${ YOUR_PRODUCT_NAME }/c/:cluster/:resource`,
    component: Dashboard,
    meta:      {
      product: YOUR_PRODUCT_NAME,
      cluster: BLANK_CLUSTER
    },
  }
];

export default routes;
