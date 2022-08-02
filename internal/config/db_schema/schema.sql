create database  networkmanager;
use networkmanager;

create table users(
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);


create table roles (
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
);

create table users_roles (
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references users(id),
    foreign key (role_id) references roles(id)
);

create table plans (
    id int not null auto_increment,
    name varchar(255) not null,
    price int not null,
    rate varchar(255) not null,
    primary key (id)
);

create table router (
    id int not null auto_increment,
    ip varchar(20) not null,
    name varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);

create table clients (
    id int not null auto_increment,
    name varchar(255) not null,
    last_name varchar(255) not null,
    address varchar(255) not null,
    phone varchar(255) not null,
    payment_date int not null,
    plan_id int not null,
    router_id int not null,
    primary key (id),
    foreign key (plan_id) references plans(id),
    foreign key (router_id) references router(id)
);

create table invoices (
    id int not null auto_increment,
    name varchar(255) not null,
    last_name varchar(255) not null,
    address varchar(255) not null,
    phone varchar(255) not null,
    plan varchar(255) not null,
    month varchar(25) not null,
    amount int not null,
    creation_date date not null,
    limit_date date not null,
    payment_date date,
    paid boolean,
    primary key (id)
);