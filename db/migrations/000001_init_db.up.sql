CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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
    department_id UUID NOT NULL REFERENCES department(id),
    career_level_id UUID NOT NULL REFERENCES career_level(id)
);

CREATE TABLE IF NOT EXISTS professor(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS evaluation_type(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    is_group_activity BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS room(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(100) NOT NULL,
    capacity INT NOT NULL,
    location VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS student(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS career_status(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    updated_at DATE
);

CREATE TABLE IF NOT EXISTS course(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    semester INT NOT NULL,
    optative BOOLEAN NOT NULL,
    course_code VARCHAR(20) NOT NULL,
    career_id UUID NOT NULL REFERENCES career(id)
);

CREATE TABLE IF NOT EXISTS professor_course(
    professor_id UUID REFERENCES professor(id),
    course_id UUID REFERENCES course(id),
    PRIMARY KEY (professor_id, course_id)
);
CREATE TABLE IF NOT EXISTS course_dependency(
    course_id UUID REFERENCES course(id),
    required_course_id UUID REFERENCES course(id),
    PRIMARY KEY (course_id, required_course_id)
);

CREATE TABLE IF NOT EXISTS course_ocurrence(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    ocurrence_year INT NOT NULL,
    code VARCHAR(20) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    capacity INT NOT NULL,
    course_id UUID NOT NULL REFERENCES course(id),
    professor_id UUID NOT NULL REFERENCES professor(id)
);

CREATE TABLE IF NOT EXISTS assistant_professor(
    professor_id UUID NOT NULL REFERENCES professor(id),
    course_ocurrence_id UUID NOT NULL REFERENCES course_ocurrence(id)
);

CREATE TABLE IF NOT EXISTS schedule(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    day_of_week CHAR(2) NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    room_id UUID NOT NULL REFERENCES room(id),
    course_ocurrence_id UUID NOT NULL REFERENCES course_ocurrence(id)
);

CREATE TABLE IF NOT EXISTS career_enrollment(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    student_id UUID NOT NULL REFERENCES student(id),
    career_id UUID NOT NULL REFERENCES career(id),
    enrollment_date  DATE NOT NULL,
    career_status_id UUID NOT NULL REFERENCES career_status(id)
);

CREATE TABLE IF NOT EXISTS course_enrollment(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    student_id UUID NOT NULL REFERENCES student(id),
    course_ocurrence_id UUID NOT NULL REFERENCES course_ocurrence(id),
    final_score DECIMAL(5,2)
);

CREATE TABLE IF NOT EXISTS career_enrollment_evaluation(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    course_enrollment_id UUID NOT NULL REFERENCES course_enrollment(id),
    evaluation_type_id UUID NOT NULL REFERENCES evaluation_type(id),
    evaluation_date DATE,
    final_score DECIMAL(5,2)
);
