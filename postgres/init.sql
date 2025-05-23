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



-- ================================
-- Seed: Sample Data Insertions
-- ================================\


-- users
INSERT INTO users (name) VALUES
('Hailey Kim'),   -- student_id = 1
('Soo Kim');      -- student_id = 2


-- timetables
-- Hailey Kim's classes (student_id = 1)
INSERT INTO timetables (student_id, semester, subject_name, day_of_week, start_time, end_time, classroom) VALUES
(1, 1, 'Algorithms', 'Mon', '09:00:00', '10:30:00', 'Building 302'),   -- id = 1
(1, 1, 'Data Structures', 'Wed', '13:00:00', '14:30:00', 'Building 304'),  -- id = 2
(1, 1, 'OS', 'Fri', '13:00:00', '14:30:00', 'Building 302'),  -- id = 3
(2, 1, 'Artificial Intelligence', 'Tue', '10:00:00', '11:30:00', 'Building 303');  -- id = 4


-- attendances
-- Hailey Kim's attendance for "Algorithms" (timetable_id = 1)
INSERT INTO attendances (student_id, timetable_id, attendance_date, attendance_time, classroom, status) VALUES
(1, 1, '2025-05-06', '09:00:00', 'Building 302', 'Present'),
(1, 1, '2025-05-13', '09:00:00', 'Building 302', 'Late'),
(2, 4, '2025-05-06', '10:00:00', 'Building 303', 'Absent');