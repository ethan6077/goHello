package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadFromListingSearchApi(listingId string) {
	baseUrl := "https://lsapi.listings-search-pipeline.resi-property.realestate.com.au/services/listings/"
	url := baseUrl + listingId

	res, errGet := http.Get(url)

	if errGet != nil {
		fmt.Println("Error fetching data from API.")
		panic(errGet)
	}

	defer res.Body.Close()

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading data from API.")
	// 	panic(err)
	// }
	// fmt.Println("Data from API: ", string(data))

	var listing Listing
	errDecoding := json.NewDecoder(res.Body).Decode(&listing)
	if errDecoding != nil {
		fmt.Println("Error decoding data from API.")
		panic(errDecoding)
	}

	fmt.Printf("Listing: %+v\n", listing)
}
