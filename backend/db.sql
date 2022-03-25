CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE tbl_room (
    id    UUID         DEFAULT uuid_generate_v4 (), 
    name  VARCHAR(64)  NOT NULL
);