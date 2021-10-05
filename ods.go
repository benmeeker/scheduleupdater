package main

import (
	"bufio"
	"fmt"
	"os"

	ods "github.com/LIJUCHACKO/ods2csv"
)

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line+"\r")
	}
	return w.Flush()
}

func readFile() {
	// TODO: Put in your filepath
	Filecontent, eerr := ods.ReadODSFile(`filepath`)
	if eerr != nil {
		fmt.Printf("Read : %s\n", eerr)
		var yes string
		fmt.Scan(&yes)
		os.Exit(0)
	}
	for _, sheet := range Filecontent.Sheets {
		if sheet.Name != "WEEKLY SCHEDULE" {
			continue
		}
		outputcontent := []string{}
		var decide bool
		for _, row := range sheet.Rows {
			rowString := ""
			for _, cell := range row.Cells {
				if cell.Text == "Name" {
					decide = true
					break
				}
				if cell.Text == "Scheduled Hours" {
					decide = false
					break
				}
				if cell.Text != "-" {
					rowString = rowString + cell.Text + ","
				}
			}
			if decide {
				if len(outputcontent) == 0 {
					rowString = "Name, Mon. Start, Mon. End,, Mon. Hrs, Tues. Start, Tues. End,, Tues. Hrs, Wed. Start, Wed. End,, Wed. Hours, Thurs. Start, Thurs. End,, Thurs. Hrs, Fri. Start, Fri. End,, Fri. Hrs, Sat. Start, Sat End,, Sat Hrs,,,,, Total Hrs"
				}
				outputcontent = append(outputcontent, rowString)
			}
		}
		fmt.Printf("writing   %s", *csvfile)
		err := writeLines(outputcontent, *csvfile)
		if err != nil {
			fmt.Printf("writing: %s", err)
			var yes string
			fmt.Scan(&yes)
			return
		}
	}
}
