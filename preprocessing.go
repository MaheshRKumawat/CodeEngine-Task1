package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func preprocess(fileName string) {
	fi, erri := os.Open(fileName)
	fo, erro := os.Create("salesP.csv")

	if erri != nil {
		log.Fatal("Failed opening file, error: %s", erri)
	}
	if erro != nil {
		log.Fatal("Failed creating file, error: %s", erro)
	}

	csvReader := csv.NewReader(fi)
	csvWriter := csv.NewWriter(fo)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error in reading records, error: ", err)
		} else {
			// Check if there is a valid price and quantity values in the dataset.
			_, errQ := strconv.ParseInt(rec[2], 10, 64)
			_, errP := strconv.ParseFloat(rec[3], 10)

			if errQ != nil || errP != nil {
				// For non-int values
				continue
			} else {
				_ = csvWriter.Write(rec)
			}
		}
	}
	// You need to call the Flush method of your CSV writer to ensure all buffered data is written to your file before closing the file.
	csvWriter.Flush()

	fi.Close()
	fo.Close()
}

// func main() {
// 	preprocess("Sales_August_2019.csv")
// }
