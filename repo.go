package zdb

import "time"

type Empty struct{}

func OrmTime(t time.Time) *Time {
	return &Time{Time: t}
}

type BaseInfo struct {
	CreateBy string `json:"create_by" gorm:"comment:创建人"`
	UpdateBy string `json:"update_by" gorm:"comment:更新人"`
	CreateAt *Time  `json:"create_at" gorm:"type:timestamp(0);autoCreateTime;comment:创建时间"`
	UpdateAt *Time  `json:"update_at" gorm:"type:timestamp(0);autoUpdateTime;comment:更新时间"`
}
