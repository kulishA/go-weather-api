CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE weather
(
    id          serial                                      not null unique,
    user_id     int references users (id) on delete cascade not null,
    location    varchar(255)                                not null,
    temp_c      float4                                      not null,
    temp_f      float4                                      not null,
    is_day      int                                         not null,
    wind_mph    float4                                      not null,
    wind_kph    float4                                      not null,
    wind_degree float4                                      not null
)