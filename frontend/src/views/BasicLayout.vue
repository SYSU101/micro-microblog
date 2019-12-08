<template>
  <div id="basic-layout">
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';

@Component({ })
export default class BasicLayout extends Vue { }

</script>

<style lang="less" scoped>
#basic-layout {
  width: 100vw;
  height: 100vh;
  background-color: #dddddd;
}
</style>
