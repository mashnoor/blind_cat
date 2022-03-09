package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/structures"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
	"time"
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

func checkHealth(service structures.Service, wg *sync.WaitGroup) {
	for true {
		resp, err := requests.Get(service.Endpoint)
		check(err)
		fmt.Println(resp.R.StatusCode)

		time.Sleep(time.Second * 5)
	}

	wg.Done()

}

func main() {
	var wg sync.WaitGroup

	config := readConfigFile()

	monitorServices := structures.ServicesHolder{}
	err := yaml.Unmarshal([]byte(config), &monitorServices)
	check(err)

	for _, service := range monitorServices.Services {
		fmt.Println(service.Endpoint)
		go checkHealth(service, &wg)
		wg.Add(1)

	}

	wg.Wait()

}
