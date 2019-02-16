drop database tables;
create database tables;
use tables;

CREATE TABLE User
(
    username varchar(64),
    f_name varchar(64)
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
    FOREIGN KEY (rightUser) REFERENCES User(username) on delete cascade
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
    FOREIGN KEY (username) REFERENCES User(username),
    FOREIGN KEY (lang_name) REFERENCES Language(lang_name)
);

CREATE TABLE LearnLang
(
    username varchar(64),
    lang_name varchar(32),
    PRIMARY KEY (username, lang_name),
    FOREIGN KEY (username) REFERENCES User(username),
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
    FOREIGN KEY (topic_title) REFERENCES Topic(title),
    FOREIGN KEY (topic_lang) REFERENCES Language(lang_name)
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



