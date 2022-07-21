CREATE TABLE "roles" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar (50) NOT NULL,
                         "desc" varchar (255) NOT NULL
);

CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username"  varchar (50) NOT NULL,
                         "email" varchar (255) NOT NULL,
                         "first_name" varchar (50) NOT NULL,
                         "last_name" varchar (50) NOT NULL,
                         "password" varchar (128) NOT NULL,
                         "delete_status" smallint NOT NULL,
                         "roles_id" bigserial NOT NULL,
                         CONSTRAINT roles_fk FOREIGN KEY("roles_id") REFERENCES roles("id")
);

CREATE TABLE "permissions" (
                               "id" bigserial PRIMARY KEY,
                               "title" varchar NOT NULL
);

CREATE TABLE "roles_permissions" (
                                     "roles_id"    int REFERENCES roles ("id") ON UPDATE CASCADE ON DELETE CASCADE,
                                     "permissions_id" int REFERENCES permissions ("id") ON UPDATE CASCADE,
                                     CONSTRAINT role_permission_pk PRIMARY KEY ("roles_id", "permissions_id")
);

CREATE TABLE "input_types" (
                               "id" bigserial PRIMARY KEY,
                               "name" varchar (50) NOT NULL
);

CREATE TABLE "inputs" (
                          "id" bigserial PRIMARY KEY,
                          "title" varchar (50) NOT NULL,
                          "input_types_id" bigserial NOT NULL ,
                          CONSTRAINT input_types_fk FOREIGN KEY("input_types_id") REFERENCES input_types("id")
);

CREATE TABLE "items" (
                         "id" bigserial PRIMARY KEY,
                         "type_id" bigserial NOT NULL,
                         "lon" decimal (7,4),
                         "lat" decimal (7,4),
                         "broken_id" bigserial,
                         "delete_status" smallint NOT NULL,
                         CONSTRAINT type_id_fk FOREIGN KEY("type_id") REFERENCES inputs("id"),
                         CONSTRAINT broken_id_fk FOREIGN KEY("broken_id") REFERENCES inputs("id")
);

CREATE TABLE "rents" (
                         "id" bigserial PRIMARY KEY,
                         "created_at" timestamp NOT NULL,
                         "date" date NOT NULL,
                         "external_id" bigserial NOT NULL,
                         "items_id" bigserial,
                         "users_id" bigserial,
                         "delete_status" smallint NOT NULL,
                         CONSTRAINT items_fk FOREIGN KEY("items_id") REFERENCES items("id"),
                         CONSTRAINT users_fk FOREIGN KEY("users_id") REFERENCES users("id")
);

INSERT INTO
    "roles" ("title", "desc")
VALUES
    ('Admin', 'Vlasnik poduzeća'),
    ('Radnik', 'Sezonac');

INSERT INTO
    "permissions" ("title")
VALUES
    ('CRUD users'),
    ('CRUD items'),
    ('User settings'),
    ('Personal update');

INSERT INTO
    "roles_permissions" ("roles_id", "permissions_id")
VALUES
    (1,1),
    (1,2),
    (1,3),
    (2,3),
    (2,4);

INSERT INTO
    "input_types" ("name")
VALUES
    ('type'),
    ('broken');

INSERT INTO
    "inputs" ("title","input_types_id")
VALUES
    ('sunbed', 1),
    ('parsol', 1),
    ('light damaged', 2),
    ('medium damaged', 2),
    ('completely broken', 2);

ALTER TABLE "users" ALTER COLUMN "delete_status" SET DEFAULT 1;

ALTER TABLE "items" ALTER COLUMN "delete_status" SET DEFAULT 1;
ALTER TABLE "items" ALTER COLUMN "broken_id" DROP NOT NULL;
ALTER TABLE "items" ALTER COLUMN "broken_id" DROP DEFAULT;

ALTER TABLE "rents" ALTER COLUMN "external_id" DROP DEFAULT;
ALTER TABLE "rents" ALTER COLUMN "items_id" DROP DEFAULT;
ALTER TABLE "rents" ALTER COLUMN "users_id" DROP DEFAULT;
ALTER TABLE "rents" ALTER COLUMN "created_at" SET DEFAULT NOW();
ALTER TABLE "rents" ALTER COLUMN "date" SET DEFAULT CURRENT_DATE;
ALTER TABLE "rents" ALTER COLUMN "delete_status" SET DEFAULT 1;


INSERT INTO
    users ("id", "username", "email", "first_name", "last_name", "password", "delete_status", "roles_id")
VALUES
    (1, 'admin', 'neki@gmail.com', 'admin', 'hotela', '$2a$12$KSZqnYNviKkCebKV1j9Ai.aBdU3Yym7ao06XDDZoM6GedNQIQ5O1W', 1, 1)