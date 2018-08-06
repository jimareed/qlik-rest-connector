package main

import (
	"io"
	"net/http"

	"github.com/tealeg/xlsx"
)

const MAX_COLS = 100

// ItemList : list of items
type ItemList []struct {
	Item string `json:"item"`
}

func main() {
	http.HandleFunc("/data", data)
	http.HandleFunc("/", health)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

func emptyrow(colvalue [MAX_COLS]string, numCols int) bool {
	for i, text := range colvalue {
		if i < numCols {
			if text != "" {
				return false
			}
		}
	}
	return true
}

func writerow(w http.ResponseWriter, colname [MAX_COLS]string, colvalue [MAX_COLS]string, numCols int) {
	for i, text := range colvalue {
		if i < numCols {
			io.WriteString(w, "  \"")
			io.WriteString(w, colname[i])
			io.WriteString(w, "\":\"")
			io.WriteString(w, text)
			io.WriteString(w, "\"")
			if i < (numCols - 1) {
				io.WriteString(w, ",")
			}
			io.WriteString(w, "\n")
		}
	}
}

func data(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "{\n\"data\": [")

	excelFileName := "/usr/local/bin/example.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		io.WriteString(w, "{\"item\":\"")
		io.WriteString(w, "error")
		io.WriteString(w, "\"}\n")
	} else {

		var firstrow = true

		for _, sheet := range xlFile.Sheets {
			var colname [MAX_COLS]string
			var colvalue [MAX_COLS]string

			for r, row := range sheet.Rows {
				for i, cell := range row.Cells {
					text := cell.String()
					if i < MAX_COLS {
						if r == 0 {
							colname[i] = text
						} else {
							colvalue[i] = text
						}
					}
				}
				if r > 0 {
					if !emptyrow(colvalue, len(row.Cells)) {
						if firstrow {
							firstrow = false
						} else {
							io.WriteString(w, ",")
						}
						io.WriteString(w, "\n{\n")
						writerow(w, colname, colvalue, len(row.Cells))
						io.WriteString(w, "}")
					}
				}
			}
		}
	}

	io.WriteString(w, "\n]\n}")

}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}
