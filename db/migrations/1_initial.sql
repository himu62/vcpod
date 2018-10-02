-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table root_object (
    id int primary key,
    path text,
    registered text -- timestamp
);

create table object (
    id int primary key,
    path text,
    type text,
    hash text,
    hash_type text,
    registered text, -- timestamp
    last_used text, -- timestamp
    score int
);

create table object_used (
    object_id int,
    timestamp text
);

create table favorite (
    id int primary key,
    object_id int,
    registered text -- timestamp
);

create table tag (
    id int primary key,
    string text,
    registered text -- timestamp
);

create table object_tag(
    id int primary key,
    tag_id int,
    object_id int,
    registered text -- timestamp
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back