package env

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type App struct {
	ServiceName       string `json:"service_name"`
	Port              int    `json:"port"`
	PathLog           string `json:"path_log"`
	AllowedDomains    string `json:"allowed_domains"`
	RSAPrivateKey     string `json:"rsa_private_key"`
	LogReviewInterval int    `json:"log_review_interval"`
	LoggerHttp        bool   `json:"logger_http"`
	UrlPortal         string `json:"url_portal"`
}

type Ws struct {
	Url string `json:"url_api"`
	Token  string `json:"token"`
}

var (
	once   sync.Once
	config = &configuration{}
)

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

type configuration struct {
	App App `json:"app"`
	DB  DB  `json:"db"`
	Ws Ws `json:"ws"`
	Smtp     Smtp     `json:"smtp"`
	Template Template `json:"template"`
}

type Template struct {
	EmailCode        string `json:"email_code"`
	EmailToken       string `json:"email_change_password"`
}

type Smtp struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func fromFile() {
	once.Do(func() {
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}
	})
}

func NewConfiguration() *configuration {
	fromFile()
	return config
}
