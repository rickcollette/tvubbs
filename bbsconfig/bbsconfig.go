package bbsconfig

import (
	"fmt"
	"os"
	"tvubbs/dbstruct"

	"gopkg.in/yaml.v3"
)

var BbsConfig *dbstruct.Sysconfig
var MenuConfig *dbstruct.MenuStruct
var MessageBaseConfig *dbstruct.MessageBaseIdx
var UserConfig *dbstruct.Userdb
var MessageConfig *dbstruct.MessageBaseData
var MenuDataConfig *dbstruct.MenuData

func LoadConfig() error {
	fmt.Printf("Checking Databases...\n")
	file, err := os.Open("data/bbsconfig.yml")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err != nil {
		return err
	}
	return decoder.Decode(&BbsConfig)
}

func LoadMenuConfig() error {
	fmt.Printf("Checking Databases...\n")
	file, err := os.Open("data/menus.yml")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err != nil {
		return err
	}
	return decoder.Decode(&MenuConfig)
}