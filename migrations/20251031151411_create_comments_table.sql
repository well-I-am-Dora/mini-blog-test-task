-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments
(
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT comments_pkey PRIMARY KEY,
    text TEXT NOT NULL CHECK (LENGTH(text) <= 2000),
    user_id UUID NOT NULL,
    post_id UUID NOT NULL,
    parent_comment_id UUID,

    path LTREE,

    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_parent_comment_id ON comments(parent_comment_id);

CREATE INDEX idx_comments_path_gist ON comments USING GIST (path);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd