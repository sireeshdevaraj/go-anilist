package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
	"github.com/sireeshdevaraj/Go-anilistv1.0.0/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getAnilistDataOfUser(userId int) ([]byte, error) {
	client := graphql.NewClient("https://graphql.anilist.co/")
	query := utils.Query // graphQL query
	formattedQuery := fmt.Sprintf(query, userId, userId)
	var response utils.Response

	req := graphql.NewRequest(formattedQuery)
	err := client.Run(context.Background(), req, &response)
	if err != nil {
		fmt.Println("ERROR: running the graphql client", err)
		return nil, err
	}
	response.TruncateResponse() // I only need one entry for my website
	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		fmt.Println("ERROR: marshalling the struct", err)
		return nil, err
	}
	// body := bytes.NewBuffer(jsonResponse)
	return jsonResponse, nil
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
		split_array := strings.Split(request.URL.Path, "/")
		id, err := strconv.Atoi(split_array[len(split_array)-1])
		if err != nil {
			http.Error(response, "Error processing the request", 400)
			return
		}
		userId := id
		data, err := getAnilistDataOfUser(userId)
		if err != nil {
			http.Error(response, "Error in querying the ID from anilist", 400)
			return
		}
		response.Header().Set("Content-Type", "Application/json")
		response.Write(data)
	}
	http.HandleFunc("/anilist/", getHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
