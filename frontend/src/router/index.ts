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
    children: [
      {
        path: '/me',
        component: () => import('../views/UserProfile.vue'),
      },
      {
        path: '/other',
        component: () => import('../views/UserList.vue'),
      },
      {
        path: '*',
        redirect: '/me',
      },
    ],
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
  mode: 'hash',
  base: process.env.BASE_URL,
  routes,
});

export default router;
