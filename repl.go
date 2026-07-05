package main
import ("fmt"
"os"
"bufio"
"strings")

func startRepl() {
	cfg := &config{}
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
if exists{
	err:= command.callback(cfg)
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
	callback func(cfg *config) error
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

}
}