package main

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/structures"
	"github.com/mashnoor/blind_cat/utility"
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

func updateToRedis(service *structures.Service, counter int) {
	utility.RedisHSet(service.Name, structures.ErrorCounter, counter)
}

func incrementErrorCounter(service *structures.Service) {
	currentErrorCounter := utility.RedisHGet(service.Name, structures.ErrorCounter)
	currentErrorCounter += 1
	updateToRedis(service, currentErrorCounter)
	fmt.Println(currentErrorCounter)

}

func checkHealth(service *structures.Service, wg *sync.WaitGroup) {
	for true {
		resp, err := requests.Get(service.Endpoint)
		check(err)
		if resp.R.StatusCode > 500 {
			incrementErrorCounter(service)

		}
		time.Sleep(time.Second * service.CheckInterval)
	}

	wg.Done()

}

func sendSlackMessage(msg string) {
	payload := make(map[string]string)
	payload["text"] = "Sup! Sending from Blind Cat :spaghetti:"
	jsonStr, err := json.Marshal(payload)

	hookUrl := "https://hooks.slack.com/services/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"
	resp, err := requests.PostJson(hookUrl, jsonStr)
	check(err)
	fmt.Println(resp.Text())
}

func main() {
	sendSlackMessage("hhh")
	//settings.InitRedis()
	//var wg sync.WaitGroup
	//
	//config := readConfigFile()
	//
	//monitorServices := structures.ServicesHolder{}
	//err := yaml.Unmarshal([]byte(config), &monitorServices)
	//check(err)
	//
	//for _, service := range monitorServices.Services {
	//	fmt.Println(service.Endpoint)
	//	go checkHealth(&service, &wg)
	//	wg.Add(1)
	//
	//}
	//
	//wg.Wait()

}
