<template>
  <div id="register">
    <RegisterForm></RegisterForm>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import RegisterForm from '@/components/RegisterForm.vue';

@Component({
  components: { RegisterForm },
})
export default class Register extends Vue { }

</script>

<style lang="less" scoped>
#register {
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