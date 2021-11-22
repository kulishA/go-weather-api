CREATE TABLE IF NOT EXISTS weather
(
    id           serial                                       not null unique,
    city_id      int references cities (id) on delete cascade not null unique,
    last_updated varchar(255)                                 not null,
    temp_c       float                                        not null,
    temp_f       float                                        not null,
    is_day       int                                          not null,
    wind_mph     float                                        not null,
    wind_kph     float                                        not null,
    wind_degree  int                                          not null,
    wind_dir     varchar(255)                                 not null,
    pressure_mb  float                                        not null,
    pressure_in  float                                        not null,
    precip_mm    float                                        not null,
    precip_in    float                                        not null,
    humidity     int                                          not null,
    cloud        int                                          not null,
    column_16    int                                          not null,
    feelslike_c  float                                        not null,
    feelslike_f  float                                        not null,
    vis_km       float                                        not null,
    vis_miles    float                                        not null,
    uv           float                                        not null,
    gust_mph     float                                        not null,
    gust_kph     float                                        not null,
    location_id  int                                          not null
);

