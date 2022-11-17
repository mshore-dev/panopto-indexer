package importer

import (
	"log"

	"github.com/asticode/go-astisub"
	"github.com/mshore-dev/panopto-indexer/database"
)

func ImportVideo(name, description, videoFile, subtitleFile string) {
	log.Printf("[importer] importing %s and %s...\n", videoFile, subtitleFile)

	id, err := database.AddVideo(name, description, videoFile)
	if err != nil {
		panic(err)
	}

	log.Printf("[importer] imported video '%s' has ID: %d\n", name, id)

	subs, err := astisub.OpenFile(subtitleFile)
	if err != nil {
		panic(err) 
	}

	log.Printf("[importer] importing %d subtitle lines...\n", len(subs.Items))

	for i := 0; i < len(subs.Items); i++ {
		database.AddSubtitleLine(id, int(subs.Items[i].StartAt.Seconds()), subs.Items[i].String())
	}

	log.Printf("[importer] imported %d lines for video %s\n", len(subs.Items), name)

	// TODO: tag newly imported video, if applicable
}