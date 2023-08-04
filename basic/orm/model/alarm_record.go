package model

import "time"

type Alarm struct {
	ID       int
	Personid int
	Name     string
	Depid    int
	Depname  string
	Times    time.Time
	Type     int
	Mark     string
}

func (Alarm) TableName() string {
	return "alarm_record"
}
