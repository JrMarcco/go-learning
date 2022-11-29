create table laptop
(
    id            bigint unsigned auto_increment comment 'id',
    uid           varchar(32)   not null default '',
    brand         varchar(32)   not null default '',
    laptop_name   varchar(64)   not null default '',
    weight        decimal(8, 2) not null default 0.00,
    price_rmb     int           not null default 0,
    released_year int           not null default 0,
    created_at    datetime      not null default now(),
    updated_at    datetime      not null default now() on update now(),
    primary key (id) using btree
) engine = innodb comment = 'account';

create index brand_index on laptop (brand);
create unique index unique_uid_index on laptop(uid);