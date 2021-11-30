/*
INSERT INTO users (id, login_email, is_blocked, user_name, user_surname, created_at, role_id)

VALUES (2, 'f1@gmail.com', false, 'Zamik', 'Galus', '2016-06-22 19:10:25-07', 1 );

INSERT INTO payment_types (id, name)
VALUES (2, 'Dollar');

INSERT INTO scooter_models (id, payment_type_id,  model_name, max_weight, speed)
VALUES (2, 2, 'Razor', 150, 25);

INSERT INTO scooters (id, model_id, owner_id, serial_number)
VALUES (2, 2, 2, 020002);
*/

INSERT INTO locations (id, latitude, longitude, label)
VALUES (1, 4, 4, 'locationLabel1');

INSERT INTO scooter_stations (id, location_id, name, is_active)
VALUES (1, 1, 'Station1', true);


/*
CREATE TABLE IF NOT EXISTS scooter_stations
(
    id          serial PRIMARY KEY,
    location_id int NOT NULL,
    name        VARCHAR(100),
    is_active   boolean,

    FOREIGN KEY (location_id) REFERENCES locations (id)
);

CREATE TABLE IF NOT EXISTS locations
(
    id        serial PRIMARY KEY,
    latitude  NUMERIC(10, 0) NOT NULL,
    longitude NUMERIC(10, 0) NOT NULL,
    label     VARCHAR(200)
);
 */