CREATE sequence positions_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE IF NOT EXISTS positions
(
    id                  int8 not null default nextval('positions_id_seq'),
    name                varchar(510),
    residence_id        varchar(510),
    position_id_parent  varchar(510),
    created_at          timestamptz,
    updated_at          timestamptz,
    deleted_at          timestamptz,
    CONSTRAINT "positions_pkey" PRIMARY KEY ("id")
);