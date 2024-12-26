package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pong/utils"
	"strconv"
)

func SaveNumber(number int) error {
	filePath, err := utils.GetFilePath()
	if err != nil {
		return err
	}

	data := strconv.Itoa(number)

	err = os.WriteFile(filePath, []byte(data), 0755)
	if err != nil {
		fmt.Printf("Unable to write to file: %v\n", err)
	}

	return nil
}

func ReadNumber() (int, error) {
	filePath, err := utils.GetFilePath()
	if err != nil {
		return 0, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return number, nil
}

func main() {
	count, err := ReadNumber()
	if err != nil {
		fmt.Println("Error reading number:", err)
	}

	fmt.Println("Starting server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		io.WriteString(w, fmt.Sprintf("pong %d", count))

		count = count + 1
		SaveNumber(count)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("Server started in port", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
