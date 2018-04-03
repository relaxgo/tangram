package param

import (
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
		body, err := ioutil.ReadAll(r.Body)
		// TODO add new body reader ?
		if err != nil {
			log.Print("read body", err)
		}
		value, _, _, err := jsonparser.Get(body, key)

		if err != nil {
			log.Println("jsonparser", err)
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
