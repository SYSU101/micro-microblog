<template>
  <div id="login">
    <LoginForm></LoginForm>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import LoginForm from '@/components/LoginForm.vue';

@Component({
  components: {
    LoginForm,
  },
})
export default class Login extends Vue { }

</script>

<style lang="less" scoped>
#login {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 100vw;
  height: 100vh;
  min-height: 500px;
  min-width: 500px;
  background-color: #dddddd;
}
</style>
