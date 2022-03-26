import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/index.vue";
import RoomView from "@/views/room/index.vue";
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "HomeView",
      component: HomeView,
    },
    {
      path: "/:id",
      name: "RoomView",
      component: RoomView,
    },
  ],
});

export default router;
