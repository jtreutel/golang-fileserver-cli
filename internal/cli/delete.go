package cli

import (
	"fmt"
	"net/http"
)

// DeleteFile deletes a file on the server
func DeleteFile(fileName string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/files/%s", fileName), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete file, status code: %d", resp.StatusCode)
	}

	return nil
}
