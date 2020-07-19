<template>
  <v-container>
      <v-row>
          <v-col cols="3">
            <v-text-field
              label="氏名コード"
            />
          </v-col>
          <v-col cols="3">
            <v-text-field
              label="氏名"
            />
          </v-col>
      </v-row>

      <v-row justify="center">
        <v-col cols="12">
          <v-simple-table fixed-header>
            <template v-slot:default>
            <colgroup class="task-category">
              <col span="1">
            </colgroup>

            <colgroup class="task-description">
              <col span="1">
            </colgroup>

            <colgroup class="task-estimate">
              <col span="1">
            </colgroup>

            <colgroup class="task-operation">
              <col span="1">
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
                      return-object
                    >
                    </v-select>
                  </td>
                  <td>
                    <v-text-field
                      v-model="taskrow.task"
                    />
                  </td>
                  <td>
                    <v-text-field
                      placeholder="例:1.00"
                      v-model="taskrow.estimate"
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
          <v-textarea outlined placeholder="追加報告がある場合はこちらに追記してください">
          </v-textarea>
        </v-col>
      </v-row>
      <v-row justify="center">
        <v-btn color="error" @click="onSubmit" x-large width="200">登録</v-btn>
      </v-row>
      <v-snackbar
        v-model="snackbar"
        :timeout="snackbarto"
        :top=true
        :color="snackbarcolor"
      >
        {{snackbarmsg}}
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
  val: number;
  name: string;
  constructor() {
    this.val = 0;
    this.name = "";
  }
}

class Taskinfo {
  id: number;
  category: string;
  task: string;
  estimate: string;

  constructor() {
    this.id = -1;
    this.category = "";
    this.task = "";
    this.estimate = "";
  }
};

@Component
export default class RegistryComp extends Vue {
  taskrows: Taskinfo[];
  categories: Category[];
  defaultCateval: string;
  comment: string;
  snackbar: boolean;
  snackbarto: number;
  snackbarcolor: string;
  snackbarmsg: string;

  initialize() {
    console.log("initialize");
    this.defaultCateval = "0";
    this.taskrows = [
      {id: 0, category: "0", task: "", estimate: ""},
      {id: 1, category: "0", task: "", estimate: ""},
      {id: 2, category: "0", task: "", estimate: ""},
    ];

    this.categories = [
      {val: 0, name: ""},
      {val: 1, name: "打ち合わせ"},
      {val: 2, name: "設計"},
      {val: 3, name: "開発"},
      {val: 4, name: "テスト"},
      {val: 5, name: "資料作成"},
    ];

    this.comment = "";
  }

  constructor() {
    super();
    this.defaultCateval = "0";
    this.taskrows = [
      {id: 0, category: "0", task: "", estimate: ""},
      {id: 1, category: "0", task: "", estimate: ""},
      {id: 2, category: "0", task: "", estimate: ""},
    ];

    this.categories = [
      {val: 0, name: ""},
      {val: 1, name: "打ち合わせ"},
      {val: 2, name: "設計"},
      {val: 3, name: "開発"},
      {val: 4, name: "テスト"},
      {val: 5, name: "資料作成"},
    ];

    this.comment = "";
    this.snackbar = false;
    this.snackbarto = 3000;
    this.snackbarcolor = "success";
    this.snackbarmsg = "";
    this.initialize();
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

  async onSubmit() {
    const data = {
      code:"111",
      name:"kudo shunsuke",
      comment:"ほげほげ",
      tasks: [
        {
          category: 1,
          task: "hoge1",
          estimate: 1.5
        },
        {
          category: 2,
          task: "hoge2",
          estimate: 3.5
        }
      ]
    };
    console.log(data);
    console.log(this.taskrows);

    const param = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data)
    };

    //const self = this;
    fetch("http://192.168.1.161:13000/users/111/reports/registry", param)
    .then((res) => {
      console.log(res);
      if (!res.ok) {
        throw Error(res.statusText);
      }
      return res;
    })
    .then((res) => {
      console.log("[SUCCESS]");
      console.log(res);
      this.snackbar = true;
      this.snackbarmsg = "日報の登録が完了しました。";
      this.snackbarcolor = "success";
      this.initialize();
    })
    .catch((err) => {
      console.log("[ERROR]");
      console.log(err);
      this.snackbar = true;
      this.snackbarmsg = "日報の登録が失敗しました。";
      this.snackbarcolor = "error";
    });
  }
}
</script>

