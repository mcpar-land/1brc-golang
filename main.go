package main

import (
	"compress/gzip"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
)

func main() {
	genFlag := flag.Bool("gen", false, "generate the billion row file as a gzipped text file")
	flag.Parse()

	if *genFlag {
		err := generateFile()
		if err != nil {
			fmt.Println("error generating billion row file:", err)
			os.Exit(1)
		}
		fmt.Println("done")
		os.Exit(0)
	}
}

func generateFile() error {
	cityNamesSrc := "https://raw.githubusercontent.com/gunnarmorling/1brc/refs/heads/main/data/weather_stations.csv"

	fmt.Println("fetching from", cityNamesSrc)

	resp, err := http.Get(cityNamesSrc)
	if err != nil {
		return err
	}

	r := csv.NewReader(resp.Body)
	r.Comma = ';'
	r.Comment = '#'

	csvRaw, err := r.ReadAll()
	if err != nil {
		return err
	}

	cityNames := []string{}

	for _, row := range csvRaw {
		cityNames = append(cityNames, row[0])
	}

	fmt.Println("successfully parsed", len(cityNames), "city names")

	file, err := os.Create("1brc.csv.gz")
	if err != nil {
		return err
	}

	gzipWriter := gzip.NewWriter(file)

	rng := rand.New(rand.NewPCG(42, 1024))

	for i := 0; i < 1_000_000_000; i++ {
		randomCity := cityNames[rng.IntN(len(cityNames))]
		row := fmt.Sprintf(
			"%s;%.2f\n",
			randomCity,
			rng.Float32()*100.0,
		)
		_, err = gzipWriter.Write([]byte(row))
		if err != nil {
			return err
		}
		if i%1_000_000 == 0 {
			fmt.Println(i, "rows written!")
		}
	}

	err = gzipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}
