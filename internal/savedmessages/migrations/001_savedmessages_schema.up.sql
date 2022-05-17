
CREATE TABLE saved_messages (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    ing_account_id INT NOT NULL,
    writer_id  INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(writer_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_ing_account
        FOREIGN KEY(ing_account_id)
            REFERENCES ing_accounts(id)
            ON DELETE CASCADE,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

