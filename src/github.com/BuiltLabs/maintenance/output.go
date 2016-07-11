package maintenance

import (
	"fmt"
	"time"
)

func (m *Maintenance) output(output string) {
	fmt.Printf("[%s] %s\n", m.timestamp.Format(time.RFC3339), output)
}
