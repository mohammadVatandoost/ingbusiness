
CREATE TABLE ing_pages (
    id SERIAL PRIMARY KEY,
    name VARCHAR (64) NOT NULL,
    token VARCHAR (256) NOT NULL,
    organization_id INT NOT NULL,
    creator_id INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(creator_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_organization
        FOREIGN KEY(organization_id)
            REFERENCES organization(id)
            ON DELETE CASCADE,
    CONSTRAINT ing_pages_name_unique UNIQUE (name),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

