package zdb

import (
	"fmt"
)

var conf *Config

func (c *Config) Dsn(database string) (string, string, error) {
	switch c.GetKind() {
	case DatabaseType_DATABASE_MYSQL:
		uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/", c.GetUsername(), c.GetPassword(), c.GetHost(), c.GetPort())
		dsn := fmt.Sprintf("%s%s?%s", uri, database, c.GetConfig())
		return uri, dsn, nil
	case DatabaseType_DATABASE_POSTGRESQL:
		uri := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s %s",
			c.GetHost(),
			c.GetPort(),
			c.GetUsername(),
			c.GetPassword(),
			c.GetConfig(),
		)
		dsn := fmt.Sprintf("%s dbname=%s", uri, database)
		return uri, dsn, nil
	default:
		return "", "", fmt.Errorf("unsupported kind: %s", c.GetKind().String())
	}
}
