<template>
  <div id="basic-layout">
    <a-layout style="height: 100vh; width: 100vw;">
      <a-layout-sider>
        <a-menu theme="dark" mode="inline" v-model="selectedKeys">
          <a-menu-item key="me">
            <router-link to="/me">
              <a-icon type="profile" />
              <span>Me</span>
            </router-link>
          </a-menu-item>
          <a-menu-item key="other">
            <router-link to="/other">
              <a-icon type="ordered-list" />
              <span>Other People</span>
            </router-link>
          </a-menu-item>
          <a-menu-item key="logout" @click.native="logout($event)" >
            <a-icon type="logout"/>
            <span>Logout</span>
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-content
          :style="{ margin: '24px 16px', padding: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
        >
          <router-view></router-view>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { SILENT_HTTP_CLIENT } from '../utils/HttpClient';

@Component({ })
export default class BasicLayout extends Vue {
  public selectedKeys: string[] = [];

  public mounted() {
    this.selectedKeys = [ this.$route.path.split('/').pop() || 'me' ];
  }

  public async logout(event: MouseEvent | { domEvent: MouseEvent }) {
    try {
      await SILENT_HTTP_CLIENT.delete<{}>('/sessions');
    } catch (error) {
      // Do nothing
    }
    if (event.hasOwnProperty('domEvent')) {
      this.$store.commit('clearUserProfile');
      this.$notification.success({ message: '登出成功', description: '您现在可以重新登陆' });
      this.$router.push('/login');
    }
  }
}

</script>

<style lang="less" scoped>
#basic-layout {
  width: 100vw;
  height: 100vh;
  background-color: #dddddd;
}
</style>
