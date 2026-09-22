create table if not exists address (
    id uuid primary key default gen_random_uuid(),
    country text not null,
    city text not null,
    street text not null
);