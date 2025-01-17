CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) NOT NULL,
    description VARCHAR(100),
    due_date VARCHAR(40),
    completed VARCHAR(5),
    created_at TIMESTAMP DEFAULT NOW()
);