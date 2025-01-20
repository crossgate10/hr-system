-- 刪除角色與部門的關聯
DELETE FROM role_department WHERE role_id = 1 AND department_id = 1 AND approver_id = 6; -- Backend Member -> HR
DELETE FROM role_department WHERE role_id = 1 AND department_id = 1 AND approver_id = 7; -- Backend Member -> Agent
DELETE FROM role_department WHERE role_id = 1 AND department_id = 1 AND approver_id = 1; -- Backend Member -> Backend Manager
DELETE FROM role_department WHERE role_id = 1 AND department_id = 2 AND approver_id = 6; -- Frontend Member -> HR
DELETE FROM role_department WHERE role_id = 1 AND department_id = 2 AND approver_id = 7; -- Frontend Member -> Agent
DELETE FROM role_department WHERE role_id = 1 AND department_id = 2 AND approver_id = 2; -- Frontend Member -> Frontend Manager
DELETE FROM role_department WHERE role_id = 1 AND department_id = 3 AND approver_id = 3; -- HR Member -> HR Manager

-- 刪除審核人員
DELETE FROM approvers WHERE id = 1;
DELETE FROM approvers WHERE id = 2;
DELETE FROM approvers WHERE id = 3;
DELETE FROM approvers WHERE id = 4;
DELETE FROM approvers WHERE id = 5;
DELETE FROM approvers WHERE id = 6;
DELETE FROM approvers WHERE id = 7;

-- 刪除部門
DELETE FROM departments WHERE id = 1;
DELETE FROM departments WHERE id = 2;
DELETE FROM departments WHERE id = 3;

-- 刪除角色
DELETE FROM roles WHERE id = 1;
DELETE FROM roles WHERE id = 2;
DELETE FROM roles WHERE id = 3;
