-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE topics (
	id bigserial NOT NULL,
	uuid text NULL DEFAULT gen_random_uuid(),
	title varchar(255) NULL,
	value varchar(255) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT topics_pkey PRIMARY KEY (id),
	CONSTRAINT uni_topics_title UNIQUE (title),
	CONSTRAINT uni_topics_value UNIQUE (value)
);
CREATE INDEX idx_topics_deleted_at ON topics USING btree (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS topics;
-- +goose StatementEnd
