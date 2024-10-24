package cli

import (
	"fmt"
	"io"
	"net/http"
)

// ListFiles lists all files on the server
func ListFiles() error {
	resp, err := http.Get("http://localhost:8080/files")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to list files, status code: %d", resp.StatusCode)
	}

	// Print the response body (list of files)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Files on server:")
	fmt.Println(string(body))

	return nil
}
