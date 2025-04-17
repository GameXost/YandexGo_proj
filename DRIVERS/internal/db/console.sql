create table drivers (
    id varchar(32),
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(50) not null,
    phone_number varchar(20) not null,
    driver_license varchar(10) not null,
    driver_license_date date not null,
    car_number varchar(10) not null,
    car_model varchar(20) not null,
    car_marks varchar(20) not null,
    car_color varchar(20) not null
)