package zdb

import (
	"fmt"
)

type DatabaseType string

const (
	DatabaseTypeMysql    DatabaseType = "mysql"
	DatabaseTypePostgres              = "postgres"
)

var conf *Config

type Config struct {
	Kind                      DatabaseType `yaml:"kind"`
	Host                      string       `yaml:"host"`
	Port                      string       `yaml:"port"`
	UserName                  string       `yaml:"username"`
	Password                  string       `yaml:"password"`
	DataBase                  string       `yaml:"database"`
	Config                    string       `yaml:"config"`
	MaxIdleConns              int          `yaml:"max_idle_conns"`
	MaxOpenConns              int          `yaml:"max_open_conns"`
	ConnMaxLifeMinutes        int          `yaml:"conn_max_life_minutes"`
	SlowThresholdSecond       int          `yaml:"slow_threshold_second"`
	SkipCallerLookup          bool         `yaml:"skip_caller_lookup"`
	IgnoreRecordNotFoundError bool         `yaml:"ignore_record_not_found_error"`
}

func (c *Config) Dsn(database string) (string, string, error) {
	switch c.Kind {
	case DatabaseTypeMysql:
		uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/", c.UserName, c.Password, c.Host, c.Port)
		dsn := fmt.Sprintf("%s%s?%s", uri, database, c.Config)
		return uri, dsn, nil
	case DatabaseTypePostgres:
		uri := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s %s",
			c.Host,
			c.Port,
			c.UserName,
			c.Password,
			c.Config,
		)
		dsn := fmt.Sprintf("%s dbname=%s", uri, database)
		return uri, dsn, nil
	default:
		return "", "", fmt.Errorf("unsupported kind: %s", c.Kind)
	}
}
