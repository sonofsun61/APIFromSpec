create table if not exists product (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    category_id uuid not null,
    price decimal(10, 2) not null check(price >= 0),
    available_stock int not null check(available_stock >= 0),
    last_update_date date default now(),
    supplier_id uuid not null,
    image_id uuid,

    constraint fk_category_id foreign key (category_id) references category(id),
    constraint fk_supplier_id foreign key (supplier_id) references supplier(id),
    constraint fk_image_id foreign key (image_id) references images(id) on delete set null
);