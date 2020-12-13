package LoadConfig

type Config struct{
	RootPath string `json:"root"`
	DataBase string `json:"database"`
	Listen string `json:"listen"`
	Port int `json:"port"`
	MQTT string `json:"mqtt"`
	TSDB string `json:"TSDB"`
	AppID string `json:"appid"`
	Secret string `json:"secret"`
}