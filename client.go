package main


import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
	"csutils"
)


func client(){
	request:= new(se.Request)
	response:= new(Response)
	request.init([]string {"https://google.com", "https://yahoo.com"},"lol")
	requestBody, err := json.Marshal(request)
	checkError("json was not encoded",&err)

	res,err:=http.Post("http://localhost:8181/", "application/json", bytes.NewBuffer(requestBody))
	checkError("bad request",&err)
	respData, err := ioutil.ReadAll(res.Body)
	checkError("No response",&err)
	json.NewDecoder(strings.NewReader(string(respData))).Decode(&response)
	fmt.Println(response.FoundAtSite)
}
func main() {
	client()
}
