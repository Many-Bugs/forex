package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission checked the stat described error is Permission is denied - src: %s", src)
	}
	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir check the stat described error is the file does not exist - src: %s, err: %v", src, err)
	}
	if fileName != "" {
		f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("Fail to OpenFile :%v", err)
		}
		return f, nil
	}
	return nil, nil
}

func ReadMultipartfileToBuffer(file multipart.File) (buffer []byte, err error) {
	defer file.Close()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadOSFileToBuffer(file *os.File) (buffer []byte, err error) {
	defer file.Close()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	reader := bufio.NewReader(file)
	part := []byte{}
	for {
		if count, err := reader.Read(part); err != nil {
			break
		} else {
			buf.Write(part[:count])
		}
	}
	return buf.Bytes(), nil
}
