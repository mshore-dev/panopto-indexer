package database

type Video struct {
	ID int
	Name string
	Description string
	TagID int
	Tag string
}

func AddVideo(name, description, file string) (int, error) {
	resp, err := db.Exec("insert into videos(name, description, video_file) values (?,?,?)", name, description, file)
	if err != nil {
		return 0, err
	}

	lastVideoID, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastVideoID), nil

}

// func getVideoByID(id int) (err error, video Video)

// func findVideosByName(query string, tag int) (err error, videos []Video)

// func findVideosByDescription(query, tag int) (err error, videos []Video)

// func tagVideo(id, tag int)