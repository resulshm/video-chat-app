import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import api from "@/plugins/api";
import "@/assets/css/tailwind.css";
const app = createApp(App);

app.use(router);
app.use(api);
app.mount("#app");
