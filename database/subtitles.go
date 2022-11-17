package database

type SubtitleLine struct {
	VideoName string
	VideoFile string
	Timestamp int
	Text string
}

func AddSubtitleLine(videoID, timestamp int, text string) (err error) {
	_, err = db.Exec("insert into subtitles(video_id, start_time, text) values (?,?,?)", videoID, timestamp, text)
	return
}

func SearchSubtitles(query string) (error, []SubtitleLine) {
	// select start_time, text from subtitles where rowid in (select rowid from subtitles_fts where subtitles_fts match 'windows OR linux' order by rank) and video_id=1 limit 5;
	// select name, video_file, start_time, text from subtitles inner join videos where subtitles.id in (select rowid from subtitles_fts where subtitles_fts match 'windows' order by rank);
	res, err := db.Query("select name, video_file, start_time, text from subtitles inner join videos on subtitles.video_id = videos.id where subtitles.id in (select rowid from subtitles_fts where subtitles_fts match ? order by rank);", query)
	if err != nil {
		return err, []SubtitleLine{}
	}
	defer res.Close()

	var results []SubtitleLine

	for res.Next() {
		var r SubtitleLine
		err := res.Scan(&r.VideoName, &r.VideoFile, &r.Timestamp, &r.Text)
		if err != nil {
			return err, []SubtitleLine{}
		}

		// log.Printf("[database/subtitles.go] %s\n", r.Text)

		results = append(results, r)
	}

	return nil, results
}

// func searchSubtitlesTag(query string, tag int) (error, []SubtitleLine)

// func removeAllFromVideo(id int)