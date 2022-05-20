
CREATE TABLE organization (
    id SERIAL PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    owner_id INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(owner_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT name_owner_unique UNIQUE (name, owner_id),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

