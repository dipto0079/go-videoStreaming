package postgres

import (
	"context"
	"testing"
	"videoStreaming/hrm/storage"

	//"github.com/google/go-cmp/cmp"
)

func TestCreateVideo(t *testing.T) {
	s := newTestStorage(t)

	var tests = []struct {
		name    string
		in      storage.Video
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_VIDEO_SUCCESS",
			in: storage.Video{
				Title:      "This is title",
				VideoFile:  "This is Video file",
				IsComplete: true,
			},
			want: 1,
		},
		{
			name: "CREATE_VIDEO_SUCCESS",
			in: storage.Video{
				Title:      "This is title 1",
				VideoFile:  "This is Video file 1",
				IsComplete: true,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateVideo(context.Background(), tt.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}