-- Create the access_level table with the "title" field and a foreign key reference to syllabus
CREATE TYPE access_level_mode_enum AS ENUM (
    'ADMIN',
    'TEACHER',
    'STUDENT'
);
CREATE TYPE programming_language_enum AS ENUM (
    'C',
    'C++',
    'Java',
    'Python',
    'JavaScript',
    'Go',
    'Rust',
    -- Add more programming languages as needed
);
CREATE TYPE difficulty_enum AS ENUM (
    'EASY',
    'MEDIUM',
    'HARD'
);
CREATE TYPE syllabus_level_enum AS ENUM (
    'UNIVERSITY',
    'COLLEGE',
    'GLOBAL'
);


CREATE TABLE access_level (
    id UUID PRIMARY KEY,
    mode access_level_mode_enum,
    syllabus_id UUID REFERENCES syllabus(id)
);


-- Define the Submissions table
CREATE TABLE submissions (
    id UUID PRIMARY KEY,
    problem_id UUID REFERENCES problems(id),
    link VARCHAR(100), 
    created_by UUID REFERENCES user(id),
    created_at TIMESTAMP, 
    score FLOAT(5, 2),
    run_time VARCHAR(10), 
    metadata JSONB, 
    lang programming_language_enum,
    CHECK (score >= 0 AND score <= 100) 
);

-- Define the Environment table
CREATE TABLE environment (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    link VARCHAR(100),
    created_by UUID REFERENCES user(id),
    created_at TIMESTAMP, 
    update_events JSONB,
    live_dockerc_ids JSONB 
);

-- Define the Problems table
CREATE TABLE problems (
    id UUID PRIMARY KEY,
    title VARCHAR(100),
    created_by UUID REFERENCES user(id), 
    created_at TIMESTAMP, 
    -- update_events JSONB, 
    environment JSONB,
    problem_link VARCHAR(100), 
    difficulty difficulty_enum, 
    syllabus_id UUID REFERENCES syllabus(id), 
    test_link VARCHAR(100)
);

-- Define the Syllabus table
CREATE TABLE syllabus (
    id UUID PRIMARY KEY,
    subject_id UUID REFERENCES subject(id),
    uni_college_id UUID REFERENCES college(id),
    syllabus_level syllabus_level_enum 
);

-- Define the Problem Events table
-- CREATE TABLE problem_events (
--     id UUID PRIMARY KEY,
--     event_message VARCHAR, -- Replace with appropriate data type
--     created_by UUID REFERENCES user(id), -- This should be a reference to the user table
--     created_at ENUM, -- Replace with appropriate data type
--     problem_link VARCHAR, -- Replace with appropriate data type
--     test_link VARCHAR -- Replace with appropriate data type
-- );

-- Define the Environment Events table
-- CREATE TABLE environment_events (
--     id UUID PRIMARY KEY,
--     event_message VARCHAR, -- Replace with appropriate data type
--     created_by UUID REFERENCES user(id), -- This should be a reference to the user table
--     created_at ENUM, -- Replace with appropriate data type
--     link VARCHAR -- Replace with appropriate data type
-- );

-- Define the University table
CREATE TABLE university (
    id UUID PRIMARY KEY,
    title VARCHAR(50), 
    description JSONB
);

-- Define the Subject table
CREATE TABLE subject (
    id UUID PRIMARY KEY,
    title VARCHAR(50), 
    description JSONB,
    university UUID REFERENCES university(id) 
);

-- Define the College table
CREATE TABLE college (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    university_id UUID REFERENCES university(id), 
    description JSONB,
);
