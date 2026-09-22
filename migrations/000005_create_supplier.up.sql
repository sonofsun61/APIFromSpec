create table if not exists supplier (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    address_id uuid not null,
    phone_number text not null,

    constraint fk_address_id foreign key (address_id) references address(id)
);