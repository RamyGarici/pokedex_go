package main
import ("fmt"
"os"
"bufio"
"strings"
"time"
"github.com/RamyGarici/pokedex_go/internal/pokecache")

func startRepl() {
	cfg := &config{
		Cache:   pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]Pokemon),
	}
	commands := getCommands()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
if len(words)==0{
	continue
}
commandName := words[0]
command,exists := commands[commandName]
var args []string
if len(words) > 1 {
	args = words[1:]
}

if exists{
	err:= command.callback(cfg, args...)
	if err!= nil{
		fmt.Println(err)
	}
	continue
}else{
	fmt.Println("Unknown command")
	continue

}}}





func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}


type cliCommand struct{
	name string
	description string
	callback func(cfg *config, args ...string) error
}
func getCommands()  map[string]cliCommand{
	return  map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map":{
		name:        "map",
		description: "List Pokemon locations" ,
		callback:    commandMap,
	},
	"mapb":{
		name:        "mapb",
		description: "List Pokemon locations (Previous page)" ,
		callback:    commandMapb,
	},
	"explore":{
		name:        "explore",
		description: "List Pokemons of a location" ,
		callback:    commandExplore,
	},
	"catch":{
		name: "catch",
		description: "Catch a pokemon",
		callback: commandCatch,

	},
}
}