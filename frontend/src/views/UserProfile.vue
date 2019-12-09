<template>
  <div id="user-profile">
    <a-row>
      <a @click="editing = true" v-if="!editing"><a-icon type="edit" />Edit</a>
      <a @click="update" v-if="editing" style="margin-right: 1em;"><a-icon type="check" />Save</a>
      <a @click="cancel" v-if="editing"><a-icon type="close" />Cancel</a>
    </a-row>
    <a-row v-for="profileItem of profileItems" :key="profileItem.key" style="margin-bottom: 0.5em">
      <a-col :span="6"><span class="label">{{ profileItem.key }}</span></a-col>
      <a-col :span="18">
        <span v-if="!editing">{{ profileItem.value }}</span>
        <a-input v-else-if="profileItem.key !== 'birthday'" v-model="profileItem.value"/>
        <a-date-picker v-else :defaultValue="toMoment(profileItem.value, 'YYYY-MM-DD')" :format="'YYYY-MM-DD'" @change="handlePickBirthday" />
      </a-col>
    </a-row>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import moment from 'moment';
import { HTTP_CLIENT } from '../utils/HttpClient';
import { IUserProfile } from '../typings';

@Component({})
export default class UserProfile extends Vue {
  public profileItems: Array<{ key: string, value: string }> = [];
  public editing = false;

  public resetItems() {
    // tslint:disable-next-line:no-string-literal
    const userProfile = this.$store.getters['getUserProfile']();
    this.profileItems = [];
    for (const key of Object.keys(userProfile)) {
      if (key === 'birthday') {
        this.profileItems.push({ key, value: moment(userProfile[key]).format('YYYY-MM-DD')});
      } else if (key !== 'id' && key !== 'username' && key !== 'password') {
        this.profileItems.push({ key, value: userProfile[key]});
      }
    }
  }

  public toMoment(...args: any[]) {
    return moment(...args);
  }

  public mounted() {
    this.resetItems();
  }

  public handlePickBirthday(date: any, dateString: string) {
    this.profileItems.forEach((item) => {
      if (item.key === 'birthday') {
        item.value = dateString;
      }
    });
  }

  public async update() {
    try {
      // tslint:disable-next-line:no-string-literal
      const id = this.$store.getters['getUserId']() || -1;
      const newProfile: { [key: string]: string } = {};
      this.profileItems.forEach((item) => {
        if (item.key !== 'birthday') {
          newProfile[item.key] = item.value;
        } else {
          newProfile[item.key] = moment(item.value, 'YYYY-MM-DD').toISOString();
        }
      });
      await HTTP_CLIENT.put<{}>(`/user/${id}`, newProfile);
      this.$notification.success({ message: '信息更新成功', description: '' });
      await this.$store.dispatch('fetchUserProfile', id);
      this.editing = false;
    } catch (error) {
      // Do nothing
    }
  }

  public cancel() {
    this.resetItems();
    this.editing = false;
  }
}

</script>

<style lang="less" scoped>
#user-profile {
  width: 100%;

  .label {
    text-align: right;
    display: inline-block;
    width: 100%;
    font-weight: bold;
    transform: translate(-1em);

    &::after {
      content: ':';
    }
  }
}
</style>