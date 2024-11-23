package main

import (
    "fmt"
    "log"
    "os"

    "github.com/go-git/go-git/v5"
    
)

func main() {
    repo := "https://github.com/koppulakoushik/version.git"
    cloneDir := `C:\Users\kkoppula\OneDrive - e2open, LLC\Documents\jenkins`

    // Remove all contents of the directory
    err := os.RemoveAll(cloneDir)
    if err != nil {
        log.Fatal(err)
    }

    // Clone the repository
    _, err = git.PlainClone(cloneDir, false, &git.CloneOptions{
        URL:      repo,
        Progress: os.Stdout,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Make changes to the file
    filepath := `C:\Users\kkoppula\OneDrive - e2open, LLC\Documents\jenkins\versionfiles.txt`
    fmt.Println("successfully cloned the repo")
    a := []string{"17"}
    b := []string{"jdk_x64_linux", "jre_x64_alpine","jre_ppc64le_linux"}
    combination(filepath, a, b)

   
}

func combination(filepath string, a []string, b []string) {
    file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for _, val := range a {
        for _, val2 := range b {
            content := fmt.Sprintf("%s,%s,%s\n", val, val2, "jdk-0.0.0")
            if _, err := file.WriteString(content); err != nil {
                log.Fatal("Error writing to file:", err)
            }
        }
    }

    log.Println("Content added successfully!")
}
