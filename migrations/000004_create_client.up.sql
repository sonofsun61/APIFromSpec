create type gender_enum as enum('male', 'female');

create table if not exists client (
    id uuid primary key default gen_random_uuid(),
    client_name text not null,
    client_surname text not null,
    birthday date not null,
    gender gender_enum,
    registration_date date not null default now(),
    address_id uuid not null,

    constraint fk_address_id foreign key (address_id) references address(id)
);