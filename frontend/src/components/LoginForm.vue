<template>
  <a-form
    id="login-form"
    :form="form"
    class="login-form"
    @submit="login"
  >
    <span class="form-title">LOGIN ACCOUNT</span>
    <a-form-item>
      <a-input
        v-decorator="[
          'username',
          { rules: [{ required: true, message: 'Please input your username!' }] },
        ]"
        placeholder="Username"
      >
        <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)" />
      </a-input>
    </a-form-item>
    <a-form-item>
      <a-input
        v-decorator="[
          'password',
          { rules: [{ required: true, message: 'Please input your Password!' }] },
        ]"
        type="password"
        placeholder="Password"
      >
        <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)" />
      </a-input>
    </a-form-item>
    <a-form-item>
      <a-button type="primary" html-type="submit" class="login-button">
        Log in
      </a-button>
      Or
      <a>
        <router-link to="/register">register now!</router-link>
      </a>
    </a-form-item>
  </a-form>
</template>

<script lang='ts'>
import { Vue, Component } from 'vue-property-decorator';
import { Form } from 'ant-design-vue';
import { SILENT_HTTP_CLIENT } from '../utils/HttpClient';

@Component({})
export default class LoginForm extends Vue {
  public form: any;

  public beforeCreate() {
    this.form = this.$form.createForm(this);
  }

  public login(event: Event) {
    event.preventDefault();
    this.form.validateFields(async (err: any, values: any) => {
      if (!err) {
        try {
          await SILENT_HTTP_CLIENT.put<{}>('/sessions', values);
          this.$notification.success({ message: '登录成功', description: '正在跳转' });
          this.$router.push('/');
        } catch (error) {
          this.$notification.error({ message: '登录失败', description: '用户名或密码错误'});
        }
      }
    });
  }
}
</script>
<style lang="less" scoped>
#login-form {
  width: 360px;
  padding: 0 48px;
  background-color: white;
  border-radius: 10px 10px #333 10px;

  .form-title {
    color: #1890ff;
    width: 100%;
    text-align: center;
    font-size: 18px;
    letter-spacing: 0.5em;
    font-weight: bold;
    display: inline-block;
    margin-bottom: 24px;
    margin-top: 48px;
  }
  .login-button {
    width: 100%;
  }
}
</style>