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

const Category = () => (
  <select>
    <option value="0"></option>
    <option value="1">資料作成</option>
    <option value="2">設計</option>
    <option value="3">コーディング</option>
    <option value="4">テスト</option>
    <option value="5">打ち合わせ</option>
    <option value="6">その他</option>
  </select>
)

const TaskRow = () => (
  <tr>
    <td>
      <Category/>
    </td>
    <td>
      <input id="task" type="text" size="40"></input>
    </td>
    <td>
      <input id="estimate" type="text" placeholder="1.00" size="10"></input>
    </td>
  </tr>
)

const Task = () => (
  <div>
    <table>
      <thead>
        <tr>
          <th>カテゴリ</th>
          <th>作業内容</th>
          <th>作業時間[h]</th>
        </tr>
      </thead>
      <tbody>
        <TaskRow/>
        <TaskRow/>
        <TaskRow/>
        <TaskRow/>
        <TaskRow/>
      </tbody>
    </table>
  </div>
)

export default App

