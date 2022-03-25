import type { Plugin } from "vue";
class Api {
  async post(url: string, data?: any) {
    try {
      const response = await fetch(url, { method: "POST", body: data });
      return { success: true, data: await response.json() };
    } catch (err) {
      console.error(err);
      return { success: false, data: null };
    }
  }
}

const $api = new Api();

const plugin: Plugin = {
  install(app) {
    app.config.globalProperties.$api = $api;
  },
};

export default plugin;
export { $api };
