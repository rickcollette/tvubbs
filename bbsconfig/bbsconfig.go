package bbsconfig

import (
	"fmt"
	"os"
	"tvubbs/dbstruct"

	"gopkg.in/yaml.v3"
)


func LoadConfig() (*dbstruct.Sysconfig, error) {
	fmt.Printf("Checking Databases...\n")
	BaseConfig := &dbstruct.Sysconfig{}
	file, err := os.Open("data/bbsconfig.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&BaseConfig); err != nil {
		return nil, err
	}
	return (BaseConfig), nil
}
