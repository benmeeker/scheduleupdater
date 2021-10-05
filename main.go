package main

import (
	"flag"
	"os"
	"time"
)

// TODO: Create your own flag variables
var csvfile = flag.String("c", `WEEKLY SCHEDULE.csv`, "Path to file WEEKLY SCHEDULE")
var excelfile = flag.String("e", `Schedule.ods`, "Path to the file Schedule")

func main() {
	file, err := os.Stat(`Schedule.ods`)
	modifiedtime := file.ModTime()
	for {
		file, err = os.Stat(`Schedule.ods`)
		if err != nil {
		}
		if modifiedtime != file.ModTime() {
			modifiedtime = file.ModTime()
			flag.Parse()
			readFile()
			notify(*csvfile)
		}
		time.Sleep(30 * time.Minute)
	}
}
