DROP DATABASE IF EXISTS contact;
CREATE DATABASE contact;
USE contact;

CREATE TABLE hosts (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    has_images TINYINT(1) NOT NULL,
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
    FOREIGN KEY (host_id) REFERENCES hosts(id)
);

CREATE TABLE messages (
    id INT NOT NULL AUTO_INCREMENT,
    original TEXT NOT NULL,
    msg TEXT NOT NULL,
    sent TINYINT(1) NOT NULL DEFAULT 0,
    host_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES hosts(id)
);

CREATE TABLE templates (
    id INT NOT NULL AUTO_INCREMENT,
    template TEXT NOT NULL,
    host_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (host_id) REFERENCES hosts(id)
);