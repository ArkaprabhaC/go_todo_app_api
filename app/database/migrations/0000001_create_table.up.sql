

CREATE TABLE IF NOT EXISTS note (
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255),
    description VARCHAR(255),
)