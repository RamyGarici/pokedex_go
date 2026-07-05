package main


func commandMap(cfg *config) error{
   url:="https://pokeapi.co/api/v2/location-area/"
   if cfg.Next != nil{
	url = *cfg.Next
   }
	return getLocation(url, cfg)
}