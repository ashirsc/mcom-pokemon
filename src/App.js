import React from 'react';
import logo from './logo.svg';
import './App.css';

import { Pokemon } from "./poke"

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <div className='header-title'>

        <p className='heavy'>Business Article Presentation</p>
        <p className='light'>
          DREW HIRSCHI
        </p>
        </div>
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <Pokemon />
    </div>
  );
}

export default App;
