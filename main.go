package main

import (
    "log"
    "github.com/spf13/cobra"
    "encoding/json"
    "io/ioutil"
    "os"    
)

// Bookmark is the Key Model ...
type Bookmark struct {
    Group string `json:"group"`
    Link string `json:"link"`
    Private bool `json:"private"`
    Tags []string
}

var rootCmd = &cobra.Command{
    Use: "gomark",
    Short: "A Commandline tool to manage bookmarks.",
    Long: `A GoLang-based Commandlinetool to manage bookmarks.`,
    Run: func (cmd *cobra.Command, args []string) {
        log.Println("welcome to the main function")
    },
}

var bookmarks []*Bookmark

var dataModified = false

// GetBookmarks delivers all bookmarks of current context
func GetBookmarks(publicOnly bool) []*Bookmark {
    if !publicOnly {
        return bookmarks        
    }
    var withoutPrivate []*Bookmark
    for _, bookmark := range bookmarks {
        if !bookmark.Private {
            withoutPrivate = append(withoutPrivate, bookmark)
        }
    }
    return withoutPrivate
}

func createInitialData (filename string) {
    err := ioutil.WriteFile(filename, []byte("[]"), os.ModePerm)
    if err != nil {
        log.Println("Error creating file:", err)
        os.Exit(1)
    }
    log.Println("file does not exists. creating", filename)
}

func readData (filename string) {
    if _, err := os.Stat(filename); err != nil {
        createInitialData(filename)
    }
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Println("Error reading file:", err)
        os.Exit(1)
    }
    json.Unmarshal(file, &bookmarks)
}

func writeData (filename string) {
    data, err := json.Marshal(bookmarks)
    if err != nil {
        log.Println("Error writing file:", err)
        os.Exit(1)
    }
    ioutil.WriteFile(filename, data, os.ModePerm)
}

func AddBookmark(bookmark *Bookmark) {
    bookmarks = append(bookmarks, bookmark)
    dataModified = true
}

func main () {
    IntitalizeAddCommand(rootCmd)
    InitListCommand(rootCmd)
    readData("bookmarks.json")
    
    if err := rootCmd.Execute(); err != nil {
        panic(err)       
    }
    
    if dataModified {
        writeData("bookmarks.json")
    }
    
}