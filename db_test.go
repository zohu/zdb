package zdb

import (
	"context"
	"testing"
	"time"
)

func TestNewDB(t *testing.T) {
	if err := Init(&Config{
		Kind:                      DatabaseType_DATABASE_POSTGRESQL,
		Host:                      "localhost",
		Port:                      "5432",
		Username:                  "postgres",
		Password:                  "JCFkQYex4f",
		Database:                  "test_db",
		Config:                    "sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdle:                   5,
		MaxAlive:                  50,
		MaxAliveMinutes:           60,
		SlowThresholdSecond:       5,
		SkipCallerLookup:          true,
		IgnoreRecordNotFoundError: true,
	}, TestTable{}); err != nil {
		t.Fatal(err)
	}
	if err := DB(context.TODO(), "").Create(&TestTable{
		TestString: "string",
		TestTime:   OrmTime(time.Now().Add(time.Hour * 24)),
		TestTime1:  time.Now().Add(time.Hour * 24),
	}).Error; err != nil {
		t.Fatal(err)
	}
	var tt TestTable
	DB(context.TODO(), "").First(&tt)
	t.Log(tt)
	t1 := tt.TestTime
	t.Log(t1)
	tz, _ := time.LoadLocation("Asia/Shanghai")
	t2 := tt.TestTime.In(tz)
	t.Log(t2)
	t.Log(t1.Time.Equal(t2))
	t.Log(t1.Before(t2))
}

type TestTable struct {
	TestId     int64     `json:"test_id" gorm:"primarykey;type:bigint;unsigned;autoIncrement"`
	TestString string    `json:"test_string" gorm:"comment:创建人"`
	TestTime   *Time     `json:"test_time" gorm:"type:timestamp(0) without time zone;autoCreateTime;comment:创建时间"`
	TestTime1  time.Time `json:"test_time_1" gorm:"type:timestamp(0) without time zone;autoCreateTime;comment:创建时间"`
}
