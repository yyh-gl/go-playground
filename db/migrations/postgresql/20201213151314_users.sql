CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR NOT NULL,
    age INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,  
    PRIMARY KEY (id)
);

INSERT INTO users (name, age)
  VALUES
    ('yusuke', 20),
    ('hanako', 21);
