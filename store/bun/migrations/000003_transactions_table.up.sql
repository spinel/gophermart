BEGIN;

CREATE TYPE transaction_type AS ENUM ('withdraw', 'refill');

CREATE TABLE IF NOT EXISTS transactions (
	id serial PRIMARY KEY not null,
    user_id integer not null,
    order_id integer not null,
    amount decimal(12,2) default '0.00',
    type transaction_type default 'refill',
    created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
    UNIQUE (user_id, order_id)
);

COMMIT;
