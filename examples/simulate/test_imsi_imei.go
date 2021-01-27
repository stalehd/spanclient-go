package main

import (
	"fmt"
	"sync/atomic"
)

const firstIMEI = int64(888888880000000)
const firstIMSI = int64(999990000000000)

var imsiCounter *int64
var imeiCounter *int64

func init() {
	imsiCounter = new(int64)
	imeiCounter = new(int64)
	atomic.StoreInt64(imsiCounter, firstIMSI)
	atomic.StoreInt64(imeiCounter, firstIMEI)
}

func setIMSIOffset(offset int64) {
	atomic.AddInt64(imsiCounter, offset)
	atomic.AddInt64(imeiCounter, offset)
}

func nextIMSI() string {
	return fmt.Sprintf("%d", atomic.AddInt64(imsiCounter, 1))
}

func nextIMEI() string {
	return fmt.Sprintf("%d", atomic.AddInt64(imeiCounter, 1))
}
