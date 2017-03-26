package param

import (
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ValueStore interface {
	Value(string) string
}

func String(vs ValueStore, key string) string {
	return vs.Value(key)
}

func Int(vs ValueStore, key string) int {
	v, _ := strconv.Atoi(String(vs, key))
	return v
}

func Float64(p ValueStore, key string) float64 {
	v, _ := strconv.ParseFloat(String(p, key), 64)
	return v
}

func Bool(vs ValueStore, key string) bool {
	v := String(vs, key)
	return strings.ToLower(v) == "true"
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
