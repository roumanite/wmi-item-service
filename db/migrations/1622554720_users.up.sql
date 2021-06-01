CREATE TABLE IF NOT EXISTS users
(
    id                         varchar(510),
    username                   varchar(255) unique,
    first_name                 varchar(510),
    last_name                  varchar(510),
    email                      varchar(510),
    birthdate                  date,
    email_verification_token   varchar(255) unique,
    email_verified_at          timestamptz,
    email_verification_sent_at timestamptz,
    encrypted_password      varchar(510),
    created_at              timestamptz,
    updated_at              timestamptz,
    deleted_at              timestamptz,
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);
