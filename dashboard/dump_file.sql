--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Debian 16.0-1.pgdg120+1)
-- Dumped by pg_dump version 16.0 (Debian 16.0-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: lab_rank; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA lab_rank;


ALTER SCHEMA lab_rank OWNER TO postgres;

--
-- Name: access_level_mode_enum; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.access_level_mode_enum AS ENUM (
    'ADMIN',
    'TEACHER',
    'STUDENT'
);


ALTER TYPE lab_rank.access_level_mode_enum OWNER TO postgres;

--
-- Name: difficulty_enum; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.difficulty_enum AS ENUM (
    'EASY',
    'MEDIUM',
    'HARD'
);


ALTER TYPE lab_rank.difficulty_enum OWNER TO postgres;

--
-- Name: programming_language_enum; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.programming_language_enum AS ENUM (
    'C',
    'C++',
    'Java',
    'Python',
    'JavaScript',
    'Go',
    'Rust'
);


ALTER TYPE lab_rank.programming_language_enum OWNER TO postgres;

--
-- Name: status; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.status AS ENUM (
    'Accepted',
    'Memory Limit Exceeded',
    'Time Limit Exceeded',
    'Output Limit Exceeded',
    'File Error',
    'Nonzero Exit Status',
    'Signalled',
    'Internal Error',
    'Queued',
    'Running'
);


ALTER TYPE lab_rank.status OWNER TO postgres;

--
-- Name: syllabus_level_enum; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.syllabus_level_enum AS ENUM (
    'UNIVERSITY',
    'COLLEGE',
    'GLOBAL'
);


ALTER TYPE lab_rank.syllabus_level_enum OWNER TO postgres;

--
-- Name: user_status_enum; Type: TYPE; Schema: lab_rank; Owner: postgres
--

CREATE TYPE lab_rank.user_status_enum AS ENUM (
    'ACTIVE',
    'INACTIVE',
    'DELETED',
    'SPAM'
);


ALTER TYPE lab_rank.user_status_enum OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: access_level; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.access_level (
    id uuid NOT NULL,
    mode lab_rank.access_level_mode_enum NOT NULL,
    syllabus_id uuid NOT NULL
);


ALTER TABLE lab_rank.access_level OWNER TO postgres;

--
-- Name: auth; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.auth (
    user_id uuid NOT NULL,
    access_ids jsonb NOT NULL,
    salt character(32) NOT NULL,
    password_hash character(64) NOT NULL
);


ALTER TABLE lab_rank.auth OWNER TO postgres;

--
-- Name: college; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.college (
    id uuid NOT NULL,
    title character varying(50) NOT NULL,
    university_id uuid NOT NULL,
    description jsonb NOT NULL
);


ALTER TABLE lab_rank.college OWNER TO postgres;

--
-- Name: environment; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.environment (
    id uuid NOT NULL,
    title character varying(50) NOT NULL,
    link character varying(100) NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    update_events jsonb NOT NULL,
    live_dockerc_ids jsonb NOT NULL
);


ALTER TABLE lab_rank.environment OWNER TO postgres;

--
-- Name: problems; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.problems (
    id uuid NOT NULL,
    title character varying(100) NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    environment jsonb NOT NULL,
    problem_link character varying(100) NOT NULL,
    difficulty lab_rank.difficulty_enum NOT NULL,
    syllabus_id uuid NOT NULL,
    test_links character varying(100) NOT NULL
);


ALTER TABLE lab_rank.problems OWNER TO postgres;

--
-- Name: subject; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.subject (
    id uuid NOT NULL,
    title character varying(50) NOT NULL,
    description jsonb NOT NULL,
    university_id uuid NOT NULL
);


ALTER TABLE lab_rank.subject OWNER TO postgres;

--
-- Name: submissions; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.submissions (
    id uuid NOT NULL,
    problem_id uuid NOT NULL,
    link character varying(100) NOT NULL,
    created_by uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    score double precision,
    run_time character varying(10),
    metadata jsonb NOT NULL,
    lang lab_rank.programming_language_enum NOT NULL,
    status lab_rank.status DEFAULT 'Queued'::lab_rank.status NOT NULL,
    CONSTRAINT submissions_score_check CHECK (((score >= (0)::double precision) AND (score <= (100)::double precision)))
);


ALTER TABLE lab_rank.submissions OWNER TO postgres;

--
-- Name: syllabus; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.syllabus (
    id uuid NOT NULL,
    subject_id uuid NOT NULL,
    uni_college_id uuid NOT NULL,
    syllabus_level lab_rank.syllabus_level_enum NOT NULL
);


ALTER TABLE lab_rank.syllabus OWNER TO postgres;

--
-- Name: university; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank.university (
    id uuid NOT NULL,
    title character varying(50) NOT NULL,
    description jsonb NOT NULL
);


ALTER TABLE lab_rank.university OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: lab_rank; Owner: postgres
--

CREATE TABLE lab_rank."user" (
    id uuid NOT NULL,
    college_id uuid NOT NULL,
    status lab_rank.user_status_enum NOT NULL,
    created_at timestamp without time zone NOT NULL,
    email character varying(100) NOT NULL,
    contact_no character varying(15) NOT NULL,
    university_id uuid NOT NULL,
    dob date NOT NULL
);


ALTER TABLE lab_rank."user" OWNER TO postgres;

--
-- Data for Name: access_level; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.access_level (id, mode, syllabus_id) FROM stdin;
\.


--
-- Data for Name: auth; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.auth (user_id, access_ids, salt, password_hash) FROM stdin;
\.


--
-- Data for Name: college; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.college (id, title, university_id, description) FROM stdin;
6d4ac218-e7ed-48c7-b1c7-c81125ad7d96    GL BAJAJ        87700875-5784-4135-9276-3d3c294a17c3    "Example description"
\.


--
-- Data for Name: environment; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.environment (id, title, link, created_by, created_at, update_events, live_dockerc_ids) FROM stdin;
8b049cae-8b39-4318-a2da-ffc5eb650c6b    Python  some:docker-image       d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 20:50:19.331687      []      []
\.


--
-- Data for Name: problems; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.problems (id, title, created_by, created_at, environment, problem_link, difficulty, syllabus_id, test_links) FROM stdin;
8230dc12-f7c2-41f0-b9a1-20a4bca44bff    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:06:56.328095      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/830266cc-7fa2-42d7-af38-f5b2404705d6     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/136d7ee5-b40f-4cf7-90a3-e1a092179adf"}]
db30eef6-7c55-470a-bc49-bb4ae7150061    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:09:06.195503      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/49e4cfb7-1915-4a98-89e0-88ae80b17f70     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/09d73772-8798-4e6c-981c-bbc11fa52549"}]
e10f7035-3712-4a43-9b29-897f4c96928a    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:10:14.230404      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/ce88c0f4-5cdc-4e29-b8d6-b2dcd196d0d8     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/8385db3b-d469-4eb9-89d2-263eef491bf6"}]
4435ff57-b9ab-415f-bed3-028a14e6793d    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:16:48.534627      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/36583708-11c5-489e-aac4-9c84caf68caf     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/07a05c8f-6ac2-4cdf-b8c5-0a3d471ea6e8"}]
223f4779-4d95-4da5-9b45-82a0c91c12dd    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:18:51.270533      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/dfd8f6f0-03ae-4673-a876-8ccb7af23634     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/f860449b-ac90-4963-a356-c656a0b5514b"}]
74a88a72-919c-4808-befe-6e372ac53656    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:24:43.06316       [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/a3d3ef35-6cca-4cb7-b52c-28d17c31a094     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/d5bea898-a94d-4807-909e-da70ffde79bc"}]
ff325e82-338e-4654-a9fd-c42dddc01bb3    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:25:48.689784      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/ff325e82-338e-4654-a9fd-c42dddc01bb3     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/ff325e82-338e-4654-a9fd-c42dddc01bb3"}]
367e4439-f4dc-416c-82dc-efa8a01f7756    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:29:50.32519       [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/367e4439-f4dc-416c-82dc-efa8a01f7756     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/367e4439-f4dc-416c-82dc-efa8a01f7756"}]
e64264ca-2044-44a8-a890-a3a753c0b31f    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:30:06.072193      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/e64264ca-2044-44a8-a890-a3a753c0b31f     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"YourTestLanguage","link":"/files/testfile/e64264ca-2044-44a8-a890-a3a753c0b31f"}]
36731bb1-1169-44b2-b9e8-4a462cae4004    Fibo    d06b0c10-7f35-4351-b87f-c3e4b731b0d2    2023-11-14 23:32:39.557555      [{"id": "8b049cae-8b39-4318-a2da-ffc5eb650c6b", "language": "Python"}] /files/problem/36731bb1-1169-44b2-b9e8-4a462cae4004     MEDIUM  bf598b70-58bb-4c5f-8cde-c88f3f1735f0    [{"language":"Python","link":"/files/testfile/36731bb1-1169-44b2-b9e8-4a462cae4004","title":""}]
\.


--
-- Data for Name: subject; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.subject (id, title, description, university_id) FROM stdin;
367f355f-1e60-4707-b6af-f4ded6b59283    Mathematics     "Study of numbers, quantity, structure, space, and change."     87700875-5784-4135-9276-3d3c294a17c3
\.


--
-- Data for Name: submissions; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.submissions (id, problem_id, link, created_by, created_at, score, run_time, metadata, lang, status) FROM stdin;
\.


--
-- Data for Name: syllabus; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.syllabus (id, subject_id, uni_college_id, syllabus_level) FROM stdin;
bf598b70-58bb-4c5f-8cde-c88f3f1735f0    367f355f-1e60-4707-b6af-f4ded6b59283    6d4ac218-e7ed-48c7-b1c7-c81125ad7d96    COLLEGE
\.


--
-- Data for Name: university; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank.university (id, title, description) FROM stdin;
87700875-5784-4135-9276-3d3c294a17c3    AKTU    {"dummy": "value"}
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: lab_rank; Owner: postgres
--

COPY lab_rank."user" (id, college_id, status, created_at, email, contact_no, university_id, dob) FROM stdin;
d06b0c10-7f35-4351-b87f-c3e4b731b0d2    6d4ac218-e7ed-48c7-b1c7-c81125ad7d96    INACTIVE        2023-11-14 20:49:03.884782      user@example.com        1234567890      87700875-5784-4135-9276-3d3c294a17c3   2006-01-02
\.


--
-- Name: access_level access_level_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.access_level
    ADD CONSTRAINT access_level_pkey PRIMARY KEY (id);


--
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (user_id);


--
-- Name: college college_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.college
    ADD CONSTRAINT college_pkey PRIMARY KEY (id);


--
-- Name: environment environment_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.environment
    ADD CONSTRAINT environment_pkey PRIMARY KEY (id);


--
-- Name: problems problems_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.problems
    ADD CONSTRAINT problems_pkey PRIMARY KEY (id);


--
-- Name: subject subject_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.subject
    ADD CONSTRAINT subject_pkey PRIMARY KEY (id);


--
-- Name: submissions submissions_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.submissions
    ADD CONSTRAINT submissions_pkey PRIMARY KEY (id);


--
-- Name: syllabus syllabus_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.syllabus
    ADD CONSTRAINT syllabus_pkey PRIMARY KEY (id);


--
-- Name: university university_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.university
    ADD CONSTRAINT university_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: access_level access_level_syllabus_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.access_level
    ADD CONSTRAINT access_level_syllabus_id_fkey FOREIGN KEY (syllabus_id) REFERENCES lab_rank.syllabus(id);


--
-- Name: auth auth_user_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.auth
    ADD CONSTRAINT auth_user_id_fkey FOREIGN KEY (user_id) REFERENCES lab_rank."user"(id);


--
-- Name: college college_university_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.college
    ADD CONSTRAINT college_university_id_fkey FOREIGN KEY (university_id) REFERENCES lab_rank.university(id);


--
-- Name: environment environment_created_by_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.environment
    ADD CONSTRAINT environment_created_by_fkey FOREIGN KEY (created_by) REFERENCES lab_rank."user"(id);


--
-- Name: problems problems_created_by_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.problems
    ADD CONSTRAINT problems_created_by_fkey FOREIGN KEY (created_by) REFERENCES lab_rank."user"(id);


--
-- Name: problems problems_syllabus_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.problems
    ADD CONSTRAINT problems_syllabus_id_fkey FOREIGN KEY (syllabus_id) REFERENCES lab_rank.syllabus(id);


--
-- Name: subject subject_university_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.subject
    ADD CONSTRAINT subject_university_id_fkey FOREIGN KEY (university_id) REFERENCES lab_rank.university(id);


--
-- Name: submissions submissions_created_by_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.submissions
    ADD CONSTRAINT submissions_created_by_fkey FOREIGN KEY (created_by) REFERENCES lab_rank."user"(id);


--
-- Name: submissions submissions_problem_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.submissions
    ADD CONSTRAINT submissions_problem_id_fkey FOREIGN KEY (problem_id) REFERENCES lab_rank.problems(id);


--
-- Name: syllabus syllabus_subject_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.syllabus
    ADD CONSTRAINT syllabus_subject_id_fkey FOREIGN KEY (subject_id) REFERENCES lab_rank.subject(id);


--
-- Name: syllabus syllabus_uni_college_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank.syllabus
    ADD CONSTRAINT syllabus_uni_college_id_fkey FOREIGN KEY (uni_college_id) REFERENCES lab_rank.college(id);


--
-- Name: user user_college_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank."user"
    ADD CONSTRAINT user_college_id_fkey FOREIGN KEY (college_id) REFERENCES lab_rank.college(id);


--
-- Name: user user_university_id_fkey; Type: FK CONSTRAINT; Schema: lab_rank; Owner: postgres
--

ALTER TABLE ONLY lab_rank."user"
    ADD CONSTRAINT user_university_id_fkey FOREIGN KEY (university_id) REFERENCES lab_rank.university(id);


--
-- PostgreSQL database dump complete
--