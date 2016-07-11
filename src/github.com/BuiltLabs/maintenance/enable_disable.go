package maintenance

func (m *Maintenance) EnableMaintenance() {
	m.checkTable()

	if err := m.enableRecord(); err != nil {
		panic(err)
	}

	m.output("action:set ** maintenance enabled")
}


func (m *Maintenance) DisableMaintenance() {
	m.checkTable()

	if err := m.disableRecord(); err != nil {
		panic(err)
	}

	m.output("action:set ** maintenance disabled")
}

func (m *Maintenance) checkTable() {
	// checks if table exists, if not, creates it
	if err := m.tableExists(); err != nil {
		if err := m.createTable(); err != nil {
			panic(err)
		}

		if err := m.waitTable(); err != nil {
			panic(err)
		}
	}
}