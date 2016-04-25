package main

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// Flag variables
var (
	showPrivate bool
	tags        []string
	group       string
	title       string
	port        string
	markPrivate bool
	filename    string
)

// Bookmark is the Key Model ...
type Bookmark struct {
	Title   string `json:"title"`
	Group   string `json:"group"`
	Link    string `json:"link"`
	Private bool   `json:"private"`
	Tags    []string
}

var GomarkCmd = &cobra.Command{
	Use:   "gomark",
	Short: "A Commandline tool to manage bookmarks.",
	Long:  `A GoLang-based Commandlinetool to manage bookmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("welcome to the main function")
	},
}

var bookmarks []*Bookmark

var dataModified = false

// GetBookmarks delivers all bookmarks of current context
func GetBookmarks(showPrivate bool) []*Bookmark {
	if showPrivate {
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

func CreateInitialData() {
	err := ioutil.WriteFile(filename, []byte("[]"), os.ModePerm)
	if err != nil {
		log.Println("Error creating file:", err)
		os.Exit(1)
	}
	log.Println("file does not exists. creating", filename)
}

func ReadData() {
	if _, err := os.Stat(filename); err != nil {
		CreateInitialData()
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &bookmarks)
}

func WriteData() {
	if !dataModified {
		return
	}
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

func InitCommands() {

	GomarkCmd.Flags().StringVarP(&filename, "file", "f", "bookmarks.json", "The bookmarks file. Default: bookmarks.json")

	GomarkCmd.AddCommand(AddCommand)
	AddCommand.Flags().BoolVarP(&markPrivate, "private", "P", false, "Also show private bookmarks.")
	AddCommand.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "tags")
	AddCommand.Flags().StringVarP(&group, "group", "g", "", "The Bookmark group")
	AddCommand.Flags().StringVarP(&filename, "file", "f", "bookmarks.json", "The bookmarks file. Default: bookmarks.json")
	AddCommand.Flags().StringVarP(&title, "title", "T", "", "The bookmarks title")

	GomarkCmd.AddCommand(ListCommand)
	ListCommand.Flags().BoolVarP(&showPrivate, "private", "P", false, "show private")
	ListCommand.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "tags")
	ListCommand.Flags().StringVarP(&filename, "file", "f", "bookmarks.json", "The bookmarks file. Default: bookmarks.json")

	GomarkCmd.AddCommand(ServeCommand)
	ServeCommand.Flags().StringVarP(&port, "port", "p", "8080", "The port to serve on.")

}

func main() {

	InitCommands()

	ReadData()
	defer WriteData()

	if err := GomarkCmd.Execute(); err != nil {
		panic(err)
	}

}
