
CREATE TABLE experiment (
    id SERIAL PRIMARY KEY,
    name VARCHAR (64) NOT NULL,
    description TEXT,
    condition_id INT NOT NULL,
    condition_params TEXT NOT NULL,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT experiment_name_unique UNIQUE (name),
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

