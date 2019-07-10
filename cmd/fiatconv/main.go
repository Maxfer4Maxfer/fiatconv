package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"fiatconv/pkg/exchanging"
	exchangrater "fiatconv/pkg/exchangerate/exchangeratesapiio"
)

func main() {

	// define a new exchang rater
	er := exchangrater.New(&http.Client{
		Timeout: time.Duration(20 * time.Second),
	})

	// define an exchange service
	cc := exchanging.NewCurrencyConverter(er)

	// read command line arrgumments
	if len(os.Args) != 4 {
		usage()
	}

	amount64, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		usage()
	}
	amount := float32(amount64)

	src := os.Args[2]
	dst := os.Args[3]

	convAmount, err := cc.Convert(amount, src, dst)

	// output a result
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(convAmount)
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("\tfiatconv <amount:float> <src_symbol:string> <dst_symbol:string>")
	os.Exit(1)
}
