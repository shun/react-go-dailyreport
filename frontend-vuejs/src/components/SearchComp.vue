<template>
  <v-container>
    <v-row>
      <v-icon>calendar_today</v-icon>
      <v-col cols="2">
        <DatePickerComponent label="From" @selected="selectedFrom" />
      </v-col>

      <v-icon>calendar_today</v-icon>
      <v-col cols="2">
        <DatePickerComponent label="To" @selected="selectedFrom" />
      </v-col>
      <v-col :align-self="center" :align="center">
        <v-btn fab dark color="indigo">
          <v-icon>search</v-icon>
        </v-btn>
      </v-col>
    </v-row>

    <v-data-table :headers="headers" :items="dailyreports" item-key="date">
      <template v-slot:item.actions="{ item }">
        <v-btn :href="item.url">詳細</v-btn>
      </template>
    </v-data-table>
  </v-container>
</template>

<style></style>

<script lang="ts">
import { Component, Vue, Prop, Emit } from "vue-property-decorator";
import DatePickerComponent from "@/components/DatePicker/DatePicker.component.vue";

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

class ReportSummary {
  date!: string;
  code!: string;
  name!: string;
  worktime!: number;
  worktimeStr!: string;
  url!: string;
}

class ReportData {
  date!: string;
  code!: string;
  name!: string;
  comment!: string;
  tasks!: ReportTask[];
}

class TableHeader {
  text!: string;
  value!: string;
  sortable?: boolean;
  align?: "start" | "center" | "end";
}

@Component({
  components: {
    DatePickerComponent,
  },
})
export default class SearchComp extends Vue {
  selectedDateFrom = new Date().toISOString().substr(0, 10);
  selectedDateTo = new Date().toISOString().substr(0, 10);
  showCalendarFrom = false;
  showCalendarTo = false;
  dailyreports: ReportSummary[] = [];
  headers: TableHeader[];

  selectedFrom(dateStr: string) {
    console.log(`hogehoge: emitted by child: ${dateStr}`);
  }

  selectedTo(dateStr: string) {
    console.log(`hogehoge: emitted by child: ${dateStr}`);
  }

  onEndSelectedDate() {
    this.showCalendarFrom = false;
    this.showCalendarTo = false;
  }

  clear() {
    console.log("clear");
  }

  getDailyreport() {
    console.log("getDailyreport");
  }

  constructor() {
    super();
    this.headers = [
      { text: "日付", value: "date", align: "center" },
      { text: "氏名コード", value: "code", align: "center" },
      { text: "氏名", value: "name", align: "center" },
      { text: "工数", value: "worktimeStr", align: "center" },
      { text: "操作", value: "actions", sortable: false, align: "center" }
    ];
    this.dailyreports = [
      {
        date: "2020-07-20",
        code: "1017460",
        name: "工藤 俊輔",
        worktime: 9.25,
        worktimeStr: "9.25",
        url: "users/1017460/2020-07-20"
      },
      {
        date: "2020-07-19",
        code: "1017460",
        name: "工藤 俊輔",
        worktime: 9.0,
        worktimeStr: "9.00",
        url: "users/1017460/2020-07-20"
      },
      {
        date: "2020-07-17",
        code: "1017460",
        name: "工藤 俊輔",
        worktime: 9,
        worktimeStr: "9.00",
        url: "users/1017460/2020-07-20"
      }
    ];
  }

  created() {
    console.log("created");
  }

  mounted() {
    console.log("mounted");
  }

  //onClick(item: ReportSummary) {
  onClick(item: string) {
    console.log(item);
    //console.log(item.code);
    //console.log(item.date);
  }
}
</script>
