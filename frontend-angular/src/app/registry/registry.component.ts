import { Component, OnInit } from '@angular/core';
import { DataStoreService } from '../data-store.service';
import { MessageService } from 'primeng/api';

class ReportData {
  name!: string;
  code!: string;
  comment!: string;
  tasks!: ReportTask[];

  constructor() {
    this.name = "";
    this.code = "";
    this.comment = "";
    this.tasks = [];
  }
}

interface ReportTask {
  category: number;
  task: string;
  estimate: number;
}

interface Category {
  label: string;
  value: number;
}

class ShozokuCategories {
  categories!: Category[];
  constructor(categories: Category[]) {
    this.categories = categories;
  }
}

interface Shozoku {
  label: string;
  value: number;
}

@Component({
  selector: 'app-registry',
  templateUrl: './registry.component.html',
  styleUrls: ['./registry.component.styl'],
  providers: [MessageService]
})
export class RegistryComponent implements OnInit {

  registryDate: Date;
  username: string;
  usercode: string;
  reports: ReportTask[];
  shozokus: Shozoku[];
  categories: Category[];
  shozoku: Shozoku;
  categorylist: ShozokuCategories[];
  comment: string;

  constructor(
    private dataStoreService: DataStoreService,
    private messageService: MessageService
             ) {
    this.categorylist = [];
    this.reports = [];
    this.registryDate = new Date();
    this.comment = "";
  }


  ngOnInit(): void {
    this.usercode = "";
    this.username = "";
    this.reports = [
      {category: 0, task: "", estimate: null}
    ];

    this.shozokus = [
      {label: "", value: 0},
      {label: "経営企画部", value: 1},
      {label: "開発統括部", value: 2},
      {label: "CSS部", value: 3},
      {label: "基盤開発部", value: 4},
      {label: "パワートレイン部", value: 5},
    ];

    this.categorylist.push(new ShozokuCategories([
    ]));
    this.categorylist.push(new ShozokuCategories([
      {label: "", value: 0},
      {label: "打ち合わせ", value: 1},
      {label: "資料作成", value: 2},
      {label: "その他", value: 3},
    ]));
    this.categorylist.push(new ShozokuCategories([
      {label: "", value: 0},
      {label: "打ち合わせ", value: 1},
      {label: "資料作成", value: 2},
      {label: "テスト", value: 3},
    ]));

  }

  onRemove(event, index) {
    this.reports.splice(index, 1);
  }

  onAppend() {
    this.reports.push({category: null, task: "", estimate: null});
  }

  onChangeCategory(event, index) {
    this.reports[index].category = event.value;
  }

  onChangeShozoku(event) {
    this.categories = this.categorylist[event.value].categories;
  }

  async onRegistery() {
    if (!this.validatePostdata()) return;
    const reportdata = new ReportData();
    reportdata.code = this.usercode;
    reportdata.name = this.username;
    reportdata.comment = this.comment;
    reportdata.tasks = this.reports;

    delete reportdata.name;
    console.log(reportdata);
    await this.dataStoreService.postDate(reportdata)
    .then((ret) => {
      console.log(ret);
      if (ret) {
        this.messageService.add({key: "showtoast", severity:'success', summary:'日報を登録しました。'});
      } else {
        this.messageService.add({key: "showtoast", severity:'error', summary:'日報の登録に失敗しました。'});
      }
    });
  }

  validatePostdata(): boolean {
    const code = this.usercode.trim();
    const name = this.username.trim();

    const reCode = /^[0-9]+$/;
    if (7 !== code.length) return false;
    if (!reCode.exec(code)) return false;

    for (const report of this.reports) {
      if (null === report.estimate) return false;
      const estimate = Number(report.estimate);
      if (isNaN(estimate)) return false;
      report.estimate = estimate;
      const task = report.task.trim();
      if (0 === report.category) return false;
      if (0 === task.length) return false;
      if (0 === report.estimate) return false;
    }

    return true;
  }
}
