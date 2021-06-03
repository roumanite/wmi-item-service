CREATE sequence items_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE IF NOT EXISTS items
(
    id                      int8 not null default nextval('items_id_seq'),
    name                    varchar(510),
    user_id_owner           varchar(510),
    category_id             varchar(510),
    display_picture_url     varchar(510),
    notes                   varchar(510),
    created_at              timestamptz,
    updated_at              timestamptz,
    deleted_at              timestamptz,
    CONSTRAINT "items_pkey" PRIMARY KEY ("id")
);