CREATE sequence residences_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE IF NOT EXISTS residences
(
    id             int8 not null default nextval('residences_id_seq'),
    user_id_owner  varchar(510) not null,
    nickname       varchar(255),
    street_address varchar(510),
    city           varchar(510),
    state          varchar(510),
    country        varchar(510),
    zip_code       varchar(510),
    building_name  varchar(510),
    created_at     timestamptz,
    updated_at     timestamptz,
    deleted_at     timestamptz,
    CONSTRAINT "residences_pkey" PRIMARY KEY ("id")
);