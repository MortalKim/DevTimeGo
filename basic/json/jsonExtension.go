package json

import (
	jsoniter "github.com/json-iterator/go"
	"math"
	"time"
	"unsafe"
)

/**
 * @Author: Kim
 * @Description: A json extension of json-iterator, which can be used to parse time.Time
 * @File:  jsonExtension
 * @Date: 12/12/2022 11:02 PM
 */

// RegisterTimeDecoderFunc For timestamp that format is "1670312934.249"(sec.nsec), it can be parsed correctly.
func RegisterTimeDecoderFunc() {
	jsoniter.RegisterTypeDecoderFunc("time.Time", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		timestamp := iter.ReadFloat64()
		*((*time.Time)(ptr)) = time.Unix(int64(timestamp), int64(math.Mod(timestamp, 1)*1e9))
	})
}
