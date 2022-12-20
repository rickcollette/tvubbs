package menuparser

import (
	"fmt"
	"os"
	"tvubbs/bbsconfig"
	"tvubbs/dbstruct"
)

var ErrorFile string = "ascii/notfound.txt"

func PrintMenu(filename string) {
	if filename == "" {
		filename = ErrorFile
	}
	menu, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading menu:", err)
		return
	}
	fmt.Println(menu)
}

func ParseMenu(c *dbstruct.MenuStruct) {
	if c == nil {
		PrintMenu(ErrorFile)
		return
	}
	if c.Type == "menu" {
		PrintMenu(c.MenuFile)
	} else if c.MenuType == "command" {

	}
}

func GetMenuName(name string) *dbstruct.MenuStruct {
	for _, v := range bbsconfig.MenuConfig.Name {
		if v == name {
			
			return v
		}
