<template>
  <a-form id="register-form" :form="form" @submit="register">
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Username">
      <a-input
        v-decorator="[
          'username',
          { rules: [{ required: true, message: 'Please input your username!', whitespace: true }] },
        ]"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Password">
      <a-input
        v-decorator="[
          'password',
          {
            rules: [
              {
                required: true,
                message: 'Please input your password!',
              },
              {
                validator: validateToNextPassword,
              },
            ],
          },
        ]"
        type="password"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Confirm Password">
      <a-input
        v-decorator="[
          'confirm',
          {
            rules: [
              {
                required: true,
                message: 'Please confirm your password!',
              },
              {
                validator: compareToFirstPassword,
              },
            ],
          },
        ]"
        type="password"
        @blur="handleConfirmBlur"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Real Name">
      <a-input
        v-decorator="[
          'name',
          { rules: [{ required: true, message: 'Please input your real name!' }] },
        ]"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Student Id">
      <a-input-number
        v-decorator="[
          'studentId',
          { rules: [{ required: true, message: 'Please input your student ID!' }] }
        ]"
        style="width: 100%"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout" label="Birthday">
      <a-date-picker v-decorator="[
        'birthday',
        { rules: [{ type: 'object', required: true, message: 'Please select your birthday!' }] }]"
      />
    </a-form-item>
    <a-form-item :label-col="labelColLayout" :wrapper-col="inputColLayout">
      <span slot="label">
        Motto&nbsp;
        <a-tooltip title="You can leave it empty and edit it whenever you like">
          <a-icon type="question-circle-o" />
        </a-tooltip>
      </span>
      <a-input v-decorator="['motto']" />
    </a-form-item>
    <a-form-item>
      <a-button class="submit-button" type="primary" html-type="submit">
        Register
      </a-button>
    </a-form-item>
  </a-form>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { HTTP_CLIENT } from '../utils/HttpClient';

@Component({})
export default class RegisterForm extends Vue {
  public labelColLayout = { span: 8 };
  public inputColLayout = { span: 16 };
  public form: any;
  public confirmDirty = false;

  public beforeCreate() {
    this.form = this.$form.createForm(this);
  }

  public validateToNextPassword(rule: any, value: string, callback: (errMsg?: string) => void) {
    if (value && this.confirmDirty) {
      this.form.validateFields(['confirm'], { force: true });
    }
    callback();
  }

  public compareToFirstPassword(rule: any, value: string, callback: (errMsg?: string) => void) {
    if (value && value !== this.form.getFieldValue('password')) {
      callback('Two passwords that you enter is inconsistent!');
    } else {
      callback();
    }
  }

  public handleConfirmBlur(event: Event) {
    this.confirmDirty = this.confirmDirty || !!(event.target as HTMLInputElement).value;
  }

  public async register(event: Event) {
    event.preventDefault();
    this.form.validateFields(async (err: any, values: any) => {
      if (!err) {
        try {
          /* tslint:disable:no-string-literal */
          const registerBody = {
            username: values['username'],
            password: values['password'],
            name: values['name'],
            studentId: values['studentId'].toString(),
            birthday: new Date(values['birthday'].format('YYYY-MM-DD')).toISOString(),
            motto: values['motto'] || '',
          };
          /* tslint:enable:no-string-literal */
          await HTTP_CLIENT.post<{}>('/users', registerBody);
          this.$notification.success({ message: '注册成功', description: '请使用注册的账号登录'});
          this.$router.push('/login');
        } catch (error) {
          // Do nothing
        }
      }
    });
  }
}
</script>

<style lang="less" scoped>
#register-form {
  width: 540px;
  padding: 48px 48px 0 48px;
  background-color: white;
  border-radius: 10px 10px #333 10px;
  .submit-button {
    width: 100%;
  }
}
</style>