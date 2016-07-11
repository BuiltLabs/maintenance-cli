package maintenance

import (
	"io/ioutil"
	"os"
	"time"
)

type Maintenance struct {
	ready     bool
	enabled   bool
	timestamp time.Time
	Output    Output

	FileTarget string
	TableName  string
	Key        string
	KeyName    string
	MetaData   string
}

func NewMaintenance() *Maintenance {
	ts := time.Now()

	o := &Output{
		timestamp: ts,
	}

	m := &Maintenance{
		timestamp: ts,
		Output:    *o,
	}

	return m
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
