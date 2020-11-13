package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type PokeAPIPokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string      `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string      `json:"back_default"`
						BackFemale       interface{} `json:"back_female"`
						BackShiny        string      `json:"back_shiny"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontDefault     string      `json:"front_default"`
						FrontFemale      interface{} `json:"front_female"`
						FrontShiny       string      `json:"front_shiny"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string      `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string      `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type BattlePoke struct {
	Stats  BattleStats
	Health int
	Name   string
}

type BattleStats struct {
	hp      int
	attack  int
	defense int
	speed   int
}

func main() {

	pokemonList := loadPokemon()
	for _, poke := range pokemonList {
		printPoke(poke)
	}

	runBattle(pokemonList)

}

func initBattlePoke(poke PokeAPIPokemon) BattlePoke {
	pStats := BattleStats{hp: poke.Stats[0].BaseStat, attack: poke.Stats[1].BaseStat, defense: poke.Stats[2].BaseStat, speed: poke.Stats[5].BaseStat}
	p := BattlePoke{Name: poke.Name, Stats: pStats, Health: poke.Stats[0].BaseStat}
	return p
}

func runBattle(pokemon [6]PokeAPIPokemon) {
	team1 := [3]PokeAPIPokemon{}
	team2 := [3]PokeAPIPokemon{}
	for i := 0; i < 3; i++ {
		team1[i] = pokemon[i]
		team2[i] = pokemon[i+3]
	}

	pokeIndex := 0
	poke1InArena := initBattlePoke(team1[pokeIndex])
	poke2InArena := initBattlePoke(team2[pokeIndex])

	fmt.Printf("****** %s vs. %s ******\n", poke1InArena.Name, poke2InArena.Name)

	for pokeIndex < 3 {
		printBattlePoke(poke1InArena)
		printBattlePoke(poke2InArena)
		poke1InArena, poke2InArena = runAttack(poke1InArena, poke2InArena)

		if poke1InArena.Health < 1 || poke2InArena.Health < 1 {
			if pokeIndex == 2 {
				break
			}
			pokeIndex++
			poke1InArena = initBattlePoke(team1[pokeIndex])
			poke2InArena = initBattlePoke(team2[pokeIndex])
		}

	}

	winner := poke1InArena
	if poke1InArena.Health < 1 {
		winner = poke2InArena
	}

	fmt.Printf("%s won!\n", winner.Name)
}

func runAttack(poke1 BattlePoke, poke2 BattlePoke) (BattlePoke, BattlePoke) {
	var attacker, defender *BattlePoke

	if poke1.Stats.speed > poke2.Stats.speed {
		attacker = &poke1
		defender = &poke2
	} else {
		attacker = &poke2
		defender = &poke1

	}

	defender.Health -= calcAttackPower(attacker.Stats.attack, defender.Stats.defense)

	if defender.Health < 1 {
		return poke1, poke2
	}

	temp := attacker
	attacker = defender
	defender = temp

	defender.Health -= calcAttackPower(attacker.Stats.attack, defender.Stats.defense)

	return poke1, poke2

}

func calcAttackPower(attack int, defense int) int {
	ap := attack - defense

	if ap < 1 {
		ap = 1
	}
	return ap
}

func printBattlePoke(poke BattlePoke) {
	fmt.Printf("%s\t\thp: %d\n", poke.Name, poke.Health)
}

func printPoke(poke PokeAPIPokemon) {

	fmt.Printf("=========== %s ===========\n", strings.ToUpper(poke.Name))
	for j := range poke.Stats {
		fmt.Printf("%s: %d, ", poke.Stats[j].Stat.Name, poke.Stats[j].BaseStat)
	}
	fmt.Println()

}

func loadPokemon() [6]PokeAPIPokemon {
	pokemonList := [6]PokeAPIPokemon{}

	for i := 0; i < 6; i++ {
		pokeNumber := rand.Intn(151)
		poke := getPokemon(pokeNumber)
		pokemonList[i] = poke

	}
	return pokemonList
}

func getPokemonFromJson(body []byte) (*PokeAPIPokemon, error) {
	var p = new(PokeAPIPokemon)
	err := json.Unmarshal(body, &p)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return p, err
}

func getPokemon(pokemonNumber int) PokeAPIPokemon {
	pokeUrl := "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(pokemonNumber)
	log.Println("fetching " + pokeUrl)

	resp, err := http.Get(pokeUrl)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	pokemon, err := getPokemonFromJson([]byte(body))

	return *pokemon

}
