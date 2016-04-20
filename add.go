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

func AddCommandHandler (cmd *cobra.Command, args []string) {
    bookmark := &Bookmark{
        Group: group,
        Link: strings.Join(args, " "),
        Private: markPrivate,
        Tags: tags,
    }
    AddBookmark(bookmark)
    log.Println("bookmark added")
}

