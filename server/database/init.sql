DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS owner CASCADE;
DROP TABLE IF EXISTS patient CASCADE;
DROP TABLE IF EXISTS office CASCADE;
DROP TABLE IF EXISTS pathology CASCADE;
DROP TABLE IF EXISTS visitSheet CASCADE;
DROP TABLE IF EXISTS employee CASCADE;
DROP TABLE IF EXISTS chef CASCADE;
DROP TABLE IF EXISTS timePeriod CASCADE;



CREATE TABLE IF NOT EXISTS users(
   username TEXT PRIMARY KEY,
   password TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS pathology(
    id SERIAL,
    title VARCHAR(100) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS patient(
    id SERIAL PRIMARY KEY,
    officeId TEXT REFERENCES office(id),
    firstname VARCHAR(120) NOT NULL,
    lastname VARCHAR(120) NOT NULL,
    sex BOOLEAN NOT NULL,
    birth DATE NOT NULL,
    security TEXT,
    phone VARCHAR(20),
    city VARCHAR(100) NOT NULL,
    complement VARCHAR(120),
    addedThe DATE DEFAULT NOW(),
    patient_pathology VARCHAR(100) REFERENCES pathology(title),
    lastVisit DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS office (
    id TEXT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    complement VARCHAR(120) NOT NULL,
    owner_number INTEGER,
    number_patient INTEGER,
    addedThe DATE DEFAULT NOW()
);



CREATE TABLE IF NOT EXISTS owner(
  id TEXT PRIMARY KEY,
  officeId TEXT REFERENCES office(id), 
  lastname VARCHAR(120) NOT NULL,
  firstname VARCHAR(120) NOT NULL,
  birth DATE NOT NULL,
  email TEXT NOT NULL,
  upin VARCHAR(20),
  isManager BOOLEAN NOT NULL,
  addedThe DATE DEFAULT NOW()
);



CREATE TABLE IF NOT EXISTS chef (
   officeId TEXT REFERENCES office(id),
   fullName VARCHAR(120)
);

CREATE TABLE IF NOT EXISTS employee(
    officeId TEXT REFERENCES office(id),
    ownerId TEXT REFERENCES owner(id),
    addedThe DATE DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS visiteSheet (
    ownerId TEXT REFERENCES owner(id),
    editedBy VARCHAR(120),
    addedAt DATE DEFAULT NOW(),
    weight SMALLINT,
    glycemia FLOAT,
    pressure FLOAT,
    temperature FLOAT
);


CREATE TABLE IF NOT EXISTS  timePeriod (
    id SERIAL PRIMARY KEY,
    ownerId TEXT REFERENCES owner(id),
    day VARCHAR(20),
    time VARCHAR(20),
    position SMALLINT,
    fullname VARCHAR(120),
    city VARCHAR(120)
);
