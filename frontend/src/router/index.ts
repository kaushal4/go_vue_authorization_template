import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import DashBoard from "../views/DashBoard.vue";
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/login",
    redirect: "/",
  },
  {
    path: "/login/:id",
    name: "login",
    component: Login,
    props: true,
  },
  {
    path: "/student",
    name: "student",
    component: DashBoard,
    props: {
      type: "student",
    },
  },
  {
    path: "/teacher",
    name: "teacher",
    component: DashBoard,
    props: {
      type: "teacher",
    },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
