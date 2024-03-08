package mexcgo

import "time"

type (
    TimeStamp int64
)

func (t TimeStamp) ToTime() time.Time {
    return time.UnixMilli(int64(t))
}
