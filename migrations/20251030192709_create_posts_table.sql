-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts
(
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT posts_pkey PRIMARY KEY,
    text TEXT NOT NULL,
    user_id UUID NOT NULL,
    comments_allowed BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd