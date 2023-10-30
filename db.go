package zdb

import (
	"context"
	"fmt"
	"github.com/zohu/zlog"
	"github.com/zohu/zutils"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func Init(c *Config, dst ...interface{}) error {
	conf = c
	db := DB(context.Background(), "")
	if db == nil {
		return fmt.Errorf(">>>>> init db failed")
	}
	if dst != nil && len(dst) > 0 {
		if err := db.AutoMigrate(dst...); err != nil {
			return fmt.Errorf(">>>>> init db table failed: %v", err)
		}
		zlog.Infof(">>>>> init db table success: %d", len(dst))
	}
	return nil
}

func DB(ctx context.Context, args ...string) *gorm.DB {
	if len(args) == 0 {
		args = append(args, "")
	}
	def := zutils.FirstTruthString(args[0], conf.DataBase)
	if v, ok := conn.Load(def); ok {
		return v.(*Orm).db.WithContext(ctx)
	}
	orm := new(Orm)
	if err := orm.init(def); err != nil {
		zlog.Errorf(">>>>> init db conn failed %v", err)
		return nil
	}
	conn.Store(def, orm)
	return orm.db.WithContext(ctx)
}
