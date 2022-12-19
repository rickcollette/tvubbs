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

func main() {
	checkInit()
}

func checkInit() {
	if !fileExists("./SYSTEM_INITIALIZED") {
		checkDatabases()
		//		createInit()
	} else {
		fmt.Println("Starting The BBS Daemon...")
	}
}

func createInit() {
	f, err := os.Create("./SYSTEM_INITIALIZED")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}

// read yaml file and print key value
func checkDatabases() (*Databases, error) {
	fmt.Printf("Checking Databases...\n")
	databases := &Databases{}
	file, err := os.Open("databases.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&databases); err != nil {
		return nil, err
	}

	return databases, nil
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