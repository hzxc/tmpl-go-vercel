package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UnixTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05.000"
	zone        = "Asia/Shanghai"
)

type Model struct {
	ID        uint     `gorm:"primarykey" json:"id"`
	CreatedAt UnixTime `gorm:"type:datetime;" json:"created_at,omitempty"`
	UpdatedAt UnixTime `gorm:"type:datetime;" json:"updated_at,omitempty"`
	// tag json:"-" skip this filed
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UnmarshalJSON implements json unmarshal interface.
// func (t *UnixTime) UnmarshalJSON(data []byte) (err error) {
// 	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
// 	*t = UnixTime(now)
// 	return
// }

func (t *UnixTime) UnmarshalJSON(data []byte) (err error) {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = UnixTime(time.Unix(i, 0))
	return
}

// MarshalJSON implements json marshal interface.
// func (t UnixTime) MarshalJSON() ([]byte, error) {
// 	b := make([]byte, 0, len(timeFormart)+2)
// 	b = append(b, '"')
// 	b = time.Time(t).AppendFormat(b, timeFormart)
// 	b = append(b, '"')
// 	return b, nil
// }

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t UnixTime) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t UnixTime) Local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

// Value ...
func (t UnixTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *UnixTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UnixTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t UnixTime) Unix() int32 {
	return int32(time.Time(t).Unix())
}

// func (t *time.Time) UnmarshalJSON(data []byte) (err error) {
// 	i, err := strconv.ParseInt(string(data), 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	t.Time = time.Unix(i, 0)
// 	return
// }

// func (t *time.Time) MarshalJSON() ([]byte, error) {
// 	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
// }
