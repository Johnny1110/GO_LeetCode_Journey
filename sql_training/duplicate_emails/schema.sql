-- LeetCode 182: Duplicate Emails
-- PostgreSQL schema creation

CREATE TABLE Person (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL
);