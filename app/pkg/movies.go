package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movies struct {
	Id          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
}

const filePath = "pkg/movies.json"

type MoviesList struct {
	Movies []Movies `json:"movies"`
}

func AddingMovies(movie Movies) error {
	var list MoviesList

	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
	}

	list.Movies = append(list.Movies, movie)

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	fmt.Printf("Movie added: %+v\n", movie)
	return nil
}

func GetMovies() MoviesList {
	var list MoviesList

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("No movies found or error reading file:", err)
		return list
	}

	if err := json.Unmarshal(data, &list); err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	return list
}
