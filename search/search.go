package search

import (
	"fmt"

	"github.com/mshore-dev/panopto-indexer/database"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Search(query string) {
	err, results := database.SearchSubtitles(query)
	if err != nil {
		panic(err)
	}

	// log.Println(len(results))

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Video Name", "Timestamp", "Match"})

	for i := 0; i < len(results); i++ {
		tw.AppendRow(table.Row{results[i].VideoName, results[i].Timestamp, results[i].Text})
	}

	fmt.Printf("%s\n", tw.Render())
}