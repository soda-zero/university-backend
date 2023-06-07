CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE career_status(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(100)
);
