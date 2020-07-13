package main

import (
    "fmt"
    "log"
	"net/http"
	"strings"
    "io/ioutil"
		"path/filepath"
		"encoding/json"

		"github.com/gorilla/mux"
		"github.com/go-yaml/yaml"
)

var productOwners ProductOwner

type ProductOwner struct {
	Product []struct {
		Name        string `yaml:"name"`
		Owner       string `yaml:"owner"`
		SlackHandle string `yaml:"slack_handle"`
		Email       string `yaml:"email"`
		Phone       string `yaml:"phone"`
	} `yaml:"product"`
}

type ProductOwnerDetails struct {
	Name string
	Owner string
	SlackHandle string
	Email string
	Phone string
}

// rootHandler is answering on / ("root")
func homePage(w http.ResponseWriter, r *http.Request) {

	// display to browser request
	fmt.Fprintf(w, "Welcome to the Home Page of the Service Owner API!")
	fmt.Println("Endpoint Hit: homePage")
}

//returnAllOwners returns contact info for all owners of all products
func returnAllProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllOwners")
	json.NewEncoder(w).Encode(productOwners.Product)
}

//returnSingleOwner returns contact information for owner of only the product specified in API call
func returnProductsByName(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnProductsByName")

	vars := mux.Vars(r)
	key := vars["prodName"]

	for _, name := range productOwners.Product{
		if name.Name == strings.ToLower(key) {
			json.NewEncoder(w).Encode(ProductOwnerDetails{Name: name.Name, Owner: name.Owner, Email: name.Email, SlackHandle: name.SlackHandle, Phone: name.Phone})
		}
	}
}

func returnProductsByOwner(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnProductsByOwner")
	vars := mux.Vars(r)
	key := vars["ownerName"]
	for _, owner := range productOwners.Product{
		if owner.Owner == strings.ToLower(key) {
			json.NewEncoder(w).Encode(ProductOwnerDetails{Name: owner.Name, Owner: owner.Owner, Email: owner.Email, SlackHandle: owner.SlackHandle, Phone: owner.Phone})
		}
	}
}

func returnProductsBySlackHandle(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnProductsBySlackHandle")
	vars := mux.Vars(r)
	key := vars["slackHandle"]
	for _, handle := range productOwners.Product{
		if handle.SlackHandle == strings.ToLower(key) {
			json.NewEncoder(w).Encode(ProductOwnerDetails{Name: handle.Name, Owner: handle.Owner, Email: handle.Email, SlackHandle: handle.SlackHandle, Phone: handle.Phone})
		}
	}
}

func returnProductsByEmail(w http.ResponseWriter, r *http.Request){
//	var users []User
	fmt.Println("Endpoint Hit: returnProductsByEmail")
	vars := mux.Vars(r)
	key := vars["emailAddress"]
	for _, email := range productOwners.Product{
		if email.Email == strings.ToLower(key) {
//			users = append(users, User{Name: owner.Name, Email: owner.Email})
//			json.NewEncoder(w).Encode(users)
			json.NewEncoder(w).Encode(ProductOwnerDetails{Name: email.Name, Owner: email.Owner, Email: email.Email, SlackHandle: email.SlackHandle, Phone: email.Phone})
		}
	}
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))
}

func handleRequests() {
  // creates a new instance of a mux router
  productOwnerRouter := mux.NewRouter().StrictSlash(true)
	productOwnerRouter.HandleFunc("/", homePage)
	productOwnerRouter.HandleFunc("/serviceowner/api/v1/product", returnAllProducts)
	productOwnerRouter.HandleFunc("/serviceowner/api/v1/product/name/{prodName}", returnProductsByName)
	productOwnerRouter.HandleFunc("/serviceowner/api/v1/product/owner/{ownerName}", returnProductsByOwner)
	productOwnerRouter.HandleFunc("/serviceowner/api/v1/product/handle/{slackHandle}", returnProductsBySlackHandle)
	productOwnerRouter.HandleFunc("/serviceowner/api/v1/product/email/{emailAddress}", returnProductsByEmail)

	//myRouter.HandleFunc("/serviceowner/api/v1/product", createNewArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":9858", productOwnerRouter))

}



func main() {
	  fmt.Println("Rest API v1.0 - Service Owner configured to listen")

    filename, _ := filepath.Abs("./service_owner.yml")
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
		}

    err = yaml.Unmarshal(yamlFile, &productOwners)
    if err != nil {
        log.Println(err)
		}

	handleRequests()

}

