
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR (128) NOT NULL,
    username VARCHAR (128) NOT NULL,
    email VARCHAR (128) NOT NULL,
    phone VARCHAR (128) NOT NULL,
    password VARCHAR (128) NOT NULL,
    profile_image TEXT NOT NULL,
    CONSTRAINT username_unique UNIQUE (username),
    CONSTRAINT email_unique UNIQUE (email),
    CONSTRAINT phone_unique UNIQUE (phone),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

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



CREATE TABLE access (
                        id SERIAL PRIMARY KEY,
                        organization_id INT NOT NULL,
                        organization_name VARCHAR (128) NOT NULL, -- It is redundancy, but it improve performance
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



