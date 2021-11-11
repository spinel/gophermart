BEGIN;

CREATE TABLE IF NOT EXISTS preorders (
	id serial PRIMARY KEY not null,
    number character varying not null unique,
    amount decimal(12,2) default '0.00',
    created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);

COMMIT;
