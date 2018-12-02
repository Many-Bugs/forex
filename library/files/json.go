package files

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JSONFileBinding struct {
	Dir string
}

func (j JSONFileBinding) BindFile(obj interface{}) error {
	f, err := Open(j.Dir, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return decodeJSON(f, obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
