package time

import (
	"fmt"
	"time"
)

func Now() string {
	n := time.Now()

	nowString := fmt.Sprintf("Time now is : %v", n.Format(time.DateTime))

	return nowString
}
