package util

import (
	"errors"
	"os"
	"path"
)

type DataType string

const (
	InvoiceDataType  DataType = "invoice"
	BusinessDataType DataType = "business"
)

func ensureDir(path string) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		if !stat.IsDir() {
			panic(errors.New(path + " is not a directory"))
		}
	}
}

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	billyHome := path.Join(home, ".billy")
	ensureDir(billyHome)
	return billyHome
}

func GetDataDir(dataType DataType) string {
	homeDir := GetHomeDir()
	dataDir := path.Join(homeDir, "data")
	ensureDir(dataDir)
	dataTypeDir := path.Join(dataDir, string(dataType))
	ensureDir(dataTypeDir)
	return dataTypeDir
}

func GetTemplateDir() string {
	homeDir := GetHomeDir()
	templateDir := path.Join(homeDir, "templates")
	ensureDir(templateDir)
	return templateDir
}

func OpenData(dataType DataType, name string) *os.File {
	fullPath := path.Join(GetDataDir(dataType), name+".json")
	file, err := os.Open(fullPath)
	PanicOnErrCloser(err, file)
	return file
}
