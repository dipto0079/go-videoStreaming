package storage

type Video struct {
	ID         int64  `db:"id"`
	Title      string `db:"title"`
	VideoFile  string `db:"video_file"`
	IsComplete bool   `db:"is_active"`
}
