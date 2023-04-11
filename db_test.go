package zdb

import (
	"context"
	"testing"
	"time"
)

func TestNewDB(t *testing.T) {
	if err := Init(&Config{
		Kind:                      "postgres",
		Host:                      "localhost",
		Port:                      "5432",
		UserName:                  "postgres",
		Password:                  "JCFkQYex4f",
		DataBase:                  "test_db",
		Config:                    "sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdleConns:              5,
		MaxOpenConns:              50,
		ConnMaxLifeMinutes:        60,
		SlowThresholdSecond:       5,
		SkipCallerLookup:          true,
		IgnoreRecordNotFoundError: true,
	}, TestTable{}); err != nil {
		t.Fatal(err)
	}
	if err := DB(context.TODO(), "").Create(&TestTable{
		TestString: "string",
		TestTime:   OrmTime(time.Now().Add(time.Hour * 24)),
	}).Error; err != nil {
		t.Fatal(err)
	}
	DB(context.TODO(), "").First(&TestTable{})
}

type TestTable struct {
	TestId     int64  `json:"test_id" gorm:"primarykey;type:bigint;unsigned;autoIncrement"`
	TestString string `json:"test_string" gorm:"comment:创建人"`
	TestTime   *Time  `json:"test_time" gorm:"type:timestamp(0) without time zone;autoCreateTime;comment:创建时间"`
}
