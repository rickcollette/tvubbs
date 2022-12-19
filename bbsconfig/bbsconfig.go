package bbsconfig

import (
	"fmt"
	"os"
	"tvubbs/dbstruct"

	"gopkg.in/yaml.v3"
)

var BbsConfig *dbstruct.Sysconfig

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
