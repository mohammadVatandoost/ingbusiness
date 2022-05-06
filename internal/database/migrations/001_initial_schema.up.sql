
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    username VARCHAR (128) NOT NULL,
    email VARCHAR (128) NOT NULL,
    phone VARCHAR (128) NOT NULL,
    password VARCHAR (128) NOT NULL,
    CONSTRAINT username_unique UNIQUE (username),
    CONSTRAINT email_unique UNIQUE (email),
    CONSTRAINT phone_unique UNIQUE (phone),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ing_accounts (
    id SERIAL PRIMARY KEY,
    username VARCHAR (128) NOT NULL,
    user_id INT,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES Users(id)
            ON DELETE CASCADE,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

