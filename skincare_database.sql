CREATE DATABASE beauty_store;

USE beauty_store;

CREATE TABLE member (
    ID_MEMBER INT AUTO_INCREMENT PRIMARY KEY,
    USERNAME VARCHAR(255) NOT NULL,
    GENDER ENUM('Male', 'Female', 'Other') NOT NULL,
    SKINTYPE VARCHAR(255) NOT NULL,
    SKINCOLOR VARCHAR(255) NOT NULL
);

CREATE TABLE product (
    ID_PRODUCT INT AUTO_INCREMENT PRIMARY KEY,
    NAME_PRODUCT VARCHAR(255) NOT NULL,
    PRICE DECIMAL(10, 2) NOT NULL
);

CREATE TABLE review_product (
    ID_REVIEW INT AUTO_INCREMENT PRIMARY KEY,
    ID_PRODUCT INT NOT NULL,
    ID_MEMBER INT NOT NULL,
    DESC_REVIEW TEXT NOT NULL,
    FOREIGN KEY (ID_PRODUCT) REFERENCES product(ID_PRODUCT),
    FOREIGN KEY (ID_MEMBER) REFERENCES member(ID_MEMBER)
);

CREATE TABLE like_review (
    ID_REVIEW INT NOT NULL,
    ID_MEMBER INT NOT NULL,
    PRIMARY KEY (ID_REVIEW, ID_MEMBER),
    FOREIGN KEY (ID_REVIEW) REFERENCES review_product(ID_REVIEW),
    FOREIGN KEY (ID_MEMBER) REFERENCES member(ID_MEMBER)
);

SELECT * FROM member;