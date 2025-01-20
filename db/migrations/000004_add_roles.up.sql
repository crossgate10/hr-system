-- 插入角色
INSERT INTO roles (id, name) VALUES (1, 'Member');
INSERT INTO roles (id, name) VALUES (2, 'Team Manager');
INSERT INTO roles (id, name) VALUES (3, 'Boss');

-- 插入部門
INSERT INTO departments (id, name) VALUES (1, 'Backend');
INSERT INTO departments (id, name) VALUES (2, 'Frontend');
INSERT INTO departments (id, name) VALUES (3, 'HR');

-- 插入審核人員
INSERT INTO approvers (id, name, seq) VALUES (1, 'Backend Manager', 3);
INSERT INTO approvers (id, name, seq) VALUES (2, 'Frontend Manager', 3);
INSERT INTO approvers (id, name, seq) VALUES (3, 'HR Manager', 3);
INSERT INTO approvers (id, name, seq) VALUES (4, 'VP of Engineering', 4);
INSERT INTO approvers (id, name, seq) VALUES (5, 'Boss', 999);
INSERT INTO approvers (id, name, seq) VALUES (6, 'HR', 1);
INSERT INTO approvers (id, name, seq) VALUES (7, 'Agent', 2);

-- 插入角色與部門的關聯
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 1, 6); -- Backend Member -> HR
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 1, 7); -- Backend Member -> Agent
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 1, 1); -- Backend Member -> Backend Manager
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 2, 6); -- Frontend Member -> HR
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 2, 7); -- Frontend Member -> Agent
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 2, 2); -- Frontend Member -> Frontend Manager
INSERT INTO role_department (role_id, department_id, approver_id) VALUES (1, 3, 3); -- HR Member -> HR Manager