// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Derek701/gopl-solutions/ch4/ex4-10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	fmt.Println("\nLess than 30 days old...")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 < 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("\nBetween 30 and 365 days old...")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 >= 30 && now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("\nMore than 365 days old...")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
