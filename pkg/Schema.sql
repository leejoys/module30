DROP TABLE IF EXISTS tasks_labels, tasks, users, labels;

CREATE TABLE users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL
);

CREATE TABLE labels (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL
);

CREATE TABLE tasks (
id SERIAL PRIMARY KEY,
opened BIGINT NOT NULL,
closed BIGINT DEFAULT 0,
author_id INTEGER REFERENCES users(id),
assigned_id INTEGER REFERENCES users(id),
title TEXT,
content TEXT
);

CREATE TABLE tasks_labels (
task_id INTEGER REFERENCES tasks(id),
label_id INTEGER REFERENCES labels(id)
);
