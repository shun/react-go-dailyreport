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
                      v-model="defaultCateval"
                    >
                    </v-select>
                  </td>
                  <td>
                    <v-text-field
                      v-bind:value="taskrow.task"
                    />
                  </td>
                  <td>
                    <v-text-field
                      placeholder="例:1.00"
                      v-bind:value="taskrow.estimate"
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
        <v-btn color="error" @click="onSubmit" x-large width="200">登録</v-btn>
      </v-row>
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
    this.name = "";
    this.age = -1;
    this.role = "";
  }
};

@Component
export default class RegistryComp extends Vue {
  taskrows: Taskinfo[];
  categories: Category[];
  defaultCateval: string;

  constructor() {
    super();
    this.defaultCateval = "0";
    this.taskrows = [
      {id: 0, category: "0", task: "11", estimate: ""},
      {id: 1, category: "0", task: "22", estimate: ""},
      {id: 2, category: "0", task: "33", estimate: ""},
    ];

    this.categories = [
      {val: 0, name: ""},
      {val: 1, name: "打ち合わせ"},
      {val: 2, name: "設計"},
      {val: 3, name: "開発"},
      {val: 4, name: "テスト"},
      {val: 5, name: "資料作成"},
    ];
  }

  onAddRow() {
    //const newtask = {id: this.taskrows.length + 1, name: "", age: 0, role: ""};
    const newtask = new Taskinfo();
    //taskrows.push(new Taskinfo(id: taskrows.length + 1, name: "", age: 0, role: ""));
    this.taskrows.push(newtask);
  }

  onRemoveRow(index: number) {
    console.log(index);
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

      const param = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data)
      };

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
      })
      .catch((err) => {
        console.log("[ERROR]");
        console.log(err);
      });
  }
}
</script>

