-- Set the search path to the "lab-rank" schema
SET search_path TO lab_rank;

-- Create the "lab-rank" schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS lab_rank;

-- Create the "access_level" table with the "title" field and a foreign key reference to syllabus
CREATE TYPE lab_rank.access_level_mode_enum AS ENUM (
    'ADMIN',
    'TEACHER',
    'STUDENT'
);

CREATE TYPE lab_rank.programming_language_enum AS ENUM (
    'C',
    'C++',
    'Java',
    'Python',
    'JavaScript',
    'Go',
    'Rust'
    -- Add more programming languages as needed
);

CREATE TYPE lab_rank.difficulty_enum AS ENUM (
    'EASY',
    'MEDIUM',
    'HARD'
);

CREATE TYPE lab_rank.syllabus_level_enum AS ENUM (
    'UNIVERSITY',
    'COLLEGE',
    'GLOBAL'
);

CREATE TYPE lab_rank.user_status_enum AS ENUM (
    'ACTIVE',
    'INACTIVE',
    'DELETED',
    'SPAM'
);

CREATE TABLE lab_rank.access_level (
    id UUID PRIMARY KEY,
    mode lab_rank.access_level_mode_enum,
    syllabus_id UUID REFERENCES lab_rank.syllabus(id)
);

-- Define the "lab-rank" schema for the rest of the tables
SET search_path TO lab_rank;

-- Define the Submissions table
CREATE TABLE lab_rank.submissions (
    id UUID PRIMARY KEY,
    problem_id UUID REFERENCES lab_rank.problems(id),
    link VARCHAR(100),
    created_by UUID REFERENCES lab_rank."user"(id),
    created_at TIMESTAMP,
    score FLOAT(5, 2),
    run_time VARCHAR(10),
    metadata JSONB,
    lang lab_rank.programming_language_enum,
    CHECK (score >= 0 AND score <= 100)
);

-- Define the Environment table
CREATE TABLE lab_rank.environment (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    link VARCHAR(100),
    created_by UUID REFERENCES lab_rank."user"(id),
    created_at TIMESTAMP,
    update_events JSONB,
    live_dockerc_ids JSONB
);

-- Define the Problems table
CREATE TABLE lab_rank.problems (
    id UUID PRIMARY KEY,
    title VARCHAR(100),
    created_by UUID REFERENCES lab_rank."user"(id),
    created_at TIMESTAMP,
    environment JSONB,
    problem_link VARCHAR(100),
    difficulty lab_rank.difficulty_enum,
    syllabus_id UUID REFERENCES lab_rank.syllabus(id),
    test_link VARCHAR(100)
);

-- Define the Syllabus table
CREATE TABLE lab_rank.syllabus (
    id UUID PRIMARY KEY,
    subject_id UUID REFERENCES lab_rank.subject(id),
    uni_college_id UUID REFERENCES lab_rank.college(id),
    syllabus_level lab_rank.syllabus_level_enum
);

-- Define the University table
CREATE TABLE lab_rank.university (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    description JSONB
);

-- Define the Subject table
CREATE TABLE lab_rank.subject (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    description JSONB,
    university UUID REFERENCES lab_rank.university(id)
);

-- Define the College table
CREATE TABLE lab_rank.college (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    university_id UUID REFERENCES lab_rank.university(id),
    description JSONB
);

-- Define the "user" table with basic user information
CREATE TABLE lab_rank."user" (
    id UUID PRIMARY KEY,
    college_id UUID REFERENCES lab_rank.college(id),
    status lab_rank.user_status_enum,
    created_at TIMESTAMP,
    email VARCHAR(100),
    contact_no VARCHAR(15),
    university_id UUID REFERENCES lab_rank.university(id),
    dob DATE
);

-- Define the "auth" table to store authentication information
CREATE TABLE lab_rank.auth (
    user_id UUID PRIMARY KEY REFERENCES lab_rank."user"(id),
    access_ids JSONB,
    salt CHAR(32),
    password_hash CHAR(64)
);
