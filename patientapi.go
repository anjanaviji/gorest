package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

//Patient details
type Patient struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
}

var patientList []Patient

func main() {

	// const n = 300000
	// fmt.Print(n / 3.56)
	//fmt.Println("Heyyy thats my first line in GO!!!!!!!!!!!")
	patientList = []Patient{{"1", "Joseph Mathew", "45", "9847563214"}, {"2", "Dipika Reddy", "32", "6345231289"}}
	fmt.Println(patientList)
	handleRequests()
	//fmt.Scanln()
	//fmt.Println("done")

}

func handleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/patient/{id}", deletePatient).Methods("DELETE")
	router.HandleFunc("/patient/{id}", viewPatient)
	router.HandleFunc("/patient", createPatient).Methods("POST")
	router.HandleFunc("/patient", updatePatient).Methods("PUT")
	http.ListenAndServe(":8080", router)
	// http.HandleFunc("/", multiplexer)

	// http.ListenAndServe(":1000", nil)
}

func multiplexer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		viewPatient(w, r)
	case "POST":
		createPatient(w, r)
	case "PUT":
		updatePatient(w, r)
	case "DELETE":
		deletePatient(w, r)

	}
}

func viewPatient(w http.ResponseWriter, r *http.Request) {

	//fmt.Print(r)
	// vars := mux.Vars(r)
	// key := vars["id"]
	var id = path.Base(r.RequestURI)
	var pat Patient
	var flag bool
	for _, patient := range patientList {
		if patient.ID == id {
			pat = patient
			flag = true
		}
	}
	if flag == true {
		json.NewEncoder(w).Encode(pat)

	} else {
		//json.NewEncoder(w).Encode("Error : Patient Doesn't Exist")
		http.Error(w, "Patient Doesn't Exist", 404)
	}

}

func createPatient(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Println("body-", reqBody)
	var patient Patient
	json.Unmarshal(reqBody, &patient)
	// var msg
	// output,err:=json.Unmarshal(reqBody, &patient)

	patientList = append(patientList, patient)
	fmt.Println(patient.Name, "---", patientList[2].Name)
	json.NewEncoder(w).Encode(patientList)

}

func updatePatient(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var patient Patient
	var flag bool
	err := json.Unmarshal(reqBody, &patient)
	if err != nil {
		//
		json.NewEncoder(w).Encode("Error : invalid resuest body")
	} else {
		fmt.Println("body-", patient, "-", err)
		for i, pat := range patientList {
			if pat.ID == patient.ID {
				flag = true
				patientList[i] = patient
				fmt.Println(pat)
			}
		}
		if flag == true {
			json.NewEncoder(w).Encode(patientList)
		} else {
			//json.NewEncoder(w).Encode("Error:Patient doesn't exists")
			http.Error(w, "Patient Doesn't Exist", 404)
		}

	}

}

func deletePatient(w http.ResponseWriter, r *http.Request) {

	// base, err := url.Parse(r.RequestURI)
	// if err != nil {
	// 	return
	// }
	//fmt.Println("here--" + path.Base(r.RequestURI))

	var id = path.Base(r.RequestURI)
	var flag bool
	for i, pat := range patientList {
		if pat.ID == id {
			flag = true
			var lastIndex = len(patientList) - 1
			patientList[i] = patientList[lastIndex]
			//var patient *Patient =nil
			//patientList[len(patientList)] =*patient
			patientList = patientList[:lastIndex]

		}
	}
	if flag == true {
		json.NewEncoder(w).Encode(patientList)
		return
	}
	http.Error(w, "Patient Doesn't Exist", 404)
	// Path params
	//base.Path += "this will get automatically encoded"
	//fmt.Println(base.Path)
	// Query params
	//params := url.Values{}
	//fmt.Println(params)
	//params.Add("q", "this will get encoded as well")
}
