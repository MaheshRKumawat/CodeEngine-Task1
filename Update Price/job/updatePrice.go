package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func updatePrice(priceFile string) {
	fp, errp := os.Open(priceFile)

	if errp != nil {
		log.Fatal("Failed opening file, error: %s", errp)
	}

	csvReaderP := csv.NewReader(fp)
	price := map[string]float64{}
	keys := []string{}
	length := 0
	count := 0

	for {
		rec, err := csvReaderP.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error in reading records, error: ", err)
		} else {
			pr, errP := strconv.ParseFloat(rec[1], 10)

			if errP != nil {
				continue
			}
			price[rec[0]] = pr
			keys = append(keys, rec[0])
			length++
		}
	}
	fp.Close()

	for _, key := range keys {
		if count <= length/10 { // Decimal value will be rounded off to floor value
			price[key] += price[key] * 0.1
		} else if count >= (length - (length / 10)) {
			price[key] -= (price[key] * 0.1)
		}
		count++
	}

	fo, erro := os.Create("UpdatedPrice.csv")
	csvWriter := csv.NewWriter(fo)
	if erro != nil {
		log.Fatal("Failed creating file, error: %s", erro)
	}

	for key, value := range price {
		prc := []string{key, strconv.FormatFloat(value, 'g', 8, 64)}
		_ = csvWriter.Write(prc)
	}
	csvWriter.Flush()
	fo.Close()
}

func main() {
	updatePrice("salesPrice.csv")
}
