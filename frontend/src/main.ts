import { createApp } from "vue";
import "./assets/styles/style.less";
import App from "./App.vue";
import { setupRouter } from "./routes";
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';

function init() {
  const app = createApp(App);
  setupRouter(app);
  app.use(ArcoVue);
  app.mount("#app");
};

void init();
