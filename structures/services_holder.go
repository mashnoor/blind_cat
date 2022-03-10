package structures

type AppConfig struct {
	Services  []Service `yaml:"services"`
	SlackUrl  string    `yaml:"slackUrl"`
	RedisHost string    `yaml:"redisHost"`
	RedisPort string    `yaml:"redisPort"`
}
