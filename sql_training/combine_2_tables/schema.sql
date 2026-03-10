-- LeetCode 175: Combine Two Tables
-- PostgreSQL schema creation

CREATE TABLE Person (
    personId SERIAL PRIMARY KEY,
    lastName VARCHAR(255) NOT NULL,
    firstName VARCHAR(255) NOT NULL
);

CREATE TABLE Address (
    addressId SERIAL PRIMARY KEY,
    personId INTEGER,
    city VARCHAR(255),
    state VARCHAR(255),
    FOREIGN KEY (personId) REFERENCES Person(personId)
);
