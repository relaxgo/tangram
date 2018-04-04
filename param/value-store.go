package param

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

type valueStore struct {
	*http.Request
}

func NewValueStore(r *http.Request) ValueStore {
	return &valueStore{r}
}

func (r *valueStore) Value(key string) string {
	contentType := GetContentType(r.Request)
	if contentType == "application/json" {
		defer r.Body.Close()
		r1, r2, err := drainBody(r.Body)
		if err != nil {
			log.Print("read failed", err)
			return ""
		}
		r.Body = r2
		body, err := ioutil.ReadAll(r1)
		// TODO add new body reader ?
		if err != nil {
			log.Print("read body", err)
			return ""
		}
		value, _, _, err := jsonparser.Get(body, key)

		if err != nil {
			log.Println("jsonparser", err)
			return ""
		}
		if len(value) > 0 {
			return string(value)
		}
	}
	if v := r.FormValue(key); v != "" {
		return v
	}
	if v := r.URL.Query().Get(key); v != "" {
		return v
	}
	return ""
}

func GetContentType(r *http.Request) string {
	s := r.Header.Get("Content-Type")
	return strings.TrimSpace(strings.Split(s, ";")[0])
}

// copy from golang.org/src/net/http/httputil/dump.go
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
