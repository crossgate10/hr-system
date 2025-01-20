CREATE TABLE roles
(
    id   INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE departments
(
    id   INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE approvers
(
    id   INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    seq  INT
);

CREATE TABLE role_department
(
    role_id       INT,
    department_id INT,
    approver_id   INT,
    PRIMARY KEY (role_id, department_id, approver_id),
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (department_id) REFERENCES departments (id),
    FOREIGN KEY (approver_id) REFERENCES approvers (id)
);