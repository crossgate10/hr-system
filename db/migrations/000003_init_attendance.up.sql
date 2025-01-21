CREATE TABLE leave_requests
(
    id               INT PRIMARY KEY AUTO_INCREMENT,
    employee_id      INT          NOT NULL,
    leave_type       VARCHAR(255) NOT NULL,
    start_time       DATETIME     NOT NULL,
    end_time         DATETIME     NOT NULL,
    total_hours      FLOAT        NOT NULL,
    substitute_id    INT          NOT NULL,
    description      TEXT,
    approvers        TEXT,
    application_time DATETIME     NOT NULL
);
