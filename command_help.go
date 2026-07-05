package main
import "fmt"

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _,val:= range getCommands(){
		fmt.Printf("%v: %v\n",val.name,val.description)
	}
	return nil
}