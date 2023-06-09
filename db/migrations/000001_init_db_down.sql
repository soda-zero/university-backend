-- Drop statements for tables with dependencies
DROP TABLE IF EXISTS course_enrollemnt;
DROP TABLE IF EXISTS assistant_professor;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS career_enrollment_evaluation;
DROP TABLE IF EXISTS career_enrollment;
DROP TABLE IF EXISTS course_dependency;
DROP TABLE IF EXISTS professor_course;
DROP TABLE IF EXISTS course_ocurrence;

-- Drop statements for remaining tables
DROP TABLE IF EXISTS career_enrollment;
DROP TABLE IF EXISTS career_status;
DROP TABLE IF EXISTS student;
DROP TABLE IF EXISTS course;
DROP TABLE IF EXISTS career;
DROP TABLE IF EXISTS professor;
DROP TABLE IF EXISTS evaluation_type;
DROP TABLE IF EXISTS room;
DROP TABLE IF EXISTS career_level;
DROP TABLE IF EXISTS department;

-- Drop statement for the extension
DROP EXTENSION IF EXISTS "uuid-ossp";

