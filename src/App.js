import React from 'react';
import logo from './logo.svg';
import './App.css';

import { Pokemon } from "./poke"
import { Arena } from "./arena"


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <div className='header-title'>

        <p className='heavy'>Poke app</p>
        <p className='light'>
          DREW HIRSCHI
        </p>
        </div>
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <Pokemon />
      <Arena/>
    </div>
  );
}

export default App;
