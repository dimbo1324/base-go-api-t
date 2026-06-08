package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	EnvAddr           = "ADDR"
	EnvDBAddr         = "DB_ADDR"
	EnvDBMaxOpenConns = "DB_MAX_OPEN_CONNS"
	EnvDBMaxIdleConns = "DB_MAX_IDLE_CONNS"
	EnvDBMaxIdleTime  = "DB_MAX_IDLE_TIME"
)

const legacyEnvDBMaxIdleTime = "DB_MAX_IDLE_TIME_MINS"

const (
	defaultAddr              = ":8080"
	defaultDBAddr            = "postgres://postgres:password@localhost/appdb?sslmode=disable"
	defaultDBMaxOpenConns    = 30
	defaultDBMaxIdleConns    = 30
	defaultDBMaxIdleTime     = 15 * time.Minute
	defaultRequestTimeout    = 60 * time.Second
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 10 * time.Second
	defaultWriteTimeout      = 30 * time.Second
	defaultIdleTimeout       = time.Minute
	defaultShutdownTimeout   = 10 * time.Second
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Addr              string
	RequestTimeout    time.Duration
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  time.Duration
}

func Load() (Config, error) {
	maxOpenConns, err := getPositiveInt(EnvDBMaxOpenConns, defaultDBMaxOpenConns)
	if err != nil {
		return Config{}, err
	}

	maxIdleConns, err := getPositiveInt(EnvDBMaxIdleConns, defaultDBMaxIdleConns)
	if err != nil {
		return Config{}, err
	}

	maxIdleTime, err := getPositiveDurationWithFallback(EnvDBMaxIdleTime, legacyEnvDBMaxIdleTime, defaultDBMaxIdleTime)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Server: ServerConfig{
			Addr:              getString(EnvAddr, defaultAddr),
			RequestTimeout:    defaultRequestTimeout,
			ReadHeaderTimeout: defaultReadHeaderTimeout,
			ReadTimeout:       defaultReadTimeout,
			WriteTimeout:      defaultWriteTimeout,
			IdleTimeout:       defaultIdleTimeout,
			ShutdownTimeout:   defaultShutdownTimeout,
		},
		DB: DBConfig{
			Addr:         getString(EnvDBAddr, defaultDBAddr),
			MaxOpenConns: maxOpenConns,
			MaxIdleConns: maxIdleConns,
			MaxIdleTime:  maxIdleTime,
		},
	}, nil
}

func getString(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getPositiveInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", key)
	}
	if parsed <= 0 {
		return 0, fmt.Errorf("%s must be greater than zero", key)
	}

	return parsed, nil
}

func getPositiveDurationWithFallback(primaryKey, fallbackKey string, fallback time.Duration) (time.Duration, error) {
	key := primaryKey
	value := os.Getenv(primaryKey)
	if value == "" {
		key = fallbackKey
		value = os.Getenv(fallbackKey)
	}
	if value == "" {
		return fallback, nil
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid duration", key)
	}
	if parsed <= 0 {
		return 0, fmt.Errorf("%s must be greater than zero", key)
	}

	return parsed, nil
}
