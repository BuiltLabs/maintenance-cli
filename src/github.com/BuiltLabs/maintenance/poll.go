package maintenance

import (
	"errors"
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

func (m *Maintenance) checkStatus() (err error) {
	m.timestamp = time.Now()
	err = m.checkFlags()

	if m.ready {
		if m.enabled {
			if err = m.createFile(); err == nil {
				m.Output.output("maintenance enabled")

				return
			}
		} else {
			if err = m.deleteFile(); err == nil {
				m.Output.output("maintenance disabled")

				return
			}
		}
	}

	m.Output.outputError(err, false)

	return
}

func (m *Maintenance) checkFlags() (err error) {
	item, err := m.lookup()

	if err == nil {
		if item.Item["enabled"] != nil {
			m.ready = true
			m.enabled = *item.Item["enabled"].BOOL
		} else {
			return errors.New(fmt.Sprintf("Record lookup failed (reason: could not find [%s: %s])", m.KeyName, m.Key))
		}

		if item.Item["meta"] != nil {
			m.MetaData = *item.Item["meta"].S
		} else {
			m.MetaData = ""
		}
	} else {
		return errors.New(fmt.Sprintf("Record lookup failed (reason: %s)", err))
	}

	return nil
}

func (m *Maintenance) createFile() error {
	if err := ioutil.WriteFile(m.FileTarget, []byte(m.MetaData), 0644); err != nil {
		return errors.New(fmt.Sprintf("File creation failed (reason: %s)", err))
	}

	return nil
}

func (m *Maintenance) deleteFile() error {
	if _, err := os.Stat(m.FileTarget); err != nil {
		// file doesn't exist, nothing to delete
		return nil
	}

	if err := os.Remove(m.FileTarget); err != nil {
		return errors.New(fmt.Sprintf("File deletion failed (reason: %s", err))
	}

	return nil
}
