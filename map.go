package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func mapper(fileName string) {
	fi, erri := os.Open(fileName)
	fo, erro := os.Create("salesM.csv")

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
			name := rec[1]
			// Check if there is a valid price and quantity values in the dataset.
			quant := rec[2]
			price := rec[3]

			cr := []string{name, quant, price}
			_ = csvWriter.Write(cr)
		}
	}
	// You need to call the Flush method of your CSV writer to ensure all buffered data is written to your file before closing the file.
	csvWriter.Flush()

	fi.Close()
	fo.Close()
}

func main() {
	mapper("salesP.csv")
}
