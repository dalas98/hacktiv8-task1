CREATE TABLE todos (
    id int NOT NULL AUTO_INCREMENT,
    task varchar(255) NOT NULL,
    status varchar(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);