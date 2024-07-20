-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE news_topics (
	news_id int8 NOT NULL,
	topic_id int8 NOT NULL,
	CONSTRAINT news_topics_pkey PRIMARY KEY (news_id, topic_id)
);
-- news_topics foreign keys
ALTER TABLE news_topics ADD CONSTRAINT fk_news_topics_news FOREIGN KEY (news_id) REFERENCES news(id);
ALTER TABLE news_topics ADD CONSTRAINT fk_news_topics_topic FOREIGN KEY (topic_id) REFERENCES topics(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS news_topics;
-- +goose StatementEnd
