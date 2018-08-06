package main

import (
	"io"
	"net/http"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
)

// MaxCols is the max columns in a row
const MaxCols = 256

// BinPath is the path where binary Files are Located
const BinPath = "/usr/local/bin/"

// DataPath is the path where xlsx Files are Located
const DataPath = "/tmp/data/"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data/{id}", data)
	r.HandleFunc("/health", health)

	http.ListenAndServe(":8080", r)
}

func getPath() string {

	path := DataPath

	if _, err := os.Stat(DataPath); os.IsNotExist(err) {
		path = BinPath
	}

	return path
}

func emptyrow(colvalue [MaxCols]string, numCols int) bool {
	for i, text := range colvalue {
		if i < numCols {
			if text != "" {
				return false
			}
		}
	}
	return true
}

func writerow(w http.ResponseWriter, colname [MaxCols]string, colvalue [MaxCols]string, numCols int) {
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

	vars := mux.Vars(r)
	id := vars["id"]

	io.WriteString(w, "{\n\"data\": [")

	excelFileName := getPath() + id + ".xlsx"
	xlFile, err := excelize.OpenFile(excelFileName)
	if err != nil {
		io.WriteString(w, "{\"error\":\"")
		io.WriteString(w, "error opening file "+excelFileName)
		io.WriteString(w, "\"}\n")
	} else {

		var firstrow = true

		for _, name := range xlFile.GetSheetMap() {
			var colname [MaxCols]string
			var colvalue [MaxCols]string

			for r, row := range xlFile.GetRows(name) {
				for i, cell := range row {
					text := cell
					if i < MaxCols {
						if r == 0 {
							colname[i] = text
						} else {
							colvalue[i] = text
						}
					}
				}
				if r > 0 {
					if !emptyrow(colvalue, len(row)) {
						if firstrow {
							firstrow = false
						} else {
							io.WriteString(w, ",")
						}
						io.WriteString(w, "\n{\n")
						writerow(w, colname, colvalue, len(row))
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
