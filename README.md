# GIT2GPT Converter

This Go application clones a Git repository, extracts text using `git2gpt`, saves the text to a file, and deletes the repository folder. It provides a web interface to perform these actions.

## Features
- Clone a Git repository
- Extract text using `git2gpt`
- Save extracted text to a file
- Delete the cloned repository folder

## Usage

1. **Clone Repository**: Clones the specified Git repository into a temporary directory.
2. **Extract Text**: Uses `git2gpt` to extract text from the cloned repository.
3. **Save Output**: Saves the extracted text to a specified file.
4. **Clean Up**: Deletes the temporary directory.

## Endpoints

- **POST** `/converter?url=<repo_url>`: Starts the conversion process. Replace `<repo_url>` with the URL of the Git repository.

## Setup

1. **Install Dependencies**: Ensure you have Go installed and set up. Install the `go-git` package using:
   ```bash
   go get github.com/go-git/go-git/v5

2. **Build and Run:**
   ```bash
   go build -o git2gpt
   ./git2gpt
   ```

3. **Access the Converter:**
   Open your browser and navigate to http://localhost:8080/converter?url=<repo_url>.


## License
This project is licensed under the MIT License. See the LICENSE file for details.
