package utils

import (
	"fmt"
	"time"
)

var Password = fmt.Sprintf("%04d", time.Now().Nanosecond()/1000%1000)
