-- Set the search path to the "lab-rank" schema
SET search_path TO lab_rank;

-- Create the "lab-rank" schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS lab_rank;

-- Define the "lab-rank" schema for the rest of the tables
SET search_path TO lab_rank;

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

-- Create the "access_level" table with all fields NOT NULL
CREATE TABLE lab_rank.access_level (
    id UUID PRIMARY KEY NOT NULL,
    mode lab_rank.access_level_mode_enum NOT NULL,
    syllabus_id UUID REFERENCES lab_rank.syllabus(id) NOT NULL
);


-- Define the Environment table with all fields NOT NULL
CREATE TABLE lab_rank.environment (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    link VARCHAR(100) NOT NULL,
    created_by UUID REFERENCES lab_rank."user"(id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    update_events JSONB NOT NULL,
    live_dockerc_ids JSONB NOT NULL
);

-- Define the Problems table with all fields NOT NULL
CREATE TABLE lab_rank.problems (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(100) NOT NULL,
    created_by UUID REFERENCES lab_rank."user"(id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    environment JSONB NOT NULL,
    problem_link VARCHAR(100) NOT NULL,
    difficulty lab_rank.difficulty_enum NOT NULL,
    syllabus_id UUID REFERENCES lab_rank.syllabus(id) NOT NULL,
    test_links JSONB NOT NULL
);

-- Define the Syllabus table with all fields NOT NULL
CREATE TABLE lab_rank.syllabus (
    id UUID PRIMARY KEY NOT NULL,
    subject_id UUID REFERENCES lab_rank.subject(id) NOT NULL,
    uni_college_id UUID REFERENCES lab_rank.college(id) NOT NULL,
    syllabus_level lab_rank.syllabus_level_enum NOT NULL
);

-- Define the University table with all fields NOT NULL
CREATE TABLE lab_rank.university (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    description JSONB NOT NULL
);

-- Define the Subject table with all fields NOT NULL
CREATE TABLE lab_rank.subject (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    description JSONB NOT NULL,
    university_id UUID REFERENCES lab_rank.university(id) NOT NULL
);

-- Define the College table with all fields NOT NULL
CREATE TABLE lab_rank.college (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    university_id UUID REFERENCES lab_rank.university(id) NOT NULL,
    description JSONB NOT NULL
);

-- Define the "user" table with all fields NOT NULL
CREATE TABLE lab_rank.user (
    id UUID PRIMARY KEY NOT NULL,
    college_id UUID REFERENCES lab_rank.college(id) NOT NULL,
    status lab_rank.user_status_enum NOT NULL,
    created_at TIMESTAMP NOT NULL,
    email VARCHAR(100) NOT NULL,
    contact_no VARCHAR(15) NOT NULL,
    university_id UUID REFERENCES lab_rank.university(id) NOT NULL,
    dob DATE NOT NULL
);

-- Define the "auth" table to store authentication information
CREATE TABLE lab_rank.auth (
    user_id UUID PRIMARY KEY REFERENCES lab_rank."user"(id) NOT NULL,
    access_ids JSONB NOT NULL,
    salt CHAR(32) NOT NULL,
    password_hash CHAR(64) NOT NULL
);

-- Define the Submissions table with all fields NOT NULL
CREATE TABLE lab_rank.submissions (
    id UUID PRIMARY KEY NOT NULL,
    problem_id UUID REFERENCES lab_rank.problems(id) NOT NULL,
    link VARCHAR(100) NOT NULL,
    created_by UUID REFERENCES lab_rank."user"(id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    score FLOAT,
    run_time VARCHAR(10),
    metadata JSONB NOT NULL,
    lang lab_rank.programming_language_enum NOT NULL,
    CHECK (score >= 0 AND score <= 100)
);
