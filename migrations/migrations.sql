create table if not exists images (
    id uuid primary key default gen_random_uuid(),
    image_id bytea not null
);

create table if not exists address (
    id uuid primary key default gen_random_uuid(),
    country text not null,
    city text not null,
    street text not null
);

create table if not exists client (
    id uuid primary key default gen_random_uuid(),
    client_name text not null,
    client_surname text not null,
    birthday date not null,
    gender enum('male', 'female'),
    registration_date date not null default now(),
    address_id uuid not null,

    constraint fk_address_id foreign key (address_id) references address(id)
);

create table if not exists supplier (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    address_id uuid not null,
    phone_number text not null,

    constraint fk_address_id foreign key (address_id) references address(id)
);

create table if not exists category (
    id uuid primary key default gen_random_uuid(),
    name text not null
)

create table if not exists product (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    category_id uuid not null,
    price decimal(10, 2) not null,
    available_stock int not null,
    last_update_date date default now(),
    supplier_id uuid not null,
    image_id uuid not null,

    constraint fk_category_id foreign key (category_id) references category(id),
    constraint fk_supplier_id foreign key (supplier_id) references supplier(id),
    constraint fk_image_id foreign key (image_id) references images(id)
);