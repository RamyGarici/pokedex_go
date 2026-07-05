package main
import("net/http"
"encoding/json"
"fmt")

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

type config struct{
	Next *string
	Previous *string
}

func getLocation(url string, cfg *config) error{
	resp, err := http.Get(url)
	if err!=nil{
		return err
	}
	var data LocationAreaResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	
for _, area := range data.Results {
    fmt.Println(area.Name)
}

cfg.Next = data.Next
cfg.Previous = data.Previous
return nil


}