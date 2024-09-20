-- Create Schemas
CREATE SCHEMA IF NOT EXISTS account;

CREATE SCHEMA IF NOT EXISTS task_management;

CREATE SCHEMA IF NOT EXISTS communication;

-- Create Tables in Schema: account
CREATE TABLE IF NOT EXISTS account."users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    profile_image TEXT
);

-- Create Tables in Schema: task_management
CREATE TABLE IF NOT EXISTS task_management.statuses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    color CHAR(6) NOT NULL
);

CREATE TABLE IF NOT EXISTS task_management.priorities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    color CHAR(6) NOT NULL
);

CREATE TABLE IF NOT EXISTS task_management.projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task_management.cards (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    start_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    finished_at TIMESTAMP,
    estimated_finished_at TIMESTAMP,
    status_id INTEGER REFERENCES task_management.statuses(id) ON DELETE SET NULL,
    manager_id INTEGER REFERENCES account."users"(id) ON DELETE SET NULL,
    assigned_id INTEGER REFERENCES account."users"(id) ON DELETE SET NULL,
    priority_id INTEGER REFERENCES task_management.priorities(id) ON DELETE SET NULL,
    parent_card_id INTEGER REFERENCES task_management.cards(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS task_management.boards (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    manager_id INTEGER REFERENCES account."users"(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    project_id INTEGER REFERENCES task_management.projects(id) ON DELETE SET NULL
);

-- Create Tables in Schema: communication
CREATE TABLE IF NOT EXISTS communication.messages (
    id SERIAL PRIMARY KEY,
    card_id INTEGER NOT NULL REFERENCES task_management.cards(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER REFERENCES account."users"(id) ON DELETE SET NULL,
    parent_message_id INTEGER REFERENCES communication.messages(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS communication.history (
    id SERIAL PRIMARY KEY,
    card_id INTEGER REFERENCES task_management.cards(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES account."users"(id) ON DELETE SET NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_before VARCHAR(50),
    status_after VARCHAR(50)
);
