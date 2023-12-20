package global

type Config struct {
	Host string `json:"host"`
	Port uint64 `json:"port"`
	Db   struct {
		Db       string `json:"db"`
		Host     string `json:"host"`
		ReadHost string `json:"readHost"`
		User     string `json:"user"`
		Pass     string `json:"pass"`
	} `json:"db"`
}

var GlobalConfig Config
