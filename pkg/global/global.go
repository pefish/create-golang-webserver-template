package global

type Config struct {
	ServerHost string `json:"server-host"`
	ServerPort uint64 `json:"server-port"`
	//Db   struct {
	//	Db       string `json:"db"`
	//	Host     string `json:"host"`
	//	ReadHost string `json:"readHost"`
	//	User     string `json:"user"`
	//	Pass     string `json:"pass"`
	//} `json:"db"`
}

var GlobalConfig Config
