package main

import (
	"fmt"
	"github.com/yosssi/gohtml"
	"io"

	//"github.com/yosssi/gohtml"
	//"io"
	"log"
	"net/http"
)

func main() {
	urlBase := "https://www.workana.com/jobs?"

	fmt.Println("\t\t\t\t\t\t\t\t\t\tHello gopper!!!")

	res, err := http.Get(urlBase)
	if err != nil {
		log.Fatal(err)
	}

	/*
		 //-------- Print BODY to go map --------
		fmt.Printf("\tHeader map:\n\n")
		for i, v := range res.Header {
			fmt.Printf("\tKey: %v \tValue: %v\n\n\n", i, v)
		}*/

	fmt.Printf("\tBody map:\n\n")
	fmt.Printf("Olhando res.Body melhor...\n")

	body, err := io.ReadAll(res.Body)

	fmt.Printf("%v", gohtml.Format(string(body)))
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	/*for i, v := range res.Header {
		fmt.Printf("\tKey: %v \tValue: %v\n\n\n", i, v)
	}*/

}
