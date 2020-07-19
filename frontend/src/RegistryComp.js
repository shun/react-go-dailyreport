import React, { Component } from 'react';

export default class Registry extends Component {
  showAlert() {
    const userinfo = document.getElementById('userinfo')
    const tasks = document.getElementById('task')
    debugger
    alert("登録完了しました");
  }

  render() {
    return (
      <div>
        <h2>日報登録</h2>
        <table id='userinfo'>
          <tbody>
          <tr>
            <th><span>氏名コード</span></th>
            <td><input id="usercode" type="text" placeholder="7桁の氏名コードを入力"></input></td>
          </tr>
          <tr>
            <th><span>氏名</span></th>
            <td><input id="username" type="text" placeholder="氏名を入力"></input></td>
          </tr>
          </tbody>
        </table><br/>

        <Task />
        <button type="button" onClick={this.showAlert}>日報登録</button>
      </div>
    )
  }
}

class Task extends Component {
  render() {
    return (
      <div>
        <table id='tasks'>
          <TaskHead/>
          <TaskBody/>
        </table>
      </div>
    )
  }
}

class CategoryOptions extends Component {
  render() {
    const categories = [
      {val: 0, option: ""},
      {val: 1, option: "資料作成"},
      {val: 2, option: "設計"},
      {val: 3, option: "コーディング"},
      {val: 4, option: "テスト"},
      {val: 5, option: "打ち合わせ"},
      {val: 6, option: "その他"},
    ];

    return (
      <select>
      {categories.map((category) => (
          <option key={category.val} value={category.val}>{category.option}</option>
      ))}
      </select>
    )

  }
}

class TaskRow extends Component {
  render() {
    const taskid="task[" + this.props.rowid + "]";
    const estimateid="estimate[" + this.props.rowid + "]";

    return (
      <tr>
        <td>
          <CategoryOptions/>
        </td>
        <td>
          <input id={taskid} type="text" size="40"></input>
        </td>
        <td>
          <input id={estimateid} type="text" placeholder="例:1.00" size="10"></input>
        </td>
      </tr>
    )
  }
}

class TaskHead extends Component {
  render() {
    return (
      <thead>
        <tr>
          <th>カテゴリ</th>
          <th>作業内容</th>
          <th>作業時間[h]</th>
        </tr>
      </thead>
    )
  }
}

class TaskBody extends Component {
  render() {
    const rownum = 10;
    let arr = [];

    for(let i = 0; i < rownum; i++) {
      arr.push(<TaskRow key={i} rowid={i} />);
    }
    return (
      <tbody>
        {arr}
      </tbody>
    )
  }
}

