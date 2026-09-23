create table if not exists images (
    id uuid primary key default gen_random_uuid(),
    image bytea not null
);