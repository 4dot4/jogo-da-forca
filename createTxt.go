package main

import (
	"fmt"
	"log"
	"os"
)

type Words struct {
	palavras []string
}

func main() {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{

		"abajur",
		"agulha",
		"alfinete",
		"algema",
		"bacia",
		"balaio",
		"balde",
		"banco",
		"bandeira",
		"bandolim",
		"batuta",
		"faca",
		"fagote",
		"fantoche",
		"farol",
		"fax",
		"garrafa",
		"mesa",
		"ventilador",
		"cadeira",
		"ferradura",
		"bolsa",
		"tv",
		"ferro",
		"filmadora",
		"filtro",
		"fio",
		"fita"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}
