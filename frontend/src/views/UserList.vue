<template>
  <a-table :columns="columns" :dataSource="data" />
</template>

<script lang='ts'>
import { Vue, Component } from 'vue-property-decorator';
import { SILENT_HTTP_CLIENT } from '../utils/HttpClient';
import { IUserProfile } from '../typings';
import moment from 'moment';

@Component({})
export default class UserList extends Vue {
  private data: Array<Partial<IUserProfile>> = [];
  private columns = [
    {
      title: 'Username',
      dataIndex: 'username',
      key: 'username',
    },
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Student ID',
      key: 'studentId',
      dataIndex: 'studentId',
    },
    {
      title: 'Birthday',
      key: 'birthday',
      dataIndex: 'birthday',
    },
    {
      title: 'Motto',
      key: 'motto',
      dataIndex: 'motto',
    },
  ];

  public async mounted() {
    this.data = (await SILENT_HTTP_CLIENT.get<{ user: Array<Partial<IUserProfile>> }>('/users'))
      .user
      .map((user) => {
        user.birthday = moment(new Date(user.birthday || 0).toISOString()).format('YYYY-MM-DD');
        return user;
      });
  }
}
</script>