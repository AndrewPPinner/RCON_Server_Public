import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Login from "../views/LoginView.vue";
import Dashboard from "../views/DashboardView.vue";
import store from "@/store";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    component: Dashboard,
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/login",
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach(async (to, from) => {
  if (!store.state.isAuthorized && to.name !== "login") {
    return { name: "login" };
  }
});

export default router;
