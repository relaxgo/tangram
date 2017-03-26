package param

import (
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type ValueStore interface {
	Value(string) string
}

func String(vs ValueStore, key string) string {
	return vs.Value(key)
}

func Bool(vs ValueStore, key string) bool {
	v, _ := strconv.ParseBool(String(vs, key))
	return v
}

func Int(vs ValueStore, key string) int {
	v, _ := strconv.ParseInt(String(vs, key), 10, 0)
	return int(v)
}

func Int8(vs ValueStore, key string) int8 {
	v, _ := strconv.ParseInt(String(vs, key), 10, 8)
	return int8(v)
}

func Int16(vs ValueStore, key string) int16 {
	v, _ := strconv.ParseInt(String(vs, key), 10, 16)
	return int16(v)
}
func Int32(vs ValueStore, key string) int32 {
	v, _ := strconv.ParseInt(String(vs, key), 10, 32)
	return int32(v)
}

func Int64(vs ValueStore, key string) int64 {
	v, _ := strconv.ParseInt(String(vs, key), 10, 64)
	return v
}

func Uint(vs ValueStore, key string) uint {
	v, _ := strconv.ParseUint(String(vs, key), 10, 0)
	return uint(v)
}

func Uint8(vs ValueStore, key string) uint8 {
	v, _ := strconv.ParseUint(String(vs, key), 10, 8)
	return uint8(v)
}

func Uint16(vs ValueStore, key string) uint16 {
	v, _ := strconv.ParseUint(String(vs, key), 10, 16)
	return uint16(v)
}
func Uint32(vs ValueStore, key string) uint32 {
	v, _ := strconv.ParseUint(String(vs, key), 10, 32)
	return uint32(v)
}

func Uint64(vs ValueStore, key string) uint64 {
	v, _ := strconv.ParseUint(String(vs, key), 10, 64)
	return v
}

func Byte(vs ValueStore, key string) byte {
	return byte(Uint8(vs, key))
}

func Float32(p ValueStore, key string) float32 {
	v, _ := strconv.ParseFloat(String(p, key), 32)
	return float32(v)
}

func Float64(p ValueStore, key string) float64 {
	v, _ := strconv.ParseFloat(String(p, key), 64)
	return v
}

type FileSize interface {
	Size() int64
}

func File(r *http.Request, key string) (file multipart.File, fsize int64) {
	file, _, err := r.FormFile(key)
	if err != nil {
		return nil, 0
	}
	switch f := file.(type) {
	case FileSize:
		fsize = f.Size()
	case *os.File:
		if s, err := f.Stat(); err == nil {
			fsize = s.Size()
		}
	}
	return
}
