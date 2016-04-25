package main

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var AddCommand = &cobra.Command{
	Use: "add",
	Run: AddCommandHandler,
}

func AddCommandHandler(cmd *cobra.Command, args []string) {

	link := strings.Join(args, " ")

	bookmark := &Bookmark{
		Title:   LoadPage(link),
		Group:   group,
		Link:    link,
		Private: markPrivate,
		Tags:    tags,
	}
	AddBookmark(bookmark)
	log.Println("bookmark added")
}
