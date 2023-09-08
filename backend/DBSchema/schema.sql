CREATE TABLE users (
   id UUID NOT NULL,
   usermane VARCHAR(254) NOT NULL,
   password VARCHAR(72) NOT NULL,
   created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
   updated_at INTEGER,
   CONSTRAINT users_id_pk PRIMARY KEY (id),
   CONSTRAINT users_username_uk UNIQUE (username)
);

COMMENT ON TABLE users IS 'Storage the users for client-manager';


CREATE TABLE routers (
   id UUID NOT NULL,
   name VARCHAR(25) NOT NULL,
   ip VARCHAR(25) NOT NULL,
   username VARCHAR(50) NOT NULL,
   password VARCHAR(100) NOT NULL,
   created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
   updated_at INTEGER,
   CONSTRAINT routers_id_pk PRIMARY KEY (id),
   CONSTRAINT routers_name_uk UNIQUE (name)
);

COMMENT ON TABLE routers IS 'Storage the routers for client-manager';

CREATE TABLE services (
	id UUID NOT NULL,
	name VARCHAR(128) NOT NULL,
	price NUMERIC(10,2) NOT NULL,
    rate VARCHAR(128) NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT services_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE services IS 'Storage the services for client-manager';


CREATE TABLE clients (
	id UUID NOT NULL,
	first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    address VARCHAR(128) NOT NULL,
    phone VARCHAR(128) NOT NULL,
    street VARCHAR(128) NOT NULL,
    userdata JSONB NOT NULL,
    router_id UUID NOT NULL,
    service_id UUID NOT NULL,
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT clients_id_pk PRIMARY KEY (id),
    CONSTRAINT clients_routers_id_fk FOREIGN KEY (router_id)
            REFERENCES routers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT clients_service_id_fk FOREIGN KEY (service_id)
            REFERENCES services (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE clients IS 'Storage the clients for client-manager';