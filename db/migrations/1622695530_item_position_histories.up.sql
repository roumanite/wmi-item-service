CREATE sequence item_position_histories_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE IF NOT EXISTS item_position_histories
(
    id                      int8 not null default nextval('item_position_histories_id_seq'),
    user_id_owner           varchar(510),
    position_id             int8,
    latest_picture_url      varchar(510),
    deletion_notes          varchar(510),
    created_at              timestamptz,
    updated_at              timestamptz,
    deleted_at              timestamptz,
    CONSTRAINT "item_position_histories_pkey" PRIMARY KEY ("id")
);