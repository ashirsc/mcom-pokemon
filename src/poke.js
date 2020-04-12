import React from 'react';
import { Button, Card } from 'antd'

const { Meta } = Card


export class Pokemon extends React.Component {

    state = {
        pokemon: [],
        loading: false
    }

    getPokemon = async (numPokemon) => {

        const endpoint = 'https://pokeapi.co/api/v2/pokemon/'
        let allPokemon = []
        let reqs = []
    
        for (let i = 1; i <= numPokemon; i++) {
            let pokemonReq = fetch(endpoint + i)
            // allPokemon.push(pokemon)
            reqs.push(pokemonReq)
        }
        const response = await Promise.all(reqs)

    
       this.setState({pokemon: response})
    }

    render() {

        return (
            <div className='poke'>
                <h1 className='poke-header'>Pokemon</h1> 
                <Button  onClick={()=> {this.getPokemon(15)}} type='primary' style={{width: 160, margin:  20}}>
                    Load Pokemon
                </Button>
                <section className='poke-container'>

                    {this.state.pokemon.map((poke,i) => {
                        return (
                            <Card
                                onClick={() => {
                                    window.open(`localhost:8080/pokemon/${i+1}/download`, "_blank")
                                }}
                                className='poke-card'
                                cover={<img alt={poke.forms[0].name} src={poke.sprites.front_default}/> }
                            >
                                <Meta title= {poke.forms[0].name.toUpperCase()} />
                                </Card>)
                    })
                }
                </section>
            </div>

        )
    }
}