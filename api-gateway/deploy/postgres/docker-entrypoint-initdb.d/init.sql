CREATE TABLE address (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    street VARCHAR(255) NOT NULL,

    CONSTRAINT unique_address UNIQUE (country, city, street)
);

CREATE TABLE client (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    birthday DATE NOT NULL,
    gender CHAR(1) NOT NULL CHECK (gender IN ('M', 'F')),
    registration_date TIMESTAMP DEFAULT now(),
    address_id UUID REFERENCES address(id)
);

CREATE TABLE supplier (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR(255) NOT NULL,
    address_id UUID REFERENCES address(id),
    phone_number VARCHAR(20)
);

CREATE TABLE image (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    image BYTEA NOT NULL
);

CREATE TABLE product (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    price NUMERIC(10, 2) NOT NULL,
    available_stock INTEGER NOT NULL,
    last_update_date TIMESTAMP DEFAULT now(),
    supplier_id UUID REFERENCES supplier(id),
    image_id UUID REFERENCES image(id) ON DELETE SET NULL
);