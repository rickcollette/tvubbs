package dbstruct

type Userdb struct {
	Username  string
	Password  string
	Status    string
	Lastlogin string
	Bulletins string
	Active    string
	Security  string
	Messages  string
	Tagline   string
	Available string
	Location  string
	Terminal  string
}

type Sysconfig struct {
	Bbsname  string   `yaml:"bbsname"`
	Sysop    string   `yaml:"sysop"`
	BindAddr string   `yaml:"bindaddr"`
	BindPort string   `yaml:"bindport"`
	Homedir  string   `yaml:"homedir"`
	Rooms    []string `yaml:"rooms"`
}

type MessageBaseIdx struct {
	Name   string `yaml:"name"`
	Level  string `yaml:"level"`
	Desc   string `yaml:"description"`
	Status string `yaml:"status"`
	Type   string `yaml:"type"`
}

type MessageBaseData struct {
	Indexid   int
	Writtenby string
	Date      string
	Subject   string
	Body      string
}

type MenuStruct struct {
	MenuNumber string `yaml:"menus"`
	Name string `yaml:"name"`
	Desc string `yaml:"description"`
	File string `yaml:"file"`
	Level string `yaml:"level"`
	Status string `yaml:"status"`
}

type MenuData struct {
	CommandNumber string `yaml:"commands"`
    MenuNumber string `yaml:"menu"`
	Command string `yaml:"command"`
	Shortcut string `yaml:"shortcut"`
	Type string `yaml:"type"`
	Arguments string `yaml:"arguments"`
	Level string `yaml:"level"`
}