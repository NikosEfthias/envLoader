package envLoader

import (
	"os"
	"encoding/csv"
	"io"
	"strings"
)

func Load(customPath ...string) error {
	const (
		key   = iota
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
		if err := os.Setenv(strings.TrimSpace(line[key]), strings.TrimSpace(line[value])); nil != err {
			return err
		}
	}
	return nil
}
