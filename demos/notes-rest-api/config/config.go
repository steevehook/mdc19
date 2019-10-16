package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	listen = "listen"

	dbURL                     = "db.url"
	dbMaxOpenConnections      = "db.max_open_connections"
	dbMaxIdleConnections      = "db.max_idle_connections"
	dbConnMaxLifetime         = "db.conn_max_lifetime"
	dbEnablePreparedStmtCache = "db.enable_prepared_stmt_cache"
)

type Manager struct{}

func New() (Manager, error) {
	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Manager{}, err
	}
	return Manager{}, nil
}

func (m Manager) Listen() string {
	return viper.GetString(listen)
}

func (m Manager) DBUrl() string {
	return viper.GetString(dbURL)
}

func (m Manager) DBMaxOpenConnections() int {
	return viper.GetInt(dbMaxOpenConnections)
}

func (m Manager) DBMaxIdleConnections() int {
	return viper.GetInt(dbMaxIdleConnections)
}

func (m Manager) DBConnMaxLifetime() time.Duration {
	return viper.GetDuration(dbConnMaxLifetime)
}

func (m Manager) DBEnablePreparedStmtCache() bool {
	return viper.GetBool(dbEnablePreparedStmtCache)
}
