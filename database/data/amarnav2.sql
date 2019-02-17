DROP DATABASE amarna;
CREATE DATABASE amarna;
USE amarna;

CREATE TABLE Language
(
    lang_name VARCHAR(32),
    PRIMARY KEY(lang_name)
);

CREATE TABLE User
(
    username VARCHAR(64),
    knownLang VARCHAR(32),
    learnLang VARCHAR(32),
    PRIMARY KEY (username),
    FOREIGN KEY (knownLang) REFERENCES Language(lang_name) ON DELETE SET NULL,
    FOREIGN KEY (learnLang) REFERENCES Language(lang_name) ON DELETE SET NULL
);

CREATE TABLE Pairing
(
    leftUser VARCHAR(64),
    rightUser VARCHAR(64),
    leftCount INT DEFAULT 1,
    rightCount INT DEFAULT 1,
    PRIMARY KEY (leftUser, rightUser),
    FOREIGN KEY (leftUser) REFERENCES User(username) ON DELETE CASCADE,
    FOREIGN KEY (rightUser) REFERENCES User(username) ON DELETE CASCADE
);

CREATE TABLE Lesson
(
    referenceID BIGINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(128) NOT NULL,
    section VARCHAR(128) NOT NULL,
    description VARCHAR(512) NOT NULL,
    PRIMARY KEY(referenceID)
);

CREATE TABLE Letter
(
    leftUser VARCHAR(64), -- also denotes from
    rightUser VARCHAR(64),
    referenceID BIGINT,
    ts TIMESTAMP,
    body VARCHAR(1024) NOT NULL,
    PRIMARY KEY (leftUser, rightUser, referenceID, ts),
    FOREIGN KEY (leftUser, rightUser) REFERENCES Pairing(leftUser, rightUser)
        ON DELETE CASCADE,
    FOREIGN KEY (referenceID) REFERENCES Lesson(referenceID)
        ON DELETE CASCADE
);

SELECT "----------- ADDING DATA ----------------_" as "";

INSERT INTO Language
    (lang_name)
VALUES
    ("English"),
    ("Spanish"),
    ("Chinese");

INSERT INTO User
    (username, knownLang, learnLang)
VALUES
    ("amvasquez", "Spanish", "English"),
    ("gqo", "English", "Spanish"),
    ("rtr", "Chinese", "English"),
    ("justin39", "English", "Chinese");

INSERT INTO Pairing
    (leftUser, rightUser)
VALUES
    ("amvasquez", "gqo"),
    ("gqo", "amvasquez"),
    ("rtr", "justin39"),
    ("justin39", "rtr");

INSERT INTO Lesson
    (title, section, description)
VALUES
    ("INTRODUCTION","NAMES","Say your name!"), -- 1
    ("INTRODUCTION","LOCATION","Say where you're from!"), -- 2
    ("HOBBIES","MUSIC","Say if you like music!"); -- 3

INSERT INTO Letter
    (leftUser, rightUser, referenceID, ts, body)
VALUES
    ("gqo", "amvasquez", 1, NOW(), "HOLA ANDREA, YO SO GRAEME. YO SOY MEJOR DE JUSTIN EN EL JUEGO DE BEAT SABER"),
    ("amvasquez", "gqo", 1, NOW(), "Hi Graeme. My name is Andrea. Refactoring makes me sad :("),
    ("justin39", "rtr", 1, NOW(), "你好瑞思，我叫家四厅"),
    ("rtr", "justin39", 1, NOW(), "Nice to meet you. My name is Reece.");