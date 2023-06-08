CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Career Planning
CREATE TABLE IF NOT EXISTS department(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS career_level(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS career(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    duration_years INT NOT NULL,
    required_optative_coursers INT NOT NULL,
    department_id UUID REFERENCES department(id),
    career_level_id UUID REFERENCES career_level(id)
);

CREATE TABLE IF NOT EXISTS course(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    semester INT NOT NULL,
    optative BOOLEAN NOT NULL,
    course_code VARCHAR(20),
    career_id UUID REFERENCES career(id)
);

CREATE TABLE IF NOT EXISTS course_dependency(
    course_id UUID REFERENCES course(id),
    required_course_id UUID REFERENCES course(id),
    PRIMARY KEY (course_id, required_course_id)
);

CREATE TABLE IF NOT EXISTS professor(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS professor_course(
    professor_id UUID REFERENCES professor(id),
    course_id UUID REFERENCES course(id)
);

-- TODO
-- Studen & Enrollment
-- CREATE TABLE IF NOT EXISTS career_status(
--     id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
--     name VARCHAR(100),
-- );
--
-- CREATE TABLE IF NOT EXISTS student(
--     id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
--     name VARCHAR(50) NOT NULL,
--     last_name VARCHAR(50) NOT NULL,
--     birth_date DATE NOT NULL,
--     email VARCHAR(100),
--     phone_number VARCHAR(20),
-- );
--
-- CREATE TABLE IF NOT EXISTS career_enrollment(
--     id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
--     student_id UUID REFERENCES student(id),
-- );
