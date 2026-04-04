package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const baseRawURL = "https://raw.githubusercontent.com/01-edu/public/master/subjects"

func ask(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	if answer == "" {
		fmt.Println("Error: answer cannot be empty.")
		os.Exit(1)
	}
	return answer
}

func prompt(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(answer)
}

func fetch(url string) (string, bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return "", false
	}
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Error: HTTP %d for %s\n", resp.StatusCode, url)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}
	return string(body), true
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", path, err)
		os.Exit(1)
	}
}

func main() {
	// Always run from the repository root: workspaceDir = current working directory
	workspaceDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	// 1. Project name
	projectName := ask("Enter the exact name of the new project: ")

	// 2. Sub-project?
	subAnswer := strings.ToLower(prompt("Is this a sub-project? (yes/no): "))
	projectPath := projectName
	if subAnswer == "yes" || subAnswer == "y" {
		parentProject := ask("Enter the exact name of the parent project: ")
		projectPath = parentProject + "/" + projectName
	}

	// ── Project Instructions ──────────────────────────────────────────────
	projectURL := fmt.Sprintf("%s/%s/README.md", baseRawURL, projectPath)
	fmt.Printf("\nFetching project instructions from:\n  %s\n", projectURL)

	content, ok := fetch(projectURL)
	if !ok {
		fmt.Printf("Error: nothing found at %s\n", projectURL)
		fmt.Println("Check that the project name is spelled correctly or check that the project name in Github is the same as the Intra's one.")
		os.Exit(1)
	}
	writeFile(filepath.Join(workspaceDir, "Project Instructions.md"), content)
	fmt.Println("Project Instructions.md  ✓")

	// ── Audit Instructions ────────────────────────────────────────────────
	auditURL := fmt.Sprintf("%s/%s/audit/README.md", baseRawURL, projectPath)
	fmt.Printf("\nFetching audit instructions from:\n  %s\n", auditURL)

	auditContent, ok := fetch(auditURL)
	if !ok {
		fmt.Printf("Warning: nothing found at %s\n", auditURL)
		fmt.Println("The audit page may not exist for this project or the project name in Github is different from the Intra's one.")
	} else {
		writeFile(filepath.Join(workspaceDir, "Audit Instructions.md"), auditContent)
		fmt.Println("Audit Instructions.md  ✓")
	}

	fmt.Println("\nSetup complete!")
}
