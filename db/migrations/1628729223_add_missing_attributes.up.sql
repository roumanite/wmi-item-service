ALTER TABLE item_position_histories ADD COLUMN item_id int8;
ALTER TABLE items ADD COLUMN is_favorite boolean not null default false;