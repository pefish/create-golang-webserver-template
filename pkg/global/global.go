package global

type Config struct {
	ServerHost string `json:"server-host" default:"0.0.0.0" usage:"Server host."`
	ServerPort int    `json:"server-port" default:"8000" usage:"Server port."`
	// DbHost     string `json:"db-host" default:"mysql" usage:"Database host."`
	// DbPort     int    `json:"db-port" default:"3306" usage:"Database port."`
	// DbDb       string `json:"db-db" default:"" usage:"Database to connect."`
	// DbUser     string `json:"db-user" default:"admin" usage:"Username to connect database."`
	// DbPass     string `json:"db-pass" default:"" usage:"Password to connect database."`
}

var GlobalConfig Config
