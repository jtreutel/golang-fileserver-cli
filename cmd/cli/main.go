package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jtreutel/golang-fileserver-cli/internal/cli"
)

func main() {
	// Define CLI flags for each command
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	uploadFilePath := uploadCmd.String("file", "", "Path to the file to upload")
	deleteFileName := deleteCmd.String("name", "", "Name of the file to delete")

	// Parse the first argument to determine the subcommand
	if len(os.Args) < 2 {
		fmt.Println("expected 'upload', 'list', or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "upload":
		uploadCmd.Parse(os.Args[2:])
		if *uploadFilePath == "" {
			fmt.Println("Please provide a file path using the -file flag.")
			os.Exit(1)
		}
		err := cli.UploadFile(*uploadFilePath)
		if err != nil {
			fmt.Printf("Error uploading file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("File uploaded successfully.")
	case "list":
		err := cli.ListFiles()
		if err != nil {
			fmt.Printf("Error listing files: %v\n", err)
			os.Exit(1)
		}
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteFileName == "" {
			fmt.Println("Please provide a file name using the -name flag.")
			os.Exit(1)
		}
		err := cli.DeleteFile(*deleteFileName)
		if err != nil {
			fmt.Printf("Error deleting file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("File deleted successfully.")
	default:
		fmt.Println("expected 'upload', 'list', or 'delete' subcommands")
		os.Exit(1)
	}
}
