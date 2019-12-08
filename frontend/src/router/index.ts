import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import store from '@/store';
import { SILENT_HTTP_CLIENT } from '@/utils/HttpClient';

Vue.use(VueRouter);

const routes: RouteConfig[] = [
  {
    path: '/login',
    component: () => import('../views/Login.vue'),
  },
  {
    path: '/register',
    component: () => import('../views/Register.vue'),
  },
  {
    path: '/',
    component: () => import('../views/BasicLayout.vue'),
    async beforeEnter(from, to, next) {
      try {
        const { userId } = await SILENT_HTTP_CLIENT.get<{ userId: number }>('/sessions');
        await store.dispatch('fetchUserProfile', userId);
        next();
      } catch (error) {
        next('/login');
      }
    },
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
