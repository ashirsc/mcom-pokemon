import React from 'react';
import axios from 'axios'



export class Pokemon extends React.Component {

    state = {
        pokemon: [],
        loading: false,
        numPokemon: 151
    }

    onChange = (value) => {
        this.setState({ numPokemon: value })
    }

    getPokemon = async (numPokemon) => {
        this.setState({ loading: true })

        const endpoint = 'https://pokeapi.co/api/v2/pokemon/'
        let reqs = []


        for (let i = 1; i <= numPokemon; i++) {
            // let pokemonReq =  (await fetch(endpoint + i)).body
            let pokemonReq = axios.get(endpoint + i)
            reqs.push(pokemonReq)
        }
        let response = await Promise.all(reqs)
        response = response.map(res => res.data)
        // response = await Promise.all(response)

        console.log(response);



        this.setState({ pokemon: response, loading: false })
    }

    render() {

        return (
            <div className='poke'>
                <h1 className='poke-header'>Pokemon</h1>
                <div>

                    <button loading={this.state.loading} onClick={() => { this.getPokemon(this.state.numPokemon) }} type='primary' style={{ width: 160, margin: 20 }}>
                        Load Pokemon
                </button>
                    <input onChange={this.onChange} value={this.state.numPokemon} />
                </div>

                <section className='poke-container'>

                    {this.state.pokemon.map((poke, i) => {
                        return (
                            <div>
                                <img className="poke-card" alt={poke.name} src={poke.sprites.front_default} />
                                <p>{poke.name.toUpperCase()}</p>
                            </div>

                        )
                    })
                    }
                </section>
            </div>

        )
    }
}