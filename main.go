package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/machinebox/graphql"
	"net/http"
	// "bytes"
	"github.com/joho/godotenv"
	// "io"
	"github.com/sireeshdevaraj/Go-anilistv1.0.0/utils"
	"log"
	"os"
)

func getAnilistDataOfUser(userId int) []byte {
	client := graphql.NewClient("https://graphql.anilist.co/")
	query := utils.Query // graphQL query
	formattedQuery := fmt.Sprintf(query, userId, userId)
	var response utils.Response

	req := graphql.NewRequest(formattedQuery)
	err := client.Run(context.Background(), req, &response)
	if err != nil {
		log.Fatal("ERROR: running the graphql client", err)
	}
	response.TruncateResponse() // I only need one entry for my website
	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		log.Fatal("ERROR: marshalling the struct", err)
	}
	// body := bytes.NewBuffer(jsonResponse)
	return jsonResponse
}

type ResponseFromServer = http.ResponseWriter
type Request = *http.Request

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Anilist handler
	getHandler := func(response ResponseFromServer, request Request) {
		// Maybe later, to allow others to fetch their own ID. Not really my case, and need to add rate limiting and stuff.
		// reqBytes,err := io.ReadAll(request.Body)
		// if err!=nil{
		// 	fmt.Println(err)
		// }
		// fmt.Println(bytes.NewBuffer(reqBytes).String())
		userId := 5631742
		data := getAnilistDataOfUser(userId)
		response.Header().Set("Content-Type", "Application/json")
		response.Write(data)
	}
	http.HandleFunc("/anilist", getHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
