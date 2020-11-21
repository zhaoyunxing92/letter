package global

import (
	"fmt"
	"time"
)

type JsonTime time.Time

//实现它的json序列化方法
func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
