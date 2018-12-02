package tool

import (
	"strconv"
	"time"
)

type TimeRecord struct {
	_start int
	_end int
}

func (this *TimeRecord) Start()  {
	this._start = time.Now().UTC().Nanosecond()
}

func (this *TimeRecord) End()  {
	this._end = time.Now().UTC().Nanosecond()
}

// MillSec 输出时间差值
// 根据Start 和 End 方法运行的结果计算出中间的时间差（单位:毫秒)
func (this *TimeRecord) MillSec() string  {
	var sec = (this._end - this._start) / 1e6
	return strconv.Itoa(sec) + "毫秒"
}


// MicroSec 输出时间差值
// 根据Start 和 End 方法运行的结果计算出中间的时间差（单位:微秒)
func (this *TimeRecord) MicroSec() string {
	var sec = (this._end - this._start) / 1e3
	return strconv.Itoa(sec) + "微秒"
}