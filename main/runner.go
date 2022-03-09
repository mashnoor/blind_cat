package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/structures"
	"gopkg.in/yaml.v2"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readConfigFile() string {
	dat, err := os.ReadFile("services.yaml")
	check(err)

	return string(dat)
}

func checkHealth() {

}

func main() {

	config := readConfigFile()

	monitorServices := structures.ServicesHolder{}
	err := yaml.Unmarshal([]byte(config), &monitorServices)
	check(err)

	for _, service := range monitorServices.Services {
		if service.Method == structures.GET {
			resp, err := requests.Get(service.Endpoint)
			check(err)
			fmt.Println(resp.R.StatusCode)
		}

	}

}
