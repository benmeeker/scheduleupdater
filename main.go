package main

import (
	"flag"
	"os"
	"time"
)

var csvfile = flag.String("c", `C:\Users\GROWJiffy\Desktop\WEEKLY SCHEDULE.csv`, "Path to file WEEKLY SCHEDULE")
var excelfile = flag.String("e", `C:\Users\GROWJiffy\Desktop\Schedule.ods`, "Path to the file Schedule")

func main() {
	file, err := os.Stat(`C:\Users\GROWJiffy\Desktop\Schedule.ods`)
	modifiedtime := file.ModTime()
	for {
		file, err = os.Stat(`C:\Users\GROWJiffy\Desktop\Schedule.ods`)
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
