CREATE TABLE DAG (
    id SERIAL PRIMARY KEY,
    dag_id VARCHAR(255) NOT NULL UNIQUE,
    path VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    schedule VARCHAR(255),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    concurrency INT,
    max_active_tasks INT,
    time_out INT,
    catchup BOOLEAN,
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);