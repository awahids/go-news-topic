-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE news (
	id bigserial NOT NULL,
	uuid text NULL DEFAULT gen_random_uuid(),
	title varchar(255) NULL,
	"content" text NULL,
	status varchar(50) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT news_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_news_deleted_at ON news USING btree (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS news;
-- +goose StatementEnd
