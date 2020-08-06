<template>
  <v-container>
    <v-row>
      <v-col cols="3">
        <v-text-field
          label="氏名コード"
          v-model="usercode"
          :rules="[() => !!usercode || 'Required']"
        />
      </v-col>
      <v-col cols="3">
        <v-text-field
          label="氏名"
          v-model="username"
          :rules="[() => !!username || 'Required']"
        />
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12">
        <v-simple-table fixed-header>
          <template v-slot:default>
            <colgroup class="task-category">
              <col span="1" />
            </colgroup>

            <colgroup class="task-description">
              <col span="1" />
            </colgroup>

            <colgroup class="task-estimate">
              <col span="1" />
            </colgroup>

            <colgroup class="task-operation">
              <col span="1" />
            </colgroup>
            <thead>
              <th>カテゴリ</th>
              <th>作業内容</th>
              <th>作業時間[h]</th>
              <th>操作</th>
            </thead>

            <tbody>
              <tr v-for="(taskrow, index) in taskrows" :key="taskrow.id">
                <td>
                  <v-select
                    label="Select"
                    :items="categories"
                    item-text="name"
                    item-value="val"
                    v-model="taskrow.category"
                  >
                  </v-select>
                </td>
                <td>
                  <v-text-field
                    v-model="taskrow.task"
                    @change="onChange($event)"
                  />
                </td>
                <td>
                  <v-text-field
                    placeholder="例:1.00"
                    v-model="taskrow.estimate"
                    type="number"
                    value=""
                  />
                </td>
                <td>
                  <v-btn color="nomal" @click="onRemoveRow(index)">削除</v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>

      <v-col cols="12">
        <v-btn color="primary" @click="onAddRow">追加</v-btn>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12">
        <v-textarea
          outlined
          placeholder="追加報告がある場合はこちらに追記してください"
          v-model="comment"
        >
        </v-textarea>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-btn
        color="error"
        @click="onSubmit"
        x-large
        width="200"
        :disabled="!isValid"
        >登録</v-btn
      >
    </v-row>
    <v-snackbar
      v-model="snackbar"
      :timeout="snackbarto"
      :top="true"
      :color="snackbarcolor"
    >
      {{ snackbarmsg }}
    </v-snackbar>
  </v-container>
</template>

<style>
.task-category {
  width: 60px;
}
.task-description {
  width: 200px;
}
.task-estimate {
  width: 10px;
}
.task-operation {
  width: 10px;
}
</style>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

class Category {
  val!: number;
  name!: string;
}

class Taskinfo {
  id!: number;
  category!: number;
  task!: string;
  estimate!: number | null;
}

class ReportTask {
  category!: number;
  task!: string;
  estimate!: number;
}

class ReportData {
  code!: string;
  name!: string;
  comment!: string;
  tasks!: ReportTask[];
}

@Component
export default class RegistryComp extends Vue {
  hoge!: number;
  taskrows!: Taskinfo[];
  categories!: Category[];
  defaultCateval!: string;
  comment!: string;
  usercode!: string;
  username!: string;
  isValid!: boolean;

  snackbar!: boolean;
  snackbarto!: number;
  snackbarcolor!: string;
  snackbarmsg!: string;

  clear() {
    this.isValid = false;
    this.hoge = 0;
    console.log("clear");
    this.defaultCateval = "0";
    this.taskrows = [
      { id: 0, category: 0, task: "", estimate: null },
      { id: 1, category: 0, task: "", estimate: null },
      { id: 2, category: 0, task: "", estimate: null }
    ];

    this.categories = [
      { val: 0, name: "" },
      { val: 1, name: "打ち合わせ" },
      { val: 2, name: "設計" },
      { val: 3, name: "開発" },
      { val: 4, name: "テスト" },
      { val: 5, name: "資料作成" }
    ];

    this.username = "";
    this.usercode = "";
    this.comment = "";
  }

  constructor() {
    super();
  }

  created() {
    this.snackbar = false;
    this.snackbarto = 3000;
    this.snackbarcolor = "success";
    this.snackbarmsg = "";
    this.clear();
  }

  onAddRow() {
    console.log("onAddRow");
    const newtask = new Taskinfo();
    newtask.id = this.taskrows.length;
    this.taskrows.push(newtask);
  }

  onRemoveRow(index: number) {
    console.log(index);
    this.taskrows.splice(index, 1);
  }

  isValidtask(task: Taskinfo): boolean {
    if (task.category !== 0 && task.task.length > 0 && task.estimate !== null) {
      return true;
    }
    return false;
  }

  createPostdata() {
    const data = new ReportData();
    data.code = this.usercode;
    data.name = this.username;
    data.comment = this.comment;

    const tasks: ReportTask[] = [];

    for (let idx = 0; idx < this.taskrows.length; idx++) {
      const row = this.taskrows[idx];
      if (!this.isValidtask(row)) {
        continue;
      }

      const task = new ReportTask();
      task.category = row.category;
      task.task = row.task;
      task.estimate = Number(row.estimate);
      tasks.push(task);
    }
    data.tasks = tasks;
    console.log(data);
  }

  onChange(evt: any) {
    console.log(evt);
  }

  async onSubmit() {
    const data = this.createPostdata();

    const param = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    };

    fetch(
      `http://192.168.1.161:13000/users/${this.usercode}reports/registry`,
      param
    )
      .then(res => {
        console.log(res);
        if (!res.ok) {
          throw Error(res.statusText);
        }
        return res;
      })
      .then(res => {
        console.log("[SUCCESS]");
        console.log(res);
        this.snackbar = true;
        this.snackbarmsg = "日報の登録が完了しました。";
        this.snackbarcolor = "success";
        this.clear();
      })
      .catch(err => {
        console.log("[ERROR]");
        console.log(err);
        this.snackbar = true;
        this.snackbarmsg = "日報の登録が失敗しました。";
        this.snackbarcolor = "error";
      });
  }
}
</script>
