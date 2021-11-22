CREATE TABLE IF NOT EXISTS cities
(
    id      serial       not null unique,
    city_id int          not null unique,
    name    varchar(255) not null,
    region  varchar(255) not null,
    country varchar(255) not null,
    lat     float4       not null,
    lon     float4       not null,
    url     varchar(255) not null
    )