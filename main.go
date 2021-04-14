package main

import (
	Lib "employee/common/lib"
	Types "employee/common/types"
	Employee "employee/model/employee"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	Lib.JsonResponse(w, Employee.GetEmployees())
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	Lib.JsonResponse(w, Employee.GetEmployeeById(id))
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Types.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	var response Types.Response4Update
	response.Result, response.Message = Employee.AddEmployee(employee)

	Lib.JsonResponse(w, response)
}

func DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var response Types.Response4Update
	response.Result, response.Message = Employee.DeleteEmployeeById(id)

	Lib.JsonResponse(w, response)
}

func UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var employee Types.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = id

	var response Types.Response4Update
	response.Result, response.Message = Employee.UpdateEmployeeById(employee)

	Lib.JsonResponse(w, response)
}

func main() {

	r := mux.NewRouter()
	/*
		Mock up data for creating employee:

		{
			"id": 1,
			"firstname": "Jonh",
			"lastname": "Smith",
			"BirthDay": {
				"year": 1988,
				"month": 12,
				"day": 23
			}
		}
	*/
	r.HandleFunc("/create/employee", CreateEmployee).Methods("PUT")
	r.HandleFunc("/get/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/get/employee/{id}", GetEmployeeById).Methods("GET")

	/*
		Mock up data for updating employee:

		{
			"firstname": "New Jonh",
			"lastname": "New Smith",
			"BirthDay": {
				"year": 1980,
				"month": 5,
				"day": 29
			}
		}
	*/
	r.HandleFunc("/update/employee/{id}", UpdateEmployeeById).Methods("POST")
	r.HandleFunc("/delete/employee/{id}", DeleteEmployeeById).Methods("DELETE")

	http.ListenAndServe(":80", r)

}
