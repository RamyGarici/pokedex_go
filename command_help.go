package main
import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _,val:= range getCommands(){
		fmt.Printf("%v: %v\n",val.name,val.description)
	}
	return nil
}