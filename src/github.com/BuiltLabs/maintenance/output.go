package maintenance

import (
	"fmt"
	"os"
	"time"
)

type Output struct {
	Verb string
	Noun string

	timestamp time.Time
}

func (o *Output) output(output string) {
	fmt.Printf("[%s] %s:%s ** %s\n", o.timestamp, o.Verb, o.Noun, output)
}

func (o *Output) outputError(err error, fatal bool) {
	o.Verb = "error"
	o.output(fmt.Sprint(err))

	if fatal {
		os.Exit(1)
	}
}
