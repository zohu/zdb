package zdb

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zohu/zlog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var conn sync.Map

type Orm struct {
	db *gorm.DB
}

func (o *Orm) init(database string) error {
	uri, dia, err := o.open(database)
	if err != nil {
		return err
	}
	err = o.sync(uri, conf.Kind, database)
	if err != nil {
		return err
	}
	if orm, err := gorm.Open(dia, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭复数表名
		},
		Logger: NewLoggerForGorm(&OptionForGorm{
			SlowThreshold:             time.Second * time.Duration(conf.SlowThresholdSecond),
			SkipCallerLookup:          conf.SkipCallerLookup,
			IgnoreRecordNotFoundError: conf.IgnoreRecordNotFoundError,
		}),
	}); err != nil {
		return fmt.Errorf("链接数据库失败 %s", err.Error())
	} else {
		d, _ := orm.DB()
		d.SetMaxIdleConns(conf.MaxIdleConns)
		d.SetMaxOpenConns(conf.MaxOpenConns)
		d.SetConnMaxLifetime(time.Minute * time.Duration(conf.ConnMaxLifeMinutes))
		o.db = orm
	}
	return nil
}

func (o *Orm) open(database string) (string, gorm.Dialector, error) {
	uri, dsn, err := conf.Dsn(database)
	if err != nil {
		return "", nil, err
	}
	switch conf.Kind {
	case "mysql":
		return uri, mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   false, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: true,  // 根据当前 MySQL 版本自动配置
		}), nil
	case "postgres":
		return uri, postgres.Open(dsn), nil
	default:
		return "", nil, fmt.Errorf("不支持的数据库类型: %s", conf.Kind)
	}
}

func (o *Orm) exec(dsn, driver, str string) error {
	if db, err := sql.Open(driver, dsn); err != nil {
		return err
	} else {
		defer db.Close()
		if err = db.Ping(); err != nil {
			return err
		}
		_, err = db.Exec(str)
		return err
	}
}

func (o *Orm) sql(kind, database string) string {
	switch kind {
	case "mysql":
		return fmt.Sprintf(
			"CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;",
			database,
		)
	case "postgres":
		return fmt.Sprintf(
			"SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')",
			database, database,
		)
	default:
		return ""
	}
}

func (o *Orm) sync(dsn, kind, database string) error {
	if err := o.exec(dsn, kind, o.sql(kind, database)); err != nil {
		return fmt.Errorf(">>>>> init db failed: %w", err)
	}
	zlog.Infof(">>>>> init db success: %s", database)
	return nil
}

func (o *Orm) dbc(ctx context.Context) *gorm.DB {
	return o.db.WithContext(ctx)
}
