package maintenance

import (
	"bytes"
	"os"
	"time"
	"fmt"
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
			if err := m.enable(); err != nil {
				panic(err)
			}

			buffer.WriteString("maintenance enabled")
		} else {
			if err := m.disable(); err != nil {
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

	if err == nil && item.Item["enabled"] != nil {
		m.ready     = true
		m.enabled   = *item.Item["enabled"].BOOL
	}

	return err
}


func (m Maintenance) enable() (err error) {
	file, err := os.Create(m.FileTarget)

	if err != nil {
		m.output(fmt.Sprintf("File creation failed (reason: %s)", err))
		return err
	}

	if err := file.Close(); err != nil {
		m.output(fmt.Sprintf("File close failed (reason: %s)", err))
		return err
	}

	return err
}

func (m Maintenance) disable() (err error) {
	if _, err := os.Stat(m.FileTarget); err != nil {
		return nil
	}

	err = os.Remove(m.FileTarget)

	return err
}
