package maintenance

import (
	"os"
	"time"
	"io/ioutil"
)

type Maintenance struct {
	ready       bool
	enabled     bool
	timestamp   time.Time

	FileTarget  string
	TableName   string
	Key         string
	KeyName     string
	MetaData    string
}

func NewMaintenance() *Maintenance {
	return &Maintenance{
		timestamp: time.Now(),
	}
}

func (m *Maintenance) ImportMetaData(filePath string) {
	// check if file exists
	if _, err := os.Stat(filePath); err != nil {
		return
	}

	if data, err := ioutil.ReadFile(filePath); err == nil {
		m.MetaData = string(data)
	}
}