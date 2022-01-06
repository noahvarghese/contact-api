DROP DATABASE IF EXISTS contact;
CREATE DATABASE contact;
USE contact;

CREATE TABLE hosts (
    id INT NOT NULL AUTO_INCREMENT,
    url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(url)
);

CREATE TABLE fields (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    host_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES host(id)
);

CREATE TABLE template (
    id INT NOT NULL AUTO_INCREMENT,
    template TEXT NOT NULL,
    host_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES host(id)
)