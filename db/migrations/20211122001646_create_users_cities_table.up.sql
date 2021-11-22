CREATE TABLE users_cities
(
    id      serial                                            not null unique,
    user_id int references users (id) on delete cascade       not null,
    city_id int references cities (city_id) on delete cascade not null
)