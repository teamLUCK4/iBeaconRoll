-- users
CREATE TABLE users (
    student_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- timetables
CREATE TABLE timetables (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users(student_id) ON DELETE CASCADE,
    semester INTEGER NOT NULL,
    subject_name VARCHAR(50) NOT NULL,
    day_of_week VARCHAR(10) CHECK (day_of_week IN ('Mon', 'Tue', 'Wed', 'Thu', 'Fri')),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    classroom VARCHAR(20) NOT NULL
);

-- attendances
CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users(student_id) ON DELETE CASCADE,
    timetable_id INTEGER REFERENCES timetables(id) ON DELETE SET NULL,
    attendance_date DATE NOT NULL,
    attendance_time TIME NOT NULL,
    classroom VARCHAR(20),
    status VARCHAR(10) CHECK (status IN ('Present', 'Absent', 'Late')) NOT NULL
);
