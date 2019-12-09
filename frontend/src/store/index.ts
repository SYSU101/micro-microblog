import Vue from 'vue';
import Vuex from 'vuex';
import { HTTP_CLIENT } from '@/utils/HttpClient';
import { IUserProfile } from '@/typings';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    userProfile: null as (IUserProfile | null),
  },
  mutations: {
    setUserProfile(state, payload) {
      state.userProfile = payload;
    },
    clearUserProfile(state) {
      state.userProfile = null;
    },
  },
  getters: {
    getUserProfile(state) {
      return () => state.userProfile;
    },
    getUserId(state) {
      return () => {
        const userProfile = state.userProfile;
        if (userProfile) {
          return userProfile.id;
        } else {
          return null;
        }
      };
    },
  },
  actions: {
    async fetchUserProfile({ commit }, userId) {
      const userProfile = await HTTP_CLIENT.get<IUserProfile>(`/user/${userId}`);
      commit('setUserProfile', userProfile);
    },
  },
});
