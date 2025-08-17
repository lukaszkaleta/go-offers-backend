CREATE TABLE if not exists users (
    id serial primary key,
    person_first_name text not null default '',
    person_last_name text not null default '',
    person_email text not null default '',
    person_phone text not null default '',
    address_line1 text not null default '',
    address_line2 text not null default '',
    address_city       text not null default '',
    address_postal_code text not null default '',
    address_district   text not null default '',
    settings_radar_perimeter int not null default 10000,
    settings_radar_position_latitude int not null default 0,
    settings_radar_position_longitude int not null default 0
);

CREATE TABLE if not exists offer (
    id serial primary key,
    description_value text not null default '',
    description_image_url text not null default '',
    address_line1 text not null default '',
    address_line2 text not null default '',
    address_city       text not null default '',
    address_postal_code text not null default '',
    address_district   text not null default '',
    position_latitude int not null default 0,
    position_longitude int not null default 0,
    price_value int not null default 0,
    price_currency text not null default 'NOK'
);

CREATE TABLE if not exists user_offer (
    user_id bigint not null references users,
    offer_id bigint not null references offer
);