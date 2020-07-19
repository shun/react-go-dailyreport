import React from 'react';
//import { Component } from 'react';
import { BrowserRouter, Route, Link } from 'react-router-dom';

const App = () => (
  <BrowserRouter>
    <div>
      <ul>
        <li><Link to='/'>Home</Link></li>
        <li><Link to='/registry'>日報登録</Link></li>
      </ul>
      <hr/>
      <Route exact path='/' component={Home} />
      <Route path='/registry' component={Registry} />
    </div>
  </BrowserRouter>
)

const Home = () => (
  <div>
    <h2>Home</h2>
    <p>Welcome to ようこそ</p>
  </div>
)

function showAlert() {
  alert("登録完了しました");
}

const Registry = () => (
  <div>
    <h2>日報登録</h2>
    <table>
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

    <Task/>
    <button type="button" onClick={showAlert}>日報登録</button>
  </div>
)

function CategoryOptions () {
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

function TaskRow({rowid}) {
  const taskid="task[" + rowid + "]";
  const estimateid="estimate[" + rowid + "]";

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

function TaskHead() {
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

function TaskBody() {
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

const Task = () => (
  <div>
    <table>
      <TaskHead/>
      <TaskBody/>
    </table>
  </div>
)

export default App

