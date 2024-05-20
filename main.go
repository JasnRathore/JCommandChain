package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"github.com/fatih/color"
	IC "github.com/JasnRathore/JCommandChain/internal_commands"
)

type Config struct {
	Aliases  map[string]string
	Multiple map[string][]string
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RunCommand(command string, WG *sync.WaitGroup) {
	defer WG.Done()

	NewCommand := exec.Command("sh", "-c", command)
	NewCommand.Stdout = os.Stdout
	NewCommand.Stderr = os.Stderr

	err := NewCommand.Run()
	Check(err)
}

func ReadConfig(CurrentDirectory string) Config {
DataBytes, err := os.ReadFile(fmt.Sprintf("%s/jcc.config.json",CurrentDirectory))
	Check(err)

	var Data Config
	err = json.Unmarshal(DataBytes, &Data)
	Check(err)
	return Data
}
	
func Execution(ConfigData Config) {
	var WG sync.WaitGroup

	for i, arg := range os.Args {
		if i==0 {
			continue
		}
		CommandValue, isInAliases := ConfigData.Aliases[arg]
		_, isInMultiple := ConfigData.Multiple[arg]
		if !isInMultiple && !isInAliases  {
			fmt.Println("Unknown command:", arg)
			continue
		} 
		if isInMultiple && isInAliases {
			color.Red("Command Found in Both Aliases and mulitple")
			fmt.Println("->",arg);
			break
		}
		if isInAliases {
			WG.Add(1)
			go RunCommand(CommandValue, &WG)
		} else {
			Commands, _ := ConfigData.Multiple[arg];
			for _, arg2 := range Commands {
				CommandValue2, isInAliases := ConfigData.Aliases[arg2]
				if !isInAliases {
					fmt.Println("Unknown command:", arg2)
					continue
				}
				WG.Add(1)
				go RunCommand(CommandValue2, &WG)
			}
		}
		
	}
	WG.Wait()
}

func main() {
	CurrentDirectory, err := os.Getwd()
	Check(err)
	
	Arguments := os.Args
	if len(Arguments) < 2 {
		fmt.Println("No Argument Provided")
		return
	}

	var IsInternal bool = IC.IsInternalCommand(Arguments[1])
	if IsInternal {
		IC.RunInternalCommand(Arguments, CurrentDirectory)
		return
	}
	
	if !IC.ConfigExists(CurrentDirectory) {
		CyanColor := color.New(color.FgCyan)
		color.Red("Config file not found;")
			
		fmt.Print("Use") 
		CyanColor.Print(" --init ")
		fmt.Println("To Create Config File;")
			
		fmt.Print("Use") 
		CyanColor.Print(" --help ")
		fmt.Print("For More Internal Commands;")
		
		return
	}
	var ConfigData Config = ReadConfig(CurrentDirectory);
	Execution(ConfigData)
}
