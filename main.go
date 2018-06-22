package envLoader

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

var envLoaded bool

//Load method Loads the variables defined in the specified file
//If no file was specified then by default it looks for the file named .env
func Load(customPath ...string) error {
	const (
		key = iota
		value
	)
	var path = ".env"
	if len(customPath) > 0 {
		path = customPath[0]
	}
	f, err := os.Open(path)
	if nil != err {
		return err
	}
	r := csv.NewReader(f)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if nil != err {
			return err
		}
		if len(line) < 2 {
			continue
		} else if strings.HasPrefix(strings.TrimSpace(line[key]), "#") {
			continue
		}
		if err := os.Setenv(strings.TrimSpace(line[key]), strings.TrimSpace(line[value])); nil != err {
			return err
		}
	}
	envLoaded = true
	return nil
}

//EnvOr checks given key (k) among environment variables and returns if it exists othervise returns the given default value (v)
func EnvOr(k, v string) string {
	if !envLoaded {
		_ = Load()
	}
	if val := os.Getenv(k); "" != val {
		return val
	}
	return v
}
