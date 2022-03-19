package postgres

import (
	"context"
	"videoStreaming/hrm/storage"
)

const insertVideo = `
	INSERT INTO videos(
		title,
	    video_file,
	    is_active
	)VALUES(
		:title,
	        :video_file,
	        :is_active
	)RETURNING id;
`

func (s *Storage) CreateVideo(ctx context.Context, t storage.Video) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertVideo)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}
