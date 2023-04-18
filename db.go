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
		return fmt.Errorf("数据库初始化失败")
	}
	if len(dst) > 0 {
		if err := db.AutoMigrate(dst...); err != nil {
			return fmt.Errorf("初始化表失败: %v", err)
		}
		zlog.Infof("初始化数据表 %d", len(dst))
	}
	return nil
}

func DB(ctx context.Context, database string) *gorm.DB {
	def := zutils.FirstTruthString(database, conf.DataBase)
	if v, ok := conn.Load(def); ok {
		return v.(*Orm).db.WithContext(ctx)
	}
	orm := new(Orm)
	if err := orm.init(def); err != nil {
		zlog.Errorf("创建数据库链接失败 %v", err)
		return nil
	}
	conn.Store(def, orm)
	return orm.db.WithContext(ctx)
}

func DBE(ctx context.Context, database, e string) *gorm.DB {
	return DB(ctx, database).Where("e = ?", e)
}
