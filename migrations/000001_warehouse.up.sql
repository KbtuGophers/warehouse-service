CREATE TABLE IF NOT EXISTS warehouse.stores (
    id VARCHAR PRIMARY KEY ,
    is_active boolean DEFAULT false,
    merchant_id VARCHAR,
    name VARCHAR not null ,
    location VARCHAR,
    rating NUMERIC DEFAULT 0,
    currency_id VARCHAR,
    city_id VARCHAR,
    schedule_id VARCHAR,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES account.accounts (id),
    FOREIGN KEY (currency_id) REFERENCES warehouse.currency (id),
    FOREIGN KEY (city_id) REFERENCES warehouse.city (id),
    FOREIGN KEY (schedule_id) REFERENCES warehouse.schedule (id)
);

CREATE TABLE IF NOT EXISTS warehouse.inventories (
    id VARCHAR PRIMARY KEY ,
    is_available boolean DEFAULT false,
    store_id VARCHAR,
    product_id VARCHAR,
    quantity INT,
    quantity_min INT,
    quantity_max INT,
    price NUMERIC,
    price_special NUMERIC,
    price_previous NUMERIC,
    FOREIGN KEY (store_id) REFERENCES warehouse.stores (id),
    FOREIGN KEY (product_id) REFERENCES product.products (id)
);


CREATE TABLE IF NOT EXISTS warehouse.currency (
    id  VARCHAR PRIMARY KEY ,
    sign VARCHAR,
    decimals INT,
    prefix BOOLEAN
);

CREATE TABLE IF NOT EXISTS warehouse.schedule (
    id VARCHAR PRIMARY KEY ,
    is_active BOOLEAN,
    periods JSONB
);

CREATE TABLE IF NOT EXISTS warehouse.delivery (
    id VARCHAR PRIMARY KEY ,
    is_active BOOLEAN,
    periods JSONB,
    area JSONB
);

CREATE TABLE IF NOT EXISTS warehouse.city (
    id VARCHAR PRIMARY KEY ,
    name VARCHAR,
    geocenter VARCHAR
);
