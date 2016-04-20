package main

import (
    "github.com/spf13/cobra"
    "log"
    "strings"
)

var AddCommand = &cobra.Command{
    Use: "add",
    Run: AddHandler,
}

var tags []string
var private bool
var group string

func AddHandler (cmd *cobra.Command, args []string) {
    bookmark := &Bookmark{
        Group: group,
        Link: strings.Join(args, " "),
        Private: private,
        Tags: tags,
    }
    AddBookmark(bookmark)
    log.Println("bookmark added")
}

func IntitalizeAddCommand(rootCmd *cobra.Command) {
    rootCmd.AddCommand(AddCommand)
    AddCommand.Flags().StringSliceVar(&tags, "tag", []string{}, "add this to add tags")
    AddCommand.Flags().BoolVar(&private, "private", false, "mark bookmark as private")
    AddCommand.Flags().StringVar(&group, "group", "", "add group")
}

