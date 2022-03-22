CREATE SEQUENCE user_id MINVALUE 1 INCREMENT 1;

CREATE TABLE users(
  id integer NOT NULL DEFAULT nextval('user_id'),
  email varchar(100),
  password varchar(100),
  apikey varchar(200),
  PRIMARY KEY(id)
);

CREATE SEQUENCE device_id MINVALUE 1 INCREMENT 1;

CREATE TABLE devices(
  id integer NOT NULL DEFAULT nextval('device_id'),
  user_id integer,
  name varchar(50),
  PRIMARY KEY(id)
);

CREATE INDEX idx_device_user ON devices (user_id);

ALTER TABLE devices
 ADD CONSTRAINT idx_device_user FOREIGN KEY ( user_id ) REFERENCES users(id);

CREATE SEQUENCE device_info_id MINVALUE 1 INCREMENT 1;

CREATE TABLE device_info(
  id integer NOT NULL DEFAULT nextval('device_info_id'),
  device_id integer,
  temperature float,
  humidity float,
  soil_moisture float,
  timestamp bigint,
  PRIMARY KEY(id)
);

CREATE INDEX idx_device_info_dev ON device_info (device_id);

ALTER TABLE device_info
 ADD CONSTRAINT idx_device_info_dev FOREIGN KEY ( device_id ) REFERENCES devices(id);
