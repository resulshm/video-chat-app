class Api {
  async post(url: string, data?: any) {
    try {
      const response = await fetch(url, { method: "POST", body: data });
      return { success: true as const, data: (await response.json()).data };
    } catch (err) {
      console.error(err);
      return { success: false, data: null };
    }
  }
}

const api = new Api();

export { api };
