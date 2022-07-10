import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
const app = createApp(App);

const components = import.meta.globEager("./components/base/*.vue");

Object.entries(components).forEach(([path, definition]) => {
  const componentName = path
    .split("/")
    .pop()
    .replace(/\.\w+$/, "");
  app.component(componentName, definition.default);
});

app.use(router);
app.mount("#app");
