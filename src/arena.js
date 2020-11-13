import React from 'react';
import axios from 'axios'
import {init} from './arenaController';

export class Arena extends React.Component {

    constructor(params) {
        super()
        this.arenaRef = React.createRef();
    }

    componentDidMount() {
        console.log('doing init');
        init(this.context)
    }

    render() {
        
        return (
            <div className='poke'>
                <h1 className='poke-header'>Pokemon Arena</h1>
                <section className='poke-container'>
                <canvas id="poke-arena" ref={(c) => {this.context = c.getContext('2d'); console.log('set context');}}
                width="800" height="1300" style={{ background: 'white' }}/>
                </section>
            </div>

        )
    }
}