-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS news (
  id bigserial NOT NULL,
	uuid text NULL DEFAULT gen_random_uuid(),
	title varchar(255) NULL,
	"content" text NULL,
	status_id int8 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT news_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_news_deleted_at ON news USING btree (deleted_at);

ALTER TABLE news ADD CONSTRAINT fk_news_status FOREIGN KEY (status_id) REFERENCES statuses(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS news;
-- +goose StatementEnd
