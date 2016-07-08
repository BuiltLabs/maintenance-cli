package main

import (
	"flag"
	"errors"
	"fmt"
	"os"
)

var enableMaintenance bool
var disableMaintenance bool
var fileTarget string
var engine string
var key string
var table string
var dataColumn string
var keyColumn string

func init() {
	flag.BoolVar(&enableMaintenance, "enable", false, "disregard aws, turns on maintenance")
	flag.BoolVar(&disableMaintenance, "disable", false, "disregard aws, turns off maintenance")
	flag.StringVar(&fileTarget, "fileTarget", "./maintenance.enable", "the file which is to be added/removed")
	flag.StringVar(&engine, "engine", "dynamodb", "The data store model used. DynamoDB is the only option right now.")
	flag.StringVar(&key, "key", "dev", "Key Value of the data record to be referenced")
	flag.StringVar(&table, "table", "maintenanceFlags", "The table which our key/data is stored in")
	flag.StringVar(&dataColumn, "dataColumn", "data", "The column which contains the data object")
	flag.StringVar(&keyColumn, "keyColumn", "key", "The column which contains our key indeces")
}

func main() {
	flag.Parse()

	if err := checkFlags(); err != nil {
		panic(err)
	}

	if true == enableMaintenance {
		enable()
	}

	if true == disableMaintenance {
		disable()
	}
}

func checkFlags() (error) {
	if true == enableMaintenance && true == disableMaintenance {
		return errors.New("One does not simply enable and disable maintenance at the same time.")
	}

	return nil;
}

func enable() {
	fmt.Println("Enabling maintenance mode...")

	file, err := os.Create(fileTarget);

	if err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}
}

func disable() {
	fmt.Println("Disabling maintenance mode...")

	if err := os.Remove(fileTarget); err != nil {
		panic(err)
	}
}