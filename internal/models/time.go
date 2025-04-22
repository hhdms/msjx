package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// JSONTime 自定义时间类型，用于处理JSON序列化和反序列化
type JSONTime time.Time

// 定义时间格式常量
const (
	TimeFormat     = "2006-01-02"
	TimeFormatFull = "2006-01-02 15:04:05"
)

// MarshalJSON 实现json.Marshaler接口
func (t JSONTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	
	// 根据时间是否包含时分秒信息选择不同的格式
	if tt.Hour() == 0 && tt.Minute() == 0 && tt.Second() == 0 {
		// 只有日期，没有时间部分
		stamp := fmt.Sprintf("\"%s\"", tt.Format(TimeFormat))
		return []byte(stamp), nil
	} else {
		// 包含时间部分
		stamp := fmt.Sprintf("\"%s\"", tt.Format(TimeFormatFull))
		return []byte(stamp), nil
	}
}

// UnmarshalJSON 实现json.Unmarshaler接口
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")

	if str == "" || str == "null" {
		*t = JSONTime(time.Time{})
		return nil
	}

	// 尝试解析日期格式
	parsed, err := time.Parse(TimeFormat, str)
	if err != nil {
		// 尝试解析日期时间格式
		parsed, err = time.Parse(TimeFormatFull, str)
		if err != nil {
			return err
		}
	}

	*t = JSONTime(parsed)
	return nil
}

// Value 实现driver.Valuer接口
func (t JSONTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

// Scan 实现sql.Scanner接口
func (t *JSONTime) Scan(value interface{}) error {
	if value == nil {
		*t = JSONTime(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*t = JSONTime(v)
	default:
		return fmt.Errorf("无法将 %v 转换为时间", value)
	}

	return nil
}

// String 返回时间的字符串表示
func (t JSONTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
