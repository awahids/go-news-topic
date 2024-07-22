-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE news_topics (
	topic_id int8 NOT NULL,
	news_id int8 NOT NULL,
	CONSTRAINT news_topics_pkey PRIMARY KEY (topic_id, news_id)
);
ALTER TABLE news_topics ADD CONSTRAINT fk_news_topics_news FOREIGN KEY (news_id) REFERENCES news(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE news_topics ADD CONSTRAINT fk_news_topics_topic FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS news_topics;
-- +goose StatementEnd
