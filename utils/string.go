package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func Print(x any) {
	// Pretty-print the resulting JSON
	prettyJSON, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate pretty JSON:", err)
	}

	// Output the final JSON
	fmt.Println(string(prettyJSON))
}
