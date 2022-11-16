package database

func addVideo(name, description string, tag int) (newVideoID int, err error)

func getVideoByID(id int)

func findVideoByName(query string, tag int)

func findVideoByDescription(query, tag int)

func tagVideo(id, tag int)