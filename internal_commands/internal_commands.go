package internal_commands

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"strings"
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

func IsInternalCommand(Command string) bool {
	switch Command {
		case "--init":
			return true
		case "--help":
			return true
		default:
			return false
	}
}

func ConfigExists(CurrentDirectory string) bool {
	Info, err := os.Stat(fmt.Sprintf("%s/jcc.config.json",CurrentDirectory))
	if err != nil {
		return false
	}
	mode := Info.Mode()
	if !mode.IsRegular() {
		return false
	}
	return true
}

func Init(CurrentDirectory string) {
	if ConfigExists(CurrentDirectory) {
		fmt.Println("A Config File Alredy Exist")
		var flag bool = true
		for flag {
			fmt.Print("Would You still like to coninue [Y/N]:")
			var response string;
			_, err := fmt.Scanln(&response)
			Check(err)
			switch strings.ToUpper(response) {
				case "Y":
					flag = false
					break;
				case "N":
					fmt.Println("Exiting ..")
					return
				default:
					fmt.Println("Invalid Input")
			}
		}
	}
	
	ConfigFile, err := os.Create(fmt.Sprintf("%s/jcc.config.json",CurrentDirectory))
	Check(err)
	defer ConfigFile.Close()

	var Payload Config = Config{
		Aliases:  make(map[string]string),
		Multiple: make(map[string][]string),
	}

	DataBytes, err := json.MarshalIndent(Payload, "", "    ")
	Check(err)
	ConfigFile.Write(DataBytes)
	fmt.Println("\nCreated config File -> jcc.config.json")
	fmt.Print("Use")
	CyanColor := color.New(color.FgCyan)
	CyanColor.Print(" --help ")
	fmt.Println("To get List Of All internal Commands & Notes")
}

func Help() {
	CyanColor := color.New(color.FgCyan)
	CyanColor.Print("--init ")
	fmt.Println("Make Config File")
	CyanColor.Print("--Help ")
	fmt.Println("Gets List Of All internal Commands & Notes")
	fmt.Println("\nFYI DONT USE Muliple Internal Commands together.")
	fmt.Println("And Dont use internal commands together with your Defined Commands .")
	//CyanColor.Print("--Add ")
	//fmt.Println("Adds Aliase And Command to Config File Through Terminal")
	//fmt.Println("Usage ")
	
}

func RunInternalCommand(Arguments []string, CurrentDirectory string ) {
	//var WG sync.WaitGroup
	switch Arguments[1] {
		case "--init":
			Init(CurrentDirectory)
		case "--help":
			Help()
		//case "--add":
		//	Add(Arguments)
	}
	
}
