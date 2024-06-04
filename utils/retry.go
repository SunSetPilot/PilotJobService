package utils

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	"PilotJobService/utils/log"
)

func getFunctionName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

func Retry[T any](maxTimes int, delay time.Duration, fn func() (T, error)) (T, error) {
	var (
		result T
		err    error
	)
	funcName := getFunctionName(fn)
	for i := -1; i < maxTimes; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if i != maxTimes-1 {
			time.Sleep(delay)
			log.Warnf("retry func %s", funcName)
		}
	}
	return result, err
}
