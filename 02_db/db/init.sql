CREATE TABLE money (
    PRIMARY KEY (id),
    id   INT(8) NOT NULL AUTOINCREMENT,
    data INT(8) NOT NULL
);

INSERT INTO money (data)
VALUES (10000),
       (20000),
       (30000);
