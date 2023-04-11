package zdb

import "time"

type Empty struct{}

func OrmTime(t time.Time) *Time {
	return &Time{Time: t}
}

type ID struct {
	Id int64 `json:"id" gorm:"primarykey;type:bigint;unsigned;autoIncrement"`
}

type Model struct {
	CreateBy   string `json:"create_by" gorm:"comment:创建人"`
	UpdateBy   string `json:"update_by" gorm:"comment:更新人"`
	CreateTime *Time  `json:"create_time" gorm:"type:timestamp(0) without time zone;autoCreateTime;comment:创建时间"`
	UpdateTime *Time  `json:"update_time" gorm:"type:timestamp(0) without time zone;autoUpdateTime;comment:更新时间"`
}
