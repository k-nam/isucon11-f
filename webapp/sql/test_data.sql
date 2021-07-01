INSERT INTO `users` (`id`, `name`, `mail_address`, `hashed_password`, `type`) VALUES
(
  '01234567-89ab-cdef-0001-000000000001',
  '佐藤太郎',
  'sato@example.com',
  '$2a$10$KEgha.chGu1/N4kHZ./rIeK1QISkv8sYk15Mqktr6BGB8xomRRe02', -- "password"
  'student'
),
(
  '01234567-89ab-cdef-0001-000000000002',
  '鈴木次郎',
  'suzuki@example.com',
  '$2a$10$KEgha.chGu1/N4kHZ./rIeK1QISkv8sYk15Mqktr6BGB8xomRRe02',
  'student'
),
(
  '01234567-89ab-cdef-0001-000000000003',
  '高橋三郎',
  'takahashi@example.com',
  '$2a$10$KEgha.chGu1/N4kHZ./rIeK1QISkv8sYk15Mqktr6BGB8xomRRe02',
  'student'
),
(
  '01234567-89ab-cdef-0001-000000000004',
  '椅子昆',
  'isu@example.com',
  '$2a$10$KEgha.chGu1/N4kHZ./rIeK1QISkv8sYk15Mqktr6BGB8xomRRe02',
  'faculty'
);

INSERT INTO `courses` (`id`, `name`, `description`, `credit`, `classroom`, `capacity`) VALUES
(
  '01234567-89ab-cdef-0002-000000000001',
  '微分積分基礎',
  '微積分の基礎を学びます。',
  2,
  'A101講義室',
  100
),
(
  '01234567-89ab-cdef-0002-000000000002',
  '線形代数基礎',
  '線形代数の基礎を学びます。',
  2,
  'B101講義室',
  100
),
(
  '01234567-89ab-cdef-0002-000000000003',
  'アルゴリズム基礎',
  'アルゴリズムの基礎を学びます。',
  2,
  'C101講義室',
  150
),
(
  '01234567-89ab-cdef-0002-000000000011',
  '微分積分応用',
  '微積分の応用を学びます。',
  2,
  'A102講義室',
  100
),
(
  '01234567-89ab-cdef-0002-000000000012',
  '線形代数応用',
  '線形代数の応用を学びます。',
  2,
  'B102講義室',
  100
),
(
  '01234567-89ab-cdef-0002-000000000013',
  'プログラミング',
  'プログラミングを学びます。',
  2,
  'C102講義室',
  100
),
(
  '01234567-89ab-cdef-0002-000000000014',
  'プログラミング演習A',
  'プログラミングの演習を行います。',
  1,
  '演習室1',
  50
),
(
  '01234567-89ab-cdef-0002-000000000015',
  'プログラミング演習B',
  'プログラミングの演習を行います。',
  1,
  '演習室1',
  50
);

INSERT INTO `course_requirements` (`course_id`, `required_course_id`) VALUES
(
  '01234567-89ab-cdef-0002-000000000011',
  '01234567-89ab-cdef-0002-000000000001'
),
(
  '01234567-89ab-cdef-0002-000000000012',
  '01234567-89ab-cdef-0002-000000000002'
);

INSERT INTO `schedules` (`id`, `period`, `day_of_week`, `semester`, `year`) VALUES
(
  '01234567-89ab-cdef-0003-000000000001',
  1,
  'monday',
  'first',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000002',
  2,
  'wednesday',
  'first',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000003',
  3,
  'friday',
  'first',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000011',
  1,
  'tuesday',
  'second',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000012',
  2,
  'thursday',
  'second',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000013',
  3,
  'friday',
  'second',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000014',
  4,
  'friday',
  'second',
  2021
),
(
  '01234567-89ab-cdef-0003-000000000015',
  4,
  'friday',
  'second',
  2021
);

INSERT INTO `course_schedules` (`course_id`, `schedule_id`) VALUES
(
  '01234567-89ab-cdef-0002-000000000001',
  '01234567-89ab-cdef-0003-000000000001'
),
(
  '01234567-89ab-cdef-0002-000000000002',
  '01234567-89ab-cdef-0003-000000000002'
),
(
  '01234567-89ab-cdef-0002-000000000003',
  '01234567-89ab-cdef-0003-000000000003'
),
(
  '01234567-89ab-cdef-0002-000000000011',
  '01234567-89ab-cdef-0003-000000000011'
),
(
  '01234567-89ab-cdef-0002-000000000012',
  '01234567-89ab-cdef-0003-000000000012'
),
(
  '01234567-89ab-cdef-0002-000000000013',
  '01234567-89ab-cdef-0003-000000000013'
),
(
  '01234567-89ab-cdef-0002-000000000014',
  '01234567-89ab-cdef-0003-000000000014'
),
(
  '01234567-89ab-cdef-0002-000000000015',
  '01234567-89ab-cdef-0003-000000000015'
);

INSERT INTO `classes` (id, course_id, title, description, attendance_code) VALUES
(
  '01234567-89ab-cdef-0004-000000000001',
  '01234567-89ab-cdef-0002-000000000001',
  '微分積分基礎第一回',
  '微分積分の導入',
  'test_code'
),
(
  '01234567-89ab-cdef-0004-000000000002',
  '01234567-89ab-cdef-0002-000000000001',
  '微分積分基礎第二回',
  '微分(1)',
  'test_code2'
),
(
  '01234567-89ab-cdef-0004-000000000003',
  '01234567-89ab-cdef-0002-000000000002',
  '線形代数基礎第一回',
  '線形代数とは',
  'test_code3'
);


INSERT INTO `assignments` (id, class_id, name, description, deadline, created_at) VALUES
(
  '01234567-89ab-cdef-0005-000000000001',
  '01234567-89ab-cdef-0004-000000000001',
  '微分積分基礎 課題 1',
  '第一回課題. 明日までに提出',
  DATE_ADD(NOW(6), INTERVAL 1 DAY),
  NOW()
),
(
  '01234567-89ab-cdef-0005-000000000002',
  '01234567-89ab-cdef-0004-000000000003',
  '線形代数基礎 課題 1',
  '第一回課題. 1週間後までに提出',
  DATE_ADD(NOW(6), INTERVAL 7 DAY),
  NOW()
);
