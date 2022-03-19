-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS videos(
        id SERIAL NOT NULL,
        title TEXT NOT NULL,
        video_file TEXT NOT NULL,
        is_active BOOLEAN DEFAULT false,
        PRIMARY KEY(id)
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS videos;
-- +goose StatementEnd
