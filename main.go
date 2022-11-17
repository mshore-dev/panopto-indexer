package main

import (
	"os"
	"flag"

	"github.com/mshore-dev/panopto-indexer/database"
	"github.com/mshore-dev/panopto-indexer/importer"
	"github.com/mshore-dev/panopto-indexer/search"
)

func main() {

	// set up subcommands and their accompanying flags

	// import command for adding a video and accompanying subtitles to the database
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	importVideoFile := importCmd.String("video", "", "Video file to import")
	importSubFile := importCmd.String("subs", "", "Subtitle file to import")
	importName := importCmd.String("name", "", "Name of imported video")
	importDescription := importCmd.String("description", "", "Description of imported video")
	// importTags := importCmd.String("tags", "", "Tags to apply to imported video")

	// remove command to remove a video and accompanying subtitles from the database
	// removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	// removeVideoID := removeCmd.Int("id", 0, "Video ID to remove")

	// search command to search the database of subtitles for a given query
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchTag := searchCmd.String("tag", "", "Filter search by tag")
	searchQuery := searchCmd.String("query", "", "Query to find in video subtitles")

	// // find-video command for searching the database of videos
	// findVideoCmd := flag.NewFlagSet("find-video", flag.ExitOnError)
	// findVideoTag := findVideoCmd.String("tag", "", "Filter videos by tag")
	// findVideoName := searchCmd.String("name", "", "Search by video name")
	// findVideoDescription := searchCmd.String("description", "", "Search by video")

	// // get-tag command to get videos with givent tag
	// getTagCmd := flag.NewFlagSet("get-tag", flag.ExitOnError)
	// getTagTag := getTagCmd.String("tag", "", "Tag to get a video list for")

	// // create-tag command to create a new tag
	// createTagCmd := flag.NewFlagSet("create-tag", flag.ExitOnError)
	// createTagName := createTagCmd.String("name", "", "Name of tag to create")

	// // delete-tag command to remove a tag
	// deleteTagCmd := flag.NewFlagSet("delete-tag", flag.ExitOnError)
	// deleteTagName := deleteTagCmd.String("name", "", "Name of tag to delete")

	// // list-tags command to list tags currently in the database
	// listTagsCmd := flag.NewFlagSet("list-tags", flag.ExitOnError)


	err := database.OpenDatabase("panopto-indexer")
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "import":
		importCmd.Parse(os.Args[2:])
		importer.ImportVideo(*importName, *importDescription, *importVideoFile, *importSubFile)
	case "search":
		searchCmd.Parse(os.Args[2:])
		if *searchTag == "" {
			// no tag specified.
			search.Search(*searchQuery)
		}
	}

}