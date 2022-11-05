package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type OperationType string 
const (
	ADDITION OperationType = "addition"
	SUBTRACTION OperationType = "subtraction"
	MULTIPLICATION OperationType = "multiplication"
	DIVISION OperationType = "division"

	SlackUsername = "Omiete"
)

var (
	keywords = []string{"add", "multiply", "mul", "divide", "div", "subtract", "sub"}
)

type Response struct {
	SlackUsername string 	`json:"slackUsername"`
	Result 		  int32  		`json:"result"`
	OperationType string  	`json:"operation_type"`
}

type request struct {
	OperationType string  	`json:"operation_type"`
	X 			  int 		`json:"x"`
	Y 			  int 		`json:"y"`
}

func CheckOperationType(body request, resp *Response) {
	switch body.OperationType {
	case string(ADDITION):
		resp.OperationType = string(ADDITION)
		resp.Result = int32(body.X + body.Y) 
		
	case string(SUBTRACTION):
		resp.OperationType = string(SUBTRACTION)
		resp.Result = int32(body.X - body.Y) 
		
	case string(MULTIPLICATION):
		resp.OperationType = string(MULTIPLICATION)
		resp.Result = int32(body.X * body.Y) 
		
	case string(DIVISION):
		resp.OperationType = string(DIVISION)
		resp.Result = int32(body.X / body.Y)

	default:
		for _, word := range keywords {

			if strings.Contains(body.OperationType, word) {

				switch word {
				case "add":
					resp.OperationType = string(ADDITION)
					resp.Result = int32(body.X + body.Y)
				case "mul":
					resp.OperationType = string(MULTIPLICATION)
					resp.Result = int32(body.X * body.Y)
				case "multiply":
					resp.OperationType = string(MULTIPLICATION)
					resp.Result = int32(body.X * body.Y)
				case "div":
					resp.OperationType = string(DIVISION)
					resp.Result = int32(body.X / body.Y)
				case "divide":
					resp.OperationType = string(DIVISION)
					resp.Result = int32(body.X / body.Y)
				case "sub":
					resp.OperationType = string(SUBTRACTION)
					resp.Result = int32(body.X - body.Y)
				case "subtract":
					resp.OperationType = string(SUBTRACTION)
					resp.Result = int32(body.X - body.Y)
				}
			}
		}
	}
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var body request
		var resp Response

		resp.SlackUsername = SlackUsername

		json.NewDecoder(req.Body).Decode(&body)
		defer req.Body.Close()

		CheckOperationType(body, &resp)

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "Post")
		json.NewEncoder(res).Encode(resp)
	})

	fmt.Println("server listening on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	
}