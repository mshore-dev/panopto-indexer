package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func OpenDatabase(name string) (err error) {
	db, err = sql.Open("sqlite3", name + ".db")
	if err != nil {
		return
	}

	// setup tables, virtual fts tables, and triggers.
	_, err = db.Exec(`
	create table if not exists videos (id integer not null primary key, name text, description text, video_file text);
	create table if not exists subtitles (id integer not null primary key, video_id integer, start_time integer, text text);
	create table if not exists tags (id integer not null primary key, name text);
	create table if not exists video_tag_map (video integer, tag integer);
	
	create virtual table if not exists subtitles_fts using fts5(text, content_rowid=id);
	create virtual table if not exists videos_fts using fts5(name, description, content_rowid=id);

	create trigger if not exists subtitles_fts_insert after insert on subtitles
	begin
		insert into subtitles_fts (rowid, text) values (new.rowid, new.text);
	end;

	create trigger if not exists videos_fts_insert after insert on videos
	begin
		insert into videos_fts (rowid, name, description) values (new.rowid, new.name, new.description);
	end;
	`)
	if err != nil {
		return
	}

	return
}