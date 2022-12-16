package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type Databases struct {
	Userdb struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Status    string `yaml:"status"`
		Lastlogin string `yaml:"lastlogin"`
		Bulletins string `yaml:"bulletins"`
		Active    string `yaml:"active"`
		Security  string `yaml:"security"`
		Messages  string `yaml:"messages"`
		Tagline   string `yaml:"tagline"`
	} `yaml:"userdb"`
	Sysconfig struct {
		Bbsname    string `yaml:"bbsname"`
		Sysop      string `yaml:"sysop"`
		Hostname   string `yaml:"hostname"`
		Listenport string `yaml:"listenport"`
		Homedir    string `yaml:"homedir"`
	} `yaml:"sysconfig"`
}

type DataBase interface{}

func main() {
	checkInit()
}

func checkInit() {
	if !fileExists("./SYSTEM_INITIALIZED") {
		checkDatabases()
		createInit()
	} else {
		fmt.Println("Starting The BBS Daemon...")
	}
}

func readYaml(filename string) map[string]interface{} {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var superList map[string]interface{}
	err = decoder.Decode(&superList)
	if err != nil {
		fmt.Println(err)
	}
	return superList
}

func createInit() {
	f, err := os.Create("./SYSTEM_INITIALIZED")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}

// read yaml file and print key value
func checkDatabases() {
	fmt.Printf("Checking Databases...\n")
	out := Data{}
	file, err := os.ReadFile("databases.yml")
	fmt.Printf("File: %s\n", file)

	if err != nil {
		fmt.Println(err)
	}
	if err := yaml.Unmarshal(file, &out); err != nil {
		fmt.Println(err)
	}
	for _, database := range out.DataBase {
		fmt.Println("Got to the loop")
		fmt.Println("Checking ", database)
		fmt.Println(database.(string))

	}
}

func fileExists(bbsfile string) bool {
	info, err := os.Stat(bbsfile)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func createDatabase(filename string) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// create table
	sql_table := `
	CREATE TABLE IF NOT EXISTS user(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT,
		status TEXT,
		lastlogin TEXT,
		bulletins TEXT,
		active BOOLEAN,
		security TEXT,
		messages TEXT,
		tagline TEXT
	);
	`
	_, err = db.Exec(sql_table)
	if err != nil {
		fmt.Println(err)
	}
}
