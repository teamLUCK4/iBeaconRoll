-- users
INSERT INTO users (name) VALUES
('Hailey Kim'),   -- student_id = 1
('Soo Kim');      -- student_id = 2


-- timetables
-- Hailey Kim's classes (student_id = 1)
INSERT INTO timetables (student_id, semester, subject_name, day_of_week, start_time, end_time, classroom) VALUES
(1, 1, 'Algorithms', 'Mon', '09:00:00', '10:30:00', 'Building 302'),   -- id = 1
(1, 1, 'Data Structures', 'Wed', '13:00:00', '14:30:00', 'Building 304'); -- id = 2

-- Soo Kim's class (student_id = 2)
INSERT INTO timetables (student_id, semester, subject_name, day_of_week, start_time, end_time, classroom) VALUES
(2, 1, 'Artificial Intelligence', 'Tue', '10:00:00', '11:30:00', 'Building 303'); -- id = 3


-- attendances
-- Hailey Kim's attendance for "Algorithms" (timetable_id = 1)
INSERT INTO attendances (student_id, timetable_id, attendance_date, attendance_time, classroom, status) VALUES
(1, 1, '2025-05-06', '09:00:00', 'Building 302', 'Present'),
(1, 1, '2025-05-13', '09:00:00', 'Building 302', 'Late');

-- Soo Kim's attendance for "Artificial Intelligence" (timetable_id = 3)
INSERT INTO attendances (student_id, timetable_id, attendance_date, attendance_time, classroom, status) VALUES
(2, 3, '2025-05-06', '10:00:00', 'Building 303', 'Absent');
