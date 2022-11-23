package junk

import (
	"fmt"
	"time"
)

func init() {

}

func Junk() string {
	currentTime := time.Now()
	return fmt.Sprintf("Junk: %s", currentTime)
}
