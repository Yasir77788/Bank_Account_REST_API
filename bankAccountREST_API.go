// Yasir Hassan
// Bank Account REST API - CRUD Operations 
// Written in GoLang programming language
package main

//  add required packages
import (
	"encoding/json" // parse JSON data that received from requests into go data and vice versa
	"fmt"
	"log"      // implement logging capabilities for the API such as logging errors in requests
	"net/http" // allow us to receive, parse and send http requests
)

type Account struct {
	Number  string `json: "AccountNumber"`
	Balance string `json: "Balance"`
	Desc    string `json: "AccountDescription"`
}

// create a dataset of accounts, store in a slice
// use global variable to be accessed by different functions
var Accounts []Account

// Home page writer, use http.ResponseWriter to create an initial message for the application
// so that requests return something right off the bat.... allows us to check the connection with the API.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our bank!")
	fmt.Println("Endpoint: /")
}

// the app to return the accounts in the dataset in a JSON format.
// Create returnAllAccounts to handle this process.
// use the Encode function to convert the Account into a json object
func returnAllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Accounts)
}

// To handle the HTTP requests. use a handleRequests function
// This function returns homepage or returnAllAccounts, based on the URL provided with the request.
// reach the API at the address http://localhost:10000 while the program is running
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/accounts", returnAllAccounts)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// Add data
// In the main function, create fictional data for the Accounts dataset and execute the handleRequest function.
func main() {
	// initialize the dataset
	Accounts = []Account{
		Account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
		Account{Number: "S3r53455345", Balance: "444.4", Desc: "Saving Account"},
	}

	handleRequests()
}
