
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    organization_id INT NOT NULL,
    creator_id INT NOT NULL,
    role_type INT NOT NULL,
    CONSTRAINT fk_organization
        FOREIGN KEY(organization_id)
            REFERENCES organization(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY(creator_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT roles_unique UNIQUE (organization_id, role_type),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

