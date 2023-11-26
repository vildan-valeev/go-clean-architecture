DROP TABLE IF EXISTS items, categories cascade;

create table categories
(
    "id"            uuid    not null unique,
    "title"         text    not null,
    "description"   text     not null,
    "tag"           text    not null
);

comment on table categories is 'Список категорий';
comment on column categories.id is 'id категории';
comment on column categories.title is 'Имя категории';
comment on column categories.description is 'Описание категории';
comment on column categories.tag is 'Тег категории';

--------------------------------
create table items
(
    "id"            uuid    not null unique,
    "title"         text    not null,
    "amount"        int     not null,
    "category_id"   uuid    not null
);

comment on table items is 'Список записей';
comment on column items.id is 'id записи';
comment on column items.title is 'Имя записи';
comment on column items.amount is 'Сумма записи';
comment on column items.category_id is 'Категория записи';
