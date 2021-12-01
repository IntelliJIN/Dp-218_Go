CREATE TABLE IF NOT EXISTS roles
(
    id         smallint PRIMARY KEY,
    name       VARCHAR(50) UNIQUE,
    is_admin    boolean,
    is_user     boolean,
    is_supplier boolean
);

CREATE TABLE IF NOT EXISTS users
(
    id          serial PRIMARY KEY,
    login_email  VARCHAR(100) UNIQUE NOT NULL,
    is_blocked   boolean,
    user_name    VARCHAR(100),
    user_surname VARCHAR(100),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id      int                 NOT NULL

);

CREATE TABLE IF NOT EXISTS supplier_prices
(
    id               serial PRIMARY KEY,
    price            money,
    payment_type_id  smallint NOT NULL,
    user_id          int      NOT NULL

);

CREATE TABLE IF NOT EXISTS payment_types
(
    ID   smallserial PRIMARY KEY,
    Name VARCHAR(45) UNIQUE
);

CREATE TABLE IF NOT EXISTS scooter_models
(
    id               serial PRIMARY KEY,
    payment_type_id  smallint       NOT NULL,
    model_name       VARCHAR(45)   UNIQUE NOT NULL,
    max_weight       NUMERIC(5, 2),
    speed            smallint       NOT NULL
);

CREATE TABLE IF NOT EXISTS scooters
(
    id            serial PRIMARY KEY,
    model_id      smallint            NOT NULL,
    owner_id      int                 NOT NULL,
    serial_number VARCHAR(45) UNIQUE NOT NULL
);

BEGIN;
INSERT INTO roles(id, name, is_admin, is_user, is_supplier) VALUES(1, 'admin role', true, false, false);
INSERT INTO roles(id, name, is_admin, is_user, is_supplier) VALUES(2, 'user role', false, true, false);
INSERT INTO roles(id, name, is_admin, is_user, is_supplier) VALUES(3, 'supplier role', false, false, true);
INSERT INTO roles(id, name, is_admin, is_user, is_supplier) VALUES(7, 'super_admin role', true, true, true);

INSERT INTO users (id, login_email, is_blocked, user_name, user_surname, created_at, role_id) VALUES (1, 'first@gmail.com', false, 'Mihail', 'Galustian', '2020-06-22 19:10:25-07', 1 );
INSERT INTO users (id, login_email, is_blocked, user_name, user_surname, created_at, role_id) VALUES (2, 'seccond@gmail.com', false, 'Elena', 'Stepanenko', '2021-06-22 19:10:25-07', 1 );
INSERT INTO users (id, login_email, is_blocked, user_name, user_surname, created_at, role_id) VALUES (3, 'third@gmail.com', false, 'Fernando', 'Alonso', '2021-06-22 19:10:25-07', 1 );

INSERT INTO payment_types (id, name) VALUES (1, 'EUR');
INSERT INTO payment_types (id, name) VALUES (2, 'USD');
INSERT INTO payment_types (id, name) VALUES (3, 'GRN');

INSERT INTO supplier_prices (id, price, payment_type_id, user_id) VALUES (1, 2, 1, 3);
INSERT INTO supplier_prices (id, price, payment_type_id, user_id) VALUES (2, 1, 2, 2);
INSERT INTO supplier_prices (id, price, payment_type_id, user_id) VALUES (3, 8, 3, 1);

INSERT INTO scooter_models (id, payment_type_id,  model_name, max_weight, speed) VALUES (1, 3, 'Xiaomi Mi Scooter Pro2', 120, 30);
INSERT INTO scooter_models (id, payment_type_id,  model_name, max_weight, speed) VALUES (2, 3, 'Kugoo G2 Pro', 130, 45);
INSERT INTO scooter_models (id, payment_type_id,  model_name, max_weight, speed) VALUES (3, 3, 'NINEBOT BY SEGWAY MAX G30', 140, 45);
INSERT INTO scooter_models (id, payment_type_id,  model_name, max_weight, speed) VALUES (4, 3, 'OIO RT5 PRO Dual', 150, 60);

INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (1, 1, 3, 300001);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (2, 1, 3, 300002);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (3, 2, 3, 300003);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (4, 2, 3, 300004);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (5, 2, 3, 300005);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (6, 3, 3, 300006);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (7, 3, 3, 300007);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (8, 4, 1, 100001);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (9, 4, 1, 100002);
INSERT INTO scooters (id, model_id, owner_id, serial_number) VALUES (10, 4, 1, 100003);
COMMIT;

