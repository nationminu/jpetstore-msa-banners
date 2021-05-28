package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"math/rand"
)

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Host on which to run")
	flag.StringVar(&port, "port", "8080", "Port on which to run")
}

type Banners struct {
	Banners []Banner `json:"banners"`
}

type Banner struct {
	Favcategory string `json:"favcategory"`
	Bannername  string `json:"bannername"`
	Descn       string `json:"descn"`
	Image       string `json:"image"`
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func forbidden(w http.ResponseWriter, r *http.Request) {
	// see http://golang.org/pkg/net/http/#pkg-constants
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("403 HTTP status code returned!"))
}

func find() Banners {

	banners := Banners{}
	data, err := ioutil.ReadFile("data/banners.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened banners.json")

	err = json.Unmarshal(data, &banners)
	if err != nil {
		log.Fatal(err)
	}

	return banners
}

func one(id string) Banner {

	banners := Banners{}
	banner := Banner{}
	data, err := ioutil.ReadFile("data/banners.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened banners.json")

	err = json.Unmarshal(data, &banners)
	if err != nil {
		log.Fatal(err)
	}

	for i := range banners.Banners {
		//fmt.Println("|",banners.banners[i].Catid , "|==|" , id,"|")
		if string(banners.Banners[i].Favcategory) == id {
			//fmt.Println(banners.banners[i])
			banner = banners.Banners[i]
		}
	}

	return banner
}

func findAll(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

	} else if r.Method == "GET" {
		fmt.Println("Endpoint Hit: getBanners")

		banners := find()

		output, err := json.Marshal(banners.Banners)
		//output, err := json.MarshalIndent(banners.Banners, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		fmt.Println(string(output))

		fmt.Println("Endpoint Hit: getBanners")
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 HTTP status code returned!"))
	}
}

func findOne(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/banners/")

	if id == "random" {

		banners := find()
		banner := Banner{}

		random := rand.Intn(len(banners.Banners))

		//fmt.Println(len(banners.Banners))
		//fmt.Println(random)
		for i := range banners.Banners {
			if i == random {
				banner = banners.Banners[i]
			}
		}

		output, err := json.Marshal(banner)
		//output, err := json.MarshalIndent(categories.Categories, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		fmt.Println(string(output))

	} else {
		banners := one(id)
		output, err := json.Marshal(banners)
		//output, err := json.MarshalIndent(categories.Categories, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		fmt.Println(string(output))
	}
}

func handleRequests() {
	http.HandleFunc("/favicon.ico", doNothing)
	http.HandleFunc("/banners", findAll)
	http.HandleFunc("/banners/", findOne)
	address := ":" + port

	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	handleRequests()
}
