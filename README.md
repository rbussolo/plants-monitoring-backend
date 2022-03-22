## Plants Monitoring - Backend

Created a backend resposability to store and provider data about the plants!

### How to works

In the path of project you have to install dependencies

`$ go get ./...`

To be honest I do not use this command.... so if do not works, tell me.

### Starting the server

To start the server (It is gonna run at port 9090).

<pre>
	cd .\src\
	go run .\main.go
</pre>

### Database

The database used is POSTGRES, maybe you have to edit "database.go" with yours credentials.

<pre>
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

-- Password is "123123"
INSERT INTO users (email, password, apikey) VALUES('rbussolo91@gmail.com', 'lsrjXOipsCRBeL8o5JZsLOG4OFcjqWprg4hYzdbKCh4=', '26a9ffa4-c373-4e2e-84a4-24561db63da5');
INSERT INTO users (email, password, apikey) VALUES('carlos@gmail.com', 'lsrjXOipsCRBeL8o5JZsLOG4OFcjqWprg4hYzdbKCh4=', 'afbbfb84-b3e5-462e-a38b-ccaf5bab5440');
INSERT INTO users (email, password, apikey) VALUES('test_authenticator@test.com', 'lsrjXOipsCRBeL8o5JZsLOG4OFcjqWprg4hYzdbKCh4=', '655951a6-cf19-4196-a0f7-9a071889431c');
</pre>
