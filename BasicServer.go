package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

///////////////////////////////////////////////////////////////////////////

func cloneRepo(url, dir string) (*git.Repository, error) {
	log.Println("cloning %s into %s", url, dir)
	/*auth, keyErr := publicKey()
	if keyErr != nil {
		return nil, keyErr
	}*/
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		Progress: os.Stdout,
		URL:      url,
		/*Auth:     auth,*/
	})
	if err != nil {
		if err == git.ErrRepositoryAlreadyExists {
			log.Println("repo was already cloned")
			return git.PlainOpen(dir)
		} else {
			log.Println("clone git repo error: %s", err)
			return nil, err
		}
	}
	return r, nil
}

func extractText(repoPath string) (string, error) {
	cmd := exec.Command("git2gpt", repoPath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func saveOutput(text, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func deleteFolder(folderPath string) error {
	return os.RemoveAll(folderPath)
}

////////////////////////////////////////////////////////////////////////////

func Convert_GIT2GPT(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	///////////////////////////
	outputFile := "C:/Users/MSI/Desktop/GIT2GPT_PROJECT/TEMP.txt"

	repoPath := "C:/Users/MSI/Desktop/GIT2GPT_PROJECT/TestFolder"
	_, err := cloneRepo(url, repoPath)
	if err != nil {
		fmt.Printf("Failed to clone repository: %s\n", err)
		os.Exit(1)
	}

	text, err := extractText(repoPath)
	if err != nil {
		fmt.Printf("Failed to extract text: %s\n", err)
		os.Exit(1)
	}

	err = saveOutput(text, outputFile)
	if err != nil {
		fmt.Printf("Failed to save output: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully extracted text and saved it into a file.")

	erro := deleteFolder(repoPath)
	if err != nil {
		fmt.Println("Error deleting folder:", erro)
	} else {
		fmt.Println("Folder deleted successfully")
	}

	/////////////////////////
	io.WriteString(w, text)
}

func main() {
	http.HandleFunc("/converter", Convert_GIT2GPT)
	http.ListenAndServe(":8080", nil)
}
