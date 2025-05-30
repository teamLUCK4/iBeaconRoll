-- users
CREATE TABLE users (
    student_id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    user_pss VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL
);

-- timetables
CREATE TABLE timetables (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES users(student_id) ON DELETE CASCADE,
    semester INTEGER NOT NULL,
    subject_name VARCHAR(50) NOT NULL,
    day_of_week VARCHAR(10) CHECK (day_of_week IN ('Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun')),
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
    status VARCHAR(10) CHECK (status IN ('Waiting', 'Ongoing', 'Completed', 'Absent')) NOT NULL
);



-- ================================
-- Seed: Sample Data Insertions
-- ================================\


-- users
INSERT INTO users (name, user_id, user_pss) VALUES
('Hailey Kim', 'hailey', '1234'),   -- student_id = 1
('Soo Kim', 'soo', '1234');      -- student_id = 2


-- timetables
-- Hailey Kim's classes (student_id = 1)
INSERT INTO timetables (student_id, semester, subject_name, day_of_week, start_time, end_time, classroom) VALUES
(1, 1, 'Algorithms', 'Mon', '09:00:00', '10:30:00', 'Building 302'),   -- id = 1
(1, 1, 'Data Structures', 'Wed', '13:00:00', '14:30:00', 'Building 304'),  -- id = 2

-- -- 테스트용 : 금요일 시간표
(1, 8, 'Artificial Intelligence', 'Fri', '10:00:00', '11:30:00', 'Building 301'),
(1, 8, 'Algorithms', 'Fri', '12:30:00', '14:30:00', 'Building 302'),
(1, 8, 'Data Structures', 'Fri', '15:00:00', '16:30:00', 'Building 303'),
(1, 8, 'OS', 'Fri', '18:00:00', '20:30:00', 'Building 304');


-- -- 테스트용 : 일요일 시간표
(1, 8, 'Artificial Intelligence', 'Sun', '10:00:00', '11:30:00', 'Building 301'),
(1, 8, 'Algorithms', 'Sun', '12:30:00', '14:30:00', 'Building 302'),
(1, 8, 'Data Structures', 'Sun', '15:00:00', '16:30:00', 'Building 303'),
(1, 8, 'OS', 'Sun', '18:00:00', '20:30:00', 'Building 304');


-- attendances
-- Hailey Kim's attendance for "Algorithms" (timetable_id = 1)
INSERT INTO attendances (student_id, timetable_id, attendance_date, attendance_time, classroom, status) VALUES
(1, 1, '2025-05-06', '09:00:00', 'Building 302', 'Waiting'),
(1, 1, '2025-05-13', '09:00:00', 'Building 302', 'Waiting'),
(1, 4, '2025-05-06', '10:00:00', 'Building 303', 'Waiting');