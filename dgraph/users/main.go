package main

import (
	"fmt"
	"github.com/rogercoll/microservices/dgraph/users/config"
)

func main() {
	sess := config.GetSession()
	fmt.Println(sess)
}