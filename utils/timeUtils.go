package utils

import "time"

/**
 * @Author: Kim
 * @Description:
 * @File:  timeUtils
 * @Date: 12/13/2022 1:39 PM
 */

// GetDayTimestamp get timestamp of day 0:00:00 to 23:59:59
func GetDayTimestamp(now time.Time) (int64, int64) {
	year, month, day := now.Date()
	//get day 0:00:00
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	//get day 23:59:59
	end := time.Date(year, month, day, 23, 59, 59, 0, time.Local).Unix()
	return start, end
}

func GetDayTimeString(now time.Time) (string, string) {
	year, month, day := now.Date()
	//get day 0:00:00
	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
	//get day 23:59:59
	end := time.Date(year, month, day, 23, 59, 59, 0, time.Local).Format("2006-01-02 15:04:05")
	return start, end
}
