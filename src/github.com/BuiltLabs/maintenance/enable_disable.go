package maintenance

import "fmt"

func (m *Maintenance) EnableMaintenance() {
	if err := m.checkTable(); err != nil {
		m.Output.outputError(err, true)
	}

	if err := m.enableRecord(); err != nil {
		m.Output.outputError(err, true)
	}

	m.Output.output("maintenance enabled")
}

func (m *Maintenance) DisableMaintenance() {
	if err := m.checkTable(); err != nil {
		m.Output.outputError(err, true)
	}

	if err := m.disableRecord(); err != nil {
		m.Output.outputError(err, true)
	}

	m.Output.output("maintenance disabled")
}

func (m *Maintenance) checkTable() (err error) {
	// checks if table exists, if not, creates it
	if err := m.tableExists(); err != nil {
		m.Output.output(fmt.Sprintf("creating table [%s]", m.TableName))
		if err := m.createTable(); err != nil {
			return err
		}

		m.Output.output(fmt.Sprintf("waiting for table creation [%s]", m.TableName))
		if err := m.waitTable(); err != nil {
			return err
		}
	}

	return nil
}
