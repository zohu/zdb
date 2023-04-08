package zdb

import (
	"context"
	"github.com/zohu/zlog"
	"github.com/zohu/zutils"
	"gorm.io/gorm"
)

func Init(c *Config, dst ...interface{}) {
	conf = c
	db := DB(context.Background(), "")
	if len(dst) > 0 {
		if err := db.AutoMigrate(dst...); err != nil {
			zlog.Errorf("初始化表失败: %v", err)
		}
		zlog.Infof("初始化数据表 %d", len(dst))
	}
}

func DB(ctx context.Context, database string) *gorm.DB {
	def := zutils.FirstValue(database, conf.DataBase)
	if v, ok := conn.Load(def); ok {
		return v.(*Orm).db.WithContext(ctx)
	}
	orm := new(Orm)
	if err := orm.init(def); err != nil {
		zlog.Errorf("创建数据库链接失败: %v", err)
	}
	conn.Store(def, orm)
	return orm.db.WithContext(ctx)
}

func DBE(ctx context.Context, database, e string) *gorm.DB {
	return DB(ctx, database).Where("e = ?", e)
}
