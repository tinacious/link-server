package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/tinacious/link-server/templates"
	"github.com/tinacious/link-server/utils"
)

func main() {
	port := utils.GetRandomPort()

	// Use configured port
	portString := os.Getenv("PORT")
	if portString != "" {
		i, err := strconv.Atoi(portString)
		if err != nil {
			fmt.Printf("⛔️ invalid port %s - using random port %d\n", portString, port)
		} else {
			port = i
		}
	}

	var validInput bool = false
	if utils.IsInputPiped() || len(os.Args) > 1 {
		validInput = true
	}

	if !validInput {
		fmt.Println("must provide path to file or pipe data from standard in")
		os.Exit(1)
	}

	links, err := getInput()

	fmt.Printf("🔗 links server at: http://localhost:%d with links:\n%s\n", port, strings.Join(links, "\n"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := templates.CreateLinksPage(links)

		err = page.Render(w)
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getInput() ([]string, error) {
	if utils.IsInputPiped() {
		return utils.ParseTextFromStandardIn()
	}

	filePath := os.Args[1]
	if filePath == "" {
		fmt.Println("must provide path to file")
		os.Exit(1)
	}
	return utils.ParseLinksFromFile(filePath)
}
