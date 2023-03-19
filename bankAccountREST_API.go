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

	"github.com/gorilla/mux" // Gorilla Mux package implements a request router and dispatcher that
	// matches incoming requests to their respective handler.
	// Also, it parses data sent through HTTP requests.
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

// global function to return an account
// Access the variables sent in the request from the mux router.
// Access the account number value that was sent by the HTTP request.
// The convention here is that the parameter’s name is number.
// Iterate through the dataset and when we find the account with the corresponding account number,
// encode the account in JSON format and write the data to the HTTPWriter w.
func returnAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["number"]
	for _, account := range Accounts {
		if account.Number == key {
			json.NewEncoder(w).Encode(account)
		}
	}
}

// To handle the HTTP requests. use a handleRequests function
// This function returns homepage or returnAllAccounts, based on the URL provided with the request.
// reach the API at the address http://localhost:10000 while the program is running
func handleRequests() {
	// create a router to handle our requests from the mux package.
	// StrictSlash defines the trailing slash behavior for new routes. The initial value is false.
	// When true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa. Ess
	// this guarantees that the application will always see the path as specified in the route.
	// use the new router variable to handle calls to the API, rather than using the built-in http package.
	// use the mux router as a custom handler for the ListenAndServe function
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/accounts", returnAllAccounts)
	router.HandleFunc("/account/{number}", returnAccount) // Use mux to reference specific parts of a record.
	log.Fatal(http.ListenAndServe(":10000", router))
}

// Add data
// In the main function, create fictional data for the Accounts dataset
// and execute the handleRequest function.
func main() {
	// initialize the dataset
	Accounts = []Account{
		Account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
		Account{Number: "S3r53455345", Balance: "444.4", Desc: "Saving Account"},
	}

	handleRequests()
}
