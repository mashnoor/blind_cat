package core

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/mashnoor/blind_cat/structures"
	"github.com/mashnoor/blind_cat/utility"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func updateToRedis(service *structures.Service, counter int64) {
	utility.RedisHSet(service.Name, structures.ErrorCounter, counter)
}

func incrementErrorCounter(service *structures.Service) {
	//log.WithFields(log.Fields{
	//	"service_name":     service.Name,
	//	"service_endpoint": service.Endpoint,
	//}).Info("counter increased")

	currentErrorCounter := utility.RedisHGet(service.Name, structures.ErrorCounter)
	currentErrorCounter += 1
	updateToRedis(service, currentErrorCounter)
	//fmt.Println(currentErrorCounter)

}

func getErrorCounter(service *structures.Service) int64 {
	return utility.RedisHGet(service.Name, structures.ErrorCounter)
}

func resetErrorCounter(service *structures.Service) {
	//log.WithFields(log.Fields{
	//	"service_name": service.Name,
	//}).Info("Error counter reset after successful reach")
	updateToRedis(service, 0)

}

func getLastNotificationSentTime(service *structures.Service) int64 {
	return utility.RedisHGet(service.Name, structures.LastNotificationSent)
}

func updateLastNotificationSentTime(service *structures.Service) {
	log.WithFields(log.Fields{
		"service_name": service.Name,
	}).Info("last notification send time updated")
	utility.RedisHSet(service.Name, structures.LastNotificationSent, time.Now().Unix())
}

func notificationDecision(service *structures.Service) {
	elapsedTime := time.Now().Unix() - getLastNotificationSentTime(service)
	timeCrossedThreshold := false
	if elapsedTime >= service.ConsecutiveNotificationDelay {
		timeCrossedThreshold = true
	}
	errorCrossedThreshold := getErrorCounter(service) > service.MaxErrorCount
	if errorCrossedThreshold && timeCrossedThreshold {
		//utility.SendSlackMessage(service.Name, true, getErrorCounter(service))
		fmt.Println("----- SEND DOWN NOTIFICATION ----------")
		updateLastNotificationSentTime(service)
	}
}

func decideUpMsg(service *structures.Service) {
	if getErrorCounter(service) > 0 {
		fmt.Println("------ SENDING UP NOTIFICATION -------------")
		//utility.SendSlackMessage(service.Name, false, 0)
	}
}

func isFailedRequest(service *structures.Service) bool {
	if service.Method == structures.GET {
		resp, err := requests.Get(service.Endpoint)
		if err != nil || resp.R.StatusCode > 400 {
			return true
		}

		return false
	} else {
		resp, err := requests.PostJson(service.Endpoint, service.JsonBody)
		if err != nil || resp.R.StatusCode > 400 {

			return true
		}

	}

	return false
}

func checkHealth(service *structures.Service, wg *sync.WaitGroup) {
	fmt.Println(service.ConsecutiveNotificationDelay)
	for true {
		log.WithFields(log.Fields{
			"name":              service.Name,
			"error_count":       getErrorCounter(service),
			"last_notification": time.Unix(getLastNotificationSentTime(service), 0),
		}).Info("")

		if isFailedRequest(service) {
			log.WithFields(log.Fields{
				"service_name": service.Name,
			}).Error("Service is down")
			incrementErrorCounter(service)
			notificationDecision(service)
		} else {
			decideUpMsg(service)
			resetErrorCounter(service)
			log.WithFields(log.Fields{
				"name": service.Name,
			}).Info("Service is up")
		}
		time.Sleep(time.Second * service.CheckInterval)
	}

	wg.Done()

}
