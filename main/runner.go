package main

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/settings"
	"github.com/mashnoor/blind_cat/structures"
	"github.com/mashnoor/blind_cat/utility"
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

func updateToRedis(service *structures.Service, counter int64) {
	utility.RedisHSet(service.Name, structures.ErrorCounter, counter)
}

func incrementErrorCounter(service *structures.Service) {
	currentErrorCounter := utility.RedisHGet(service.Name, structures.ErrorCounter)
	currentErrorCounter += 1
	updateToRedis(service, currentErrorCounter)
	fmt.Println(currentErrorCounter)

}

func getErrorCounter(service *structures.Service) int64 {
	return utility.RedisHGet(service.Name, structures.ErrorCounter)
}

func resetErrorCounter(service *structures.Service) {
	updateToRedis(service, 0)

}

func getLastNotificationSentTime(service *structures.Service) int64 {
	return utility.RedisHGet(service.Name, structures.LastNotificationSent)
}

func notificationDecision(service *structures.Service) {
	elapsedTime := time.Now().Unix() - getLastNotificationSentTime(service)
	timeCrossedThreshold := false
	if elapsedTime >= service.ConsecutiveNotificationDelay {
		timeCrossedThreshold = true
	}
	errorCrossedThreshold := getErrorCounter(service) > service.MaxErrorCount
	if errorCrossedThreshold && timeCrossedThreshold {
		//utility.SendSlackMessage(service.Name, true)
		fmt.Println("Time to send notification")
	}
}

func checkHealth(service *structures.Service, wg *sync.WaitGroup) {
	for true {
		resp, err := requests.Get(service.Endpoint)
		check(err)
		if resp.R.StatusCode > 500 {
			incrementErrorCounter(service)
			notificationDecision(service)

		} else {
			resetErrorCounter(service)
		}
		time.Sleep(time.Second * service.CheckInterval)
	}

	wg.Done()

}

func main() {
	//sendSlackMessage("breb_postpaid")
	settings.InitRedis()
	var wg sync.WaitGroup

	config := readConfigFile()

	monitorServices := structures.ServicesHolder{}
	err := yaml.Unmarshal([]byte(config), &monitorServices)
	check(err)

	for _, service := range monitorServices.Services {
		fmt.Println(service.Endpoint)
		go checkHealth(&service, &wg)
		wg.Add(1)

	}

	wg.Wait()

}
