package global

type Config struct {
	ServerPort uint64 `json:"serverPort"`
	//Db   struct {
	//	Db       string `json:"db"`
	//	Host     string `json:"host"`
	//	ReadHost string `json:"readHost"`
	//	User     string `json:"user"`
	//	Pass     string `json:"pass"`
	//} `json:"db"`
}

var GlobalConfig Config
