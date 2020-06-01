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

func root(x float64) float64{
	return math.Sqrt(x)
}

func getnum (w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		log.Printf("Received request for sqrt %s\n", query)
		number := query.Get("number")
		convToFlt, err := strconv.ParseFloat(number, 64)
		if err != nil{
			log.Printf("The error occurred %v\n", err)
			w.Write([]byte("An error occurred, please check the entered number"))
		} else{
			log.Printf("The number received is %v\n", convToFlt)
			w.Write([]byte(fmt.Sprintf("The number is  %s\n", number)))
			d :=root(convToFlt)
			log.Printf("Counting the square root, the result is %v\n", d)
			convToStr := strconv.FormatFloat(d, 'f', 10,32)
			w.Write([]byte(fmt.Sprintf("The square root for %s is %s...", number, convToStr)))

		}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/sqrt", getnum)

	log.Fatal(http.ListenAndServe(":8080", nil))
}