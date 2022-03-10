package structures

import "time"

const (
	GET  string = "GET"
	POST        = "POST"
)

type Service struct {
	Name                         string        `yaml:"name"`
	Endpoint                     string        `yaml:"endpoint"`
	Method                       string        `yaml:"method"`
	Body                         string        `yaml:"body"`
	MaxErrorCount                int64         `yaml:"MaxErrorCount"`
	CheckInterval                time.Duration `yaml:"checkInterval"`
	ConsecutiveNotificationDelay int64         `yaml:"consecutiveNotificationDelay"`
	//SlackUrl                     string        `yaml:"slackUrl"`
	Timeout int `yaml:"timeout"`
}
