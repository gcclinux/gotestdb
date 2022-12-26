package conn

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Host []string
	Port []string
	User []string
	Pass []string
	DB   []string
}

func Getconf() (Host, Port, User, Pass, DB string) {

	file, _ := os.Open("conf/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	Host = conf.Host[0]
	Port = conf.Port[0]
	User = conf.User[0]
	Pass = conf.Pass[0]
	DB = conf.DB[0]

	return Host, Port, User, Pass, DB
}
