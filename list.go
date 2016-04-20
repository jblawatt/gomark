package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
    Use: "list",
    Aliases: []string{"ls"},
    Run: ListCommandHandler,
}

func FilterBookmarks() []*Bookmark {
    var result []*Bookmark
    
    checkIn := func (t string, bt []string) bool {
        for _, bts := range bt {
            if t == bts {
                return true
            }
        }
        return false
    }
    
    if len(tags) == 0 {
        // Return all Bookmarks if not tags to filter
        return GetBookmarks(showPrivate)
    }
    
    for _, bookmark := range GetBookmarks(showPrivate) {
        for _, tag := range tags {
            if checkIn(tag, bookmark.Tags) {
                result = append(result, bookmark)                   
            }
        }
    }
    return result
}

func ListCommandHandler (cmd *cobra.Command, args []string) {
    data := FilterBookmarks()
    
    for i, bookmark := range data {
        fmt.Println(i, bookmark.Link)
    }
    
    if len(data) == 0 {
        fmt.Println("no bookmarks")
    }
}

