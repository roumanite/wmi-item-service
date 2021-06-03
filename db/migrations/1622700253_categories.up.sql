CREATE sequence categories_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE IF NOT EXISTS categories
(
    id                      int8 not null default nextval('categories_id_seq'),
    name                    varchar(510),
    created_at              timestamptz,
    updated_at              timestamptz,
    deleted_at              timestamptz,
    CONSTRAINT "categories_pkey" PRIMARY KEY ("id")
);