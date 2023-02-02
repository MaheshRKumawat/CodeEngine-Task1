package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func reduce(fileName string) {
	fi, erri := os.Open(fileName)

	if erri != nil {
		log.Fatal("Failed opening file, error: %s", erri)
	}

	totalSale := map[string]int64{}
	price := map[string]float64{}
	csvReader := csv.NewReader(fi)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error in reading records, error: ", err)
		} else {
			quant, errQ := strconv.ParseInt(rec[1], 10, 64)
			pr, errP := strconv.ParseFloat(rec[2], 10)

			if errQ != nil || errP != nil {
				// For non-int values
				continue
			}
			totalSale[rec[0]] += quant
			price[rec[0]] = pr
		}
	}
	fi.Close()

	frq, errq := os.Create("salesQuantity.csv")
	frp, errp := os.Create("salesPrice.csv")
	if errq != nil {
		log.Fatal("Failed creating file, error: %s", errq)
	}
	if errp != nil {
		log.Fatal("Failed creating file, error: %s", errp)
	}

	// Sorting Algo
	keys := []string{}
	for key := range totalSale {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return totalSale[keys[i]] > totalSale[keys[j]]
	})

	csvWriterQ := csv.NewWriter(frq)
	csvWriterP := csv.NewWriter(frp)

	for _, name := range keys {
		quantStr := strconv.Itoa(int(totalSale[name]))
		crQ := []string{name, quantStr}
		_ = csvWriterQ.Write(crQ)
	}

	for _, name := range keys {
		crP := []string{name, strconv.FormatFloat(price[name], 'g', 8, 64)}
		_ = csvWriterP.Write(crP)
	}

	// You need to call the Flush method of your CSV writer to ensure all buffered data is written to your file before closing the file.
	csvWriterQ.Flush()
	csvWriterP.Flush()
	frq.Close()
	frp.Close()
}

func main() {
	reduce("salesM.csv")
}
