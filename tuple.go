package clickhouse

import (
	"time"
)

func Tuple(v interface{}) interface{} {
	return v
}

func TupleFixedString(len int, v interface{}) interface{} {
	return v
}

func TupleDate(v []time.Time) interface{} {
	return v
}

func TupleDateTime(v []time.Time) interface{} {
	return v
}
