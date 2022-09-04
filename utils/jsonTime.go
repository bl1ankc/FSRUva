package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	YYYYMMDDHHMISS = "2006-01-02 15:04:05" //常规类型
)

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(YYYYMMDDHHMISS))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*t).Time, err = time.ParseInLocation(`"`+YYYYMMDDHHMISS+`"`, string(data), time.Local)
	return err
}

type TestTime struct {
	Time JSONTime `json:"time"`
	Name string   `json:"name"`
}

//格式化输出
func (t *TestTime) ToString() string {
	return fmt.Sprintf("time:%+v,name:%+v", t.Time, t.Name)
}

func TestJsonMarshal(t *testing.T) {

	//测试序列化
	testMarshal := &TestTime{
		Time: JSONTime{time.Now()},
		Name: "序列化",
	}
	marshal, err := json.Marshal(testMarshal)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("序列化字符串：", string(marshal))

	//测试反序列化
	src := []byte(`{"time":"2020-05-26 20:20:44","name":"反序列化"}`)
	var testUnMarshal TestTime
	err = json.Unmarshal(src, &testUnMarshal)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("反序列化结构体：", testUnMarshal.ToString())
}
