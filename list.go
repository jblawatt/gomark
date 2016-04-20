package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
    Use: "list",
    Aliases: []string{"ls"},
    Run: ListHandler,
}

var privateLs bool

func ListHandler (cmd *cobra.Command, args []string) {
    data := GetBookmarks(!privateLs)
    
    for i, bookmark := range data {
        fmt.Println(i, bookmark.Link)
    }
    
    if len(data) == 0 {
        fmt.Println("no bookmarks")
    }
}

func InitListCommand(rootCmd *cobra.Command) {
    rootCmd.AddCommand(ListCommand)
    ListCommand.Flags().BoolVar(&privateLs, "private", false, "show private")
}
