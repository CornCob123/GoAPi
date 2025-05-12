package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"
type Response struct {
	Message string `json:message`
	Price float64 `json:"price"`
	Status  int    `json:"status"`
}
type CoinMarketCapResponse struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}
func getBTC(w http.ResponseWriter, r *http.Request) {
		// Replace with your actual CoinMarketCap API key
	apiKey := "" // Replace this with your CoinMarketCap API Key

	// Make the API request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	// Add query parameters
	query := req.URL.Query()
	query.Add("symbol", "BTC")
	query.Add("convert", "USD")
	req.URL.RawQuery = query.Encode()

	// Initialize HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var result CoinMarketCapResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Extract the Bitcoin price in USD
	btcPrice := result.Data["BTC"].Quote["USD"].Price
	fmt.Printf("Current Bitcoin Price: $%.2f\n", btcPrice)

	// Creating the response object
	response := Response{
		Message: "Bitcoin Price Retrieved Successfully",
		Price:   btcPrice,
		Status:  200,
	}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding the response as JSON and writing to the response writer
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/bitcoin", getBTC)
    fmt.Println("Server starting at port 8080...")
    http.ListenAndServe(":8080", nil)
}