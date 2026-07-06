package main


func commandMap(cfg *config, args ...string) error {
   url:="https://pokeapi.co/api/v2/location-area/"
   if cfg.Next != nil{
	url = *cfg.Next
   }
	return getLocation(url, cfg)
}
func commandMapb(cfg *config, args ...string) error {
   url:="https://pokeapi.co/api/v2/location-area/"
   if cfg.Previous != nil{
	url = *cfg.Previous
   }
	return getLocation(url, cfg)
}