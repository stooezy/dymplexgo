package config

import (
	"time"

	"github.com/stooezy/dymplexgo/pkg/util"
)

type Database struct {
	Host            string
	Port            int
	Database        string
	Username        string
	Password        string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type HttpServer struct {
	Debug         bool
	ListenAddress string
	BaseURL       string
}

type Logger struct {
	Level string
}

type Server struct {
	Database   Database
	HttpServer HttpServer
	Logger     Logger
}

func DefaultServerConfig() Server {
	return Server{
		Database: Database{
			Host:            util.GetEnv("DB_HOST", "postgres"),
			Port:            util.GetEnvAsInt("DB_PORT", 5432),
			Database:        util.GetEnv("DB_DATABASE", "dymplex"),
			Username:        util.GetEnv("DB_USERNAME", "postgres"),
			Password:        util.GetEnv("DB_PASSWORD", ""),
			SSLMode:         util.GetEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns:    util.GetEnvAsInt("DB_MAX_OPEN_CONNS", 5),
			MaxIdleConns:    util.GetEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: time.Second * time.Duration(util.GetEnvAsInt("DB_CONN_MAX_LIFE_TIME_SEC", 60)),
		},
		HttpServer: HttpServer{
			Debug:         util.GetEnvAsBool("HTTP_SERVER_DEBUG", false),
			ListenAddress: util.GetEnv("HTTP_SERVER_LISTEN_ADDRESS", ":9000"),
			BaseURL:       util.GetEnv("HTTP_SERVER_BASE_URL", "http://localhost:9000"),
		},
		Logger: Logger{
			Level: util.GetEnv("LOGGER_LEVEL", "debug"),
		},
	}
}
