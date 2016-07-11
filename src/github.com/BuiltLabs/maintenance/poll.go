package maintenance

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func (m *Maintenance) PollStatus() {
	for {
		m.checkStatus()
		time.Sleep(time.Second * 5)
	}
}

func (m *Maintenance) checkStatus() {
	var buffer bytes.Buffer
	buffer.WriteString("action:poll ** ")

	m.timestamp = time.Now()

	if err := m.checkFlags(); err != nil {
		panic(err)
	}

	if m.ready {
		if m.enabled {
			if err := m.createFile(); err != nil {
				panic(err)
			}

			buffer.WriteString("maintenance enabled")
		} else {
			if err := m.deleteFile(); err != nil {
				panic(err)
			}

			buffer.WriteString("maintenance disabled")
		}
	} else {
		buffer.WriteString("unknown error")
	}

	m.output(buffer.String())
}

func (m *Maintenance) checkFlags() (err error) {
	item, err := m.lookup()

	if err == nil {
		if item.Item["enabled"] != nil {
			m.ready = true
			m.enabled = *item.Item["enabled"].BOOL
		}

		if item.Item["meta"] != nil {
			m.MetaData = *item.Item["meta"].S
		} else {
			m.MetaData = ""
		}
	}

	return err
}

func (m *Maintenance) createFile() (err error) {
	err = ioutil.WriteFile(m.FileTarget, []byte(m.MetaData), 0644)

	if err != nil {
		m.output(fmt.Sprintf("File creation failed (reason: %s)", err))
		return err
	}

	return err
}

func (m *Maintenance) deleteFile() (err error) {
	if _, err := os.Stat(m.FileTarget); err != nil {
		return nil
	}

	err = os.Remove(m.FileTarget)

	return err
}
