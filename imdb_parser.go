package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
	)

var filename string = "WATCHLIST.csv"

type Movie struct{
	position 	string	
	code 		string
	created 	string
	modified 	string
	description 	string
	title 		string
	title_type 	string
	director	string
	you_rated	string
	imdb_rating	string
	runtime_mins	string
	year		string
	genres		string // It may be more usefull to use genres as []string
	num_votes	string
	release_date	string
	url		string
}

type MovieSlice []Movie

func (m *MovieSlice) Append(entry Movie) {
	*m = append(*m, entry)
}
 
var movs MovieSlice
 
func main() {
	file, err := os.Open("./WATCHLIST.csv")
	file.Seek(0, os.SEEK_SET)
	defer file.Close()
	if(err != nil) {
		fmt.Println("Something bad happened!")	
	} else {
		reader := csv.NewReader(file)
		for {
			row, err := reader.Read()
			if(err == io.EOF) { break }
			if(err != nil) { fmt.Println("ERROR!", err) }
			if(len(row) == 0) { continue }
			position := row[0]
			code := row[1]
			created := row[2]
			modified := row[3]
			description := row[4]
			title := row[5]
			title_type := row[6]
			director := row[7]
			you_rated := row[8]
			imdb_rating := row[9]
			runtime_mins := row[10]
			year := row[11]
			genres := row [12]
			num_votes := row[13]
			release_date := row[14]
			url := row[15]
			movs.Append(Movie{ position, code, created, modified, description, title, title_type, director, you_rated, imdb_rating, runtime_mins, year, genres, num_votes, release_date, url })
 
		}
			println("Loaded", len(movs), "Movies")
	}
	for i, _ := range movs {
 		fmt.Println(movs[i].title, movs[i].director)
	}
}

