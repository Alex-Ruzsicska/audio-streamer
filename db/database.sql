CREATE TABLE IF NOT EXISTS folder (
  	id integer PRIMARY KEY AUTOINCREMENT,
  	name text not NULL,
  	cover_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS audio (
  	id integer PRIMARY KEY AUTOINCREMENT,
  	folder_id INT,
  	name text not NULL,
  	url TEXT NOT NULL,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (id)
  			ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS genre (
  	id integer PRIMARY KEY AUTOINCREMENT,
  	name text not NULL
);

CREATE TABLE IF NOT EXISTS author (
  	id integer PRIMARY KEY AUTOINCREMENT,
  	name text not NULL
);

CREATE TABLE IF NOT EXISTS genre_folder (
  	genre_id int,
  	folder_id int,
  	PRIMARY KEY (genre_id, folder_id),
  	FOREIGN KEY (genre_id)
  		REFERENCES genre (id)
  			ON DELETE CASCADE,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (id)
  			ON DELETE CASCADE
  
);

CREATE TABLE IF NOT EXISTS author_folder (
  	author_id int,
  	folder_id int,
  	PRIMARY KEY (author_id, folder_id),
  	FOREIGN KEY (author_id)
  		REFERENCES author (id)
  			ON DELETE CASCADE,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (id)
  			ON DELETE CASCADE
  
);

INSERT INTO folder (name, cover_url) 
VALUES("Test Folder", "/mock/cover/test.jpg");

INSERT INTO folder (name, cover_url) 
VALUES("Test Folder 2", "/mock/cover/test.jpg");

INSERT INTO folder (name, cover_url) 
VALUES("Test Folder 3", "/mock/cover/test.jpg");

INSERT INTO audio (folder_id, name, url) 
VALUES(1, "Test Audio", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(1, "Test Audio 2", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(1, "Test Audio 3", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(2, "Test Audio 4", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(2, "Test Audio 5", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(3, "Test Audio 6", "/mock/audios/test");

INSERT INTO audio (folder_id, name, url) 
VALUES(3, "Test Audio 7", "/mock/audios/test");

INSERT INTO genre (name) 
VALUES("Test Genre");

INSERT INTO genre (name) 
VALUES("Test Genre 2");

INSERT INTO genre (name) 
VALUES("Test Genre 3");

INSERT INTO author (name) 
VALUES("Test Author");

INSERT INTO author (name) 
VALUES("Test Author 2");

INSERT INTO author (name) 
VALUES("Test Author 3");

INSERT INTO author_folder (author_id, folder_id) 
VALUES(1, 1);

INSERT INTO author_folder (author_id, folder_id) 
VALUES(2, 2);

INSERT INTO author_folder (author_id, folder_id) 
VALUES(3, 3);

INSERT INTO genre_folder (genre_id, folder_id) 
VALUES(1, 1);

INSERT INTO genre_folder (genre_id, folder_id) 
VALUES(2, 2);

INSERT INTO genre_folder (genre_id, folder_id) 
VALUES(3, 3);