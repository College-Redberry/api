-- Create Schemas
CREATE SCHEMA account;

CREATE SCHEMA task_management;

CREATE SCHEMA communication;

-- Create Tables in Schema: account
CREATE TABLE account."users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    profile_image TEXT
);

-- Create Tables in Schema: task_management
CREATE TABLE task_management.projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE task_management.boards (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    manager_id SMALLINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    project_id SMALLINT,
    FOREIGN KEY (manager_id) REFERENCES account."user"(id),
    FOREIGN KEY (project_id) REFERENCES task_management.project(id)
);

CREATE TABLE task_management.statuses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    color CHAR(6) NOT NULL
);

CREATE TABLE task_management.priorities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    color CHAR(6) NOT NULL
);

CREATE TABLE task_management.cards (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    start_at TIMESTAMP,
    updated_at TIMESTAMP,
    finished_at TIMESTAMP,
    estimated_finished_at TIMESTAMP,
    status_id SMALLINT,
    manager_id SMALLINT,
    assigned_id SMALLINT,
    priority_id SMALLINT,
    parent_card_id SMALLINT,
    FOREIGN KEY (status_id) REFERENCES task_management.status(id),
    FOREIGN KEY (manager_id) REFERENCES account."user"(id),
    FOREIGN KEY (assigned_id) REFERENCES account."user"(id),
    FOREIGN KEY (priority_id) REFERENCES task_management.priority(id),
    FOREIGN KEY (parent_card_id) REFERENCES task_management.card(id)
);

-- Create Tables in Schema: communication
CREATE TABLE communication.messages (
    id SERIAL PRIMARY KEY,
    card_id SMALLINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    user_id SMALLINT,
    parent_message_id SMALLINT,
    FOREIGN KEY (card_id) REFERENCES task_management.card(id),
    FOREIGN KEY (user_id) REFERENCES account."user"(id),
    FOREIGN KEY (parent_message_id) REFERENCES communication.messages(id)
);

CREATE TABLE communication.history (
    id SERIAL PRIMARY KEY,
    card_id SMALLINT,
    user_id SMALLINT,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_before VARCHAR(50),
    status_after VARCHAR(50),
    FOREIGN KEY (card_id) REFERENCES task_management.card(id),
    FOREIGN KEY (user_id) REFERENCES account."user"(id)
);
