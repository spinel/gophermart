BEGIN;

CREATE TYPE order_status AS ENUM ('NEW', 'PROCESSING', 'INVALID', 'PROCESSED');

CREATE TABLE IF NOT EXISTS orders (
	id serial PRIMARY KEY not null,
    user_id integer not null,
    number bigint not null unique,
    status order_status default 'NEW',
    accural decimal(12,2) default '0.00',
	uploaded_at timestamp,
    created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

COMMIT;
