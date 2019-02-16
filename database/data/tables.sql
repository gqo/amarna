drop database tables;
create database tables;
use tables;

CREATE TABLE User
(
    username varchar(64),
    f_name varchar(64),
    l_name
    PRIMARY KEY (username)
);

CREATE Table Pairing 
(
    leftUser varchar(64),
    rightUser varchar(64),
    leftUserLang varchar(32),
    rightUserLang varchar(32),
    PRIMARY KEY (leftUser, rightUser),
    FOREIGN KEY (leftUser) REFERENCES User(username) on delete cascade,
    FOREIGN KEY (rightUser) REFERENCES User(username) on delete cascade,
    FOREIGN KEY (leftUserLang) REFERENCES Language(lang_name),
    FOREIGN KEY (rightUserLang) REFERENCES Language(lang_name)
);

CREATE TABLE Language
(
    lang_name varchar(32),
    PRIMARY KEY(lang_name)
);

CREATE TABLE KnownLang
(
    username varchar(64),
    lang_name varchar(32),
    PRIMARY KEY (username, lang_name),
    FOREIGN KEY (username) REFERENCES User(username) on delete cascade,
    FOREIGN KEY (lang_name) REFERENCES Language(lang_name)
);

CREATE TABLE LearnLang
(
    username varchar(64),
    lang_name varchar(32),
    PRIMARY KEY (username, lang_name),
    FOREIGN KEY (username) REFERENCES User(username) on delete cascade,
    FOREIGN KEY (lang_name) REFERENCES Language(lang_name)
);

CREATE TABLE Topic
(
    title varchar(128),
    lang_name varchar(32),
    length int,
    PRIMARY KEY (title, lang_name),
    FOREIGN KEY (lang_name) REFERENCES Language(lang_name)
);

CREATE TABLE Section
(
    topic_title varchar(128),
    topic_lang varchar(32),
    week int,
    description varchar(512),
    PRIMARY KEY (topic_title, topic_lang, week),
    FOREIGN KEY (topic_title) REFERENCES Topic(title) on delete cascade,
    FOREIGN KEY (topic_lang) REFERENCES Language(lang_name) on delete cascade
)

CREATE TABLE Letter
(
    leftUser varchar(64),
    rightUser varchar(64),
    date timestamp,
    topic_title varchar(128),
    topic_lang varchar(32),
    week int,
    body varchar(1024),
    PRIMARY KEY (leftUser, rightUser, topic_title, topic_lang, week),
    FOREIGN KEY(leftUser) REFERENCES User(username),
    FOREIGN KEY (rightUser) REFERENCES User(username),
    FOREIGN KEY (topic_title) REFERENCES Topic(title),
    FOREIGN KEY (topic_lang) REFERENCES Language(lang_name),
    FOREIGN KEY (week) REFERENCES Section(week)
);


SELECT "------------------------- Adding Data -------------------------" as "";
-- Insert dummy data

INSERT INTO User
    (username, f_name, l_name)
VALUES
    ("amvasquez", "Andrea", "Vasquez"),
    ("gqo", "Graeme", "Ferguson"),
    ("rtr", "Reece", "Rodriguez"),
    ("justin39", "Justin", "Wang"),
    ("andrew29", "Andrew", "Chen");

INSERT INTO Pairing
    (leftUser, rightUser, leftUserLang, rightUserLang)
VALUES
    ("amvasquez", "rtr", "Chinese", "Spanish"),
    ("rtr", "amvasquez", "Spanish", "Chinese"),
    ("gqo", "andrew29", "Chinese", "English"),
    ("andrew29", "gqo", "English", "Chinese"),
    ("justin39", "rtr", "Chinese", "English"),
    ("rtr", "justin39", "English", "Chinese");

INSERT INTO Language
    (lang_name)
VALUES
    ("English"),
    ("Spanish"),
    ("Chinese");

INSERT INTO KnownLang
    (username, lang_name)
VALUES 
    ("amvasquez", "English"),
    ("amvasquez", "Spanish"),
    ("rtr", "Chinese"),
    ("gqo", "English"),
    ("gqo", "Spanish"),
    ("andrew29", "Chinese"),
    ("justin39", "English");

INSERT INTO LearnLang
    (username, lang_name)
VALUES
    ("amvasquez", "Chinese"),
    ("gqo", "Chinese"),
    ("justin39", "Chinese"),
    ("andrew29", "English"),
    ("rtr", "English"),
    ("rtr", "Spanish");

INSERT INTO Topic
    (title, lang_name, length
VALUES
    ("Introduction", "English", 100),
    ("Introduction", "Spanish", 100),
    ("Introduction", "Chinese", 100);

INSERT INTO Section
    (topic_title, topic_lang, week, description)
VALUES
    ("Introduction", "English", 1, "introduction to your new Pal! Introduce your name, where you live, and occupation"),
    ("Introduction", "Chinese", 1, "introduction to your new Pal! Introduce your name, where you live, and occupation"),
    ("Introduction", "Spanish", 1, "introduction to your new Pal! Introduce your name, where you live, and occupation");
    
INSERT INTO Letter
    (leftUser, rightUser, date, topic_title, topic_lang, week, body)
VALUES
    ("rtr", "amvasquez", NOW(), "Introduction", "Spanish", 1, "HOLA ANDREA, YO SO REECE. YO SOY MEJOR DE JUSTIN EN EL JUEGO DE BEAT SABER"),
    ("justin39", "rtr", NOW(), "Introduction", "Chinese", 1, "你好瑞思，我叫家四厅");






