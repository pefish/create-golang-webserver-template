package global

type Config struct {
	Host string `json:"host"`
	Port uint64 `json:"port"`
}

var GlobalConfig Config
