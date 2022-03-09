package structures

const (
	GET  string = "GET"
	POST        = "POST"
)

type Service struct {
	Name                         string `yaml:"name"`
	Endpoint                     string `yaml:"endpoint"`
	Method                       string `yaml:"method"`
	Body                         string `yaml:"body"`
	MaxErrorCount                int    `yaml:"MaxErrorCount"`
	CheckInterval                uint64 `yaml:"checkInterval"`
	ConsecutiveNotificationDelay int    `yaml:"consecutiveNotificationDelay"`
	SlackUrl                     string `yaml:"slackUrl"`
	Timeout                      int    `yaml:"timeout"`
}
