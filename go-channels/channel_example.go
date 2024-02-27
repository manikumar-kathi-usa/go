package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type City struct {
	ID        int
	STATECODE string
	STATENAME string
	CITY      string
}

func main() {
	f, err := os.Open("us_cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := readRow(f)
	for c := range rows {
		log.Println(c)
	}
}

func readRow(r io.Reader) chan City {
	out := make(chan City)

	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				log.Fatal(err)
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			id, err := strconv.Atoi(row[0])
			if err != nil {
				log.Fatal(err)
			}

			out <- City{
				ID:        id,
				STATECODE: row[1],
				STATENAME : row[2],
				CITY: row[3],
			}
		}
		close(out)
	}()
	return out
}
