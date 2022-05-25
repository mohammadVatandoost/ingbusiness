
CREATE TABLE saved_messages (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    image TEXT NOT NULL,
    organization_id INT NOT NULL,
    writer_id  INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(writer_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_organization
        FOREIGN KEY(organization_id)
            REFERENCES organization(id)
            ON DELETE CASCADE,
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

