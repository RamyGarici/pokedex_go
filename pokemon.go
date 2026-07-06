package main
import("net/http"
"encoding/json"
"fmt"
"io"
"github.com/RamyGarici/pokedex_go/internal/pokecache"
)

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}



type config struct{
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
}

func getLocation(url string, cfg *config) error{
	body, ok:= cfg.Cache.Get(url)
	if !ok{
	resp, err := http.Get(url)
	 
	if err!=nil{
		return err
	}
	defer resp.Body.Close()
    body,err = io.ReadAll(resp.Body)
	if err!=nil{
		return err
	}
    cfg.Cache.Add(url,body)}
   
	var data LocationAreaResponse
	err := json.Unmarshal(body,&data)
	if err!=nil{
		return err
	}
	
for _, area := range data.Results {
    fmt.Println(area.Name)
}

cfg.Next = data.Next
cfg.Previous = data.Previous
return nil


}

func getPokemon(url string, cfg *config) error{
	body, ok := cfg.Cache.Get(url)
	if !ok{
		resp,err:= http.Get(url)
		if err!=nil{
			return err
		}
		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		if err!=nil{
			return err
		}
		cfg.Cache.Add(url, body)
	}
	var data LocationAreaDetail
	err:= json.Unmarshal(body, &data)
	if err!=nil{
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}