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
	if c.MenuType == "menu" {
		PrintMenu(c.MenuFile)
	} else if c.MenuType == "command" {
		
