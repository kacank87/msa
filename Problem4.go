package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type AutoGenerated struct {
	Help    string `json:"help"`
	Success bool   `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
		Fields     []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"fields"`
		Records []struct {
			ID            int    `json:"_id"`
			Sex           string `json:"sex"`
			NoOfGraduates string `json:"no_of_graduates"`
			TypeOfCourse  string `json:"type_of_course"`
			Year          string `json:"year"`
		} `json:"records"`
		Links struct {
			Start string `json:"start"`
			Next  string `json:"next"`
		} `json:"_links"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"result"`
}

func addcol(fname string, column []string) {
	// read the file
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(column)
	w.Flush()
}

func write_csv(csvname string, year string, sex string, typeofcourse string, nograduates string) {
	var csvFile *os.File

	if _, err := os.Stat("./" + csvname + ".csv"); errors.Is(err, os.ErrNotExist) {
		//fmt.Println("ganok")
		csvFile, err = os.Create("./" + csvname + ".csv")
		if err != nil {
			fmt.Println(err)
		}
	}
	defer csvFile.Close() // This code close the file.
	tulis := []string{
		year, sex, typeofcourse, nograduates,
	}
	go addcol("./"+csvname+".csv", tulis)

}
func read_json_decode(url string) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(response.Body)
	customer := &AutoGenerated{}
	decoder.Decode(customer)
	data := customer.Result.Records
	fmt.Println(len(data))

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
		go write_csv(data[i].Year, data[i].Year, data[i].Sex, data[i].TypeOfCourse, data[i].NoOfGraduates)
	}

}
func main4() {
	read_json_decode("https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=100")
	time.Sleep(10 * time.Second)

}