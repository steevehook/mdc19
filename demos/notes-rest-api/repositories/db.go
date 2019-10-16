package repositories

import (
	"time"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

type DBOptions struct {
	URL                     string
	MaxOpenConnections      int
	MaxIdleConnections      int
	ConnectionMaxLifetime   time.Duration
	EnablePreparedStmtCache bool
}

func DBConnection(options DBOptions) (sqlbuilder.Database, error, ) {
	url, err := mysql.ParseURL(options.URL)
	if err != nil {
		return nil, err
	}

	conn, err := mysql.Open(url)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(options.MaxOpenConnections)
	conn.SetMaxIdleConns(options.MaxIdleConnections)
	conn.SetConnMaxLifetime(options.ConnectionMaxLifetime)
	conn.SetPreparedStatementCache(options.EnablePreparedStmtCache)
	return conn, nil
}
