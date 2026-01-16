-- +goose Up
CREATE TABLE IF NOT EXISTS "chats" (
     "created_at" timestamptz,
     "title" varchar(200) NOT NULL,
     "id" bigserial,
     PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "messages" (
    "created_at" timestamptz,
    "text" varchar(5000) NOT NULL,
    "chat_id" bigint NOT NULL,
    "id" bigserial,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_messages_chat" FOREIGN KEY ("chat_id")
        REFERENCES "chats"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chats;