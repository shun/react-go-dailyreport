import React, { Component } from 'react';
import { BrowserRouter, Route, Link } from 'react-router-dom';
import Home from './HomeComp';
import Registry from './RegistryComp.js';

export default class App extends Component {
  render() {
    return (
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
  }
}

