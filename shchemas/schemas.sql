drop schema if exists ken cascade;
create schema ken;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


create table ken.contract
(
    id          uuid primary key default uuid_generate_v4(),
    description text,
    amount      float8,
    periods     integer,
    date        timestamp
);

create table ken.payment
(
    id          uuid not null default uuid_generate_v4() primary key,
    id_contract uuid references ken.contract (id),
    amount      decimal,
    date        timestamp
);

insert into ken.contract
values ('09cfdad3-9263-470a-8231-74fa6d483913', 'Prestaçao de serviços escolares', 6000, 12, '2022-01-01T10:00:00');

insert into ken.payment
values (default, '09cfdad3-9263-470a-8231-74fa6d483913', 6000, '2022-01-05T10:00:00');
