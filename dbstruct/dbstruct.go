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
	Desc   string `yaml:"desc"`
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
