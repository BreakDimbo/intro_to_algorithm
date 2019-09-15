package util

import (
	"fmt"
	"time"
)

func ConsumeTime(t time.Time) {
	fmt.Printf("time consumed is %fs\n", time.Now().Sub(t).Seconds())
}
