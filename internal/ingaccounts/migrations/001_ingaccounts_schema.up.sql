
CREATE TABLE ing_accounts (
    id SERIAL PRIMARY KEY,
    name VARCHAR (64) NOT NULL,
    token VARCHAR (256),
    user_id INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT ing_accounts_name_unique UNIQUE (name),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

