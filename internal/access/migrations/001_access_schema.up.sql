
CREATE TABLE access (
    id SERIAL PRIMARY KEY,
    organization_id INT NOT NULL,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    CONSTRAINT fk_organization
        FOREIGN KEY(organization_id)
            REFERENCES organization(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_roles
        FOREIGN KEY(role_id)
            REFERENCES roles(id)
            ON DELETE CASCADE,
    CONSTRAINT access_unique UNIQUE (organization_id, user_id, role_id),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

