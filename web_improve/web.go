package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request for greeting")
	fmt.Fprintf(w, "wassup, man!!!\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request for headers")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}


func sqrtStr (u string)(float64, error){

	convToFlt, err := strconv.ParseFloat(u, 64)
	if err != nil{
		log.Printf("The error occurred %v\n", err)
	} else {
		log.Printf("The converted number is %v\n", convToFlt)
	}
	return math.Sqrt(convToFlt), err
}


func count (w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	log.Printf("Received request for sqrt %s\n", query)
	number := query.Get("number")
	x, err := sqrtStr(number)

	if err != nil {
		w.Write([]byte("An error occurred, please check the entered number"))
	} else {
		w.Write([]byte(fmt.Sprintf("The number is  %s\n", number)))
		log.Printf("Counting the square root, the result is %v\n", x)
		convToStr := strconv.FormatFloat(x, 'f', 10,32)
		w.Write([]byte(fmt.Sprintf("The square root for %s is %s...", number, convToStr)))
	}
}


func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/sqrt", count)

	log.Fatal(http.ListenAndServe(":8080", nil))
}