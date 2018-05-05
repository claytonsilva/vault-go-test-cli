package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

func callHello() {
	httpClient := &http.Client{}
	url := os.Getenv("VAULT_ADDR")

	// create a vault client
	client, err := api.NewClient(&api.Config{Address: url, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}

	// the login path
	// this is configurable, change userpass to ldap etc
	path := "secret/sample/hello"

	// PUT call to get a token
	secret, err := client.Logical().Read(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(secret.Data["hello"])

}

func main() {

	for {
		callHello()
		time.Sleep(1 * time.Second)
	}

}
