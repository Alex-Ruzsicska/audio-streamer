CREATE TABLE IF NOT EXISTS folder (
  	id intEGER PRIMARY KEY AUTOINCREMENT,
  	name text not NULL,
  	cover_url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS audio (
  	id integer PRIMARY KEY AUTOINCREMENT,
  	folder_id INT,
  	name text not NULL,
  	url TEXT NOT NULL,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (folder_id)
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
  		REFERENCES genre (genre_id)
  			ON DELETE CASCADE,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (folder_id)
  			ON DELETE CASCADE
  
);

CREATE TABLE IF NOT EXISTS author_folder (
  	author_id int,
  	folder_id int,
  	PRIMARY KEY (author_id, folder_id),
  	FOREIGN KEY (author_id)
  		REFERENCES author (author_id)
  			ON DELETE CASCADE,
  	FOREIGN KEY (folder_id)
  		REFERENCES folder (folder_id)
  			ON DELETE CASCADE
  
);

INSERT INTO folder (name, cover_url) 
VALUES("Test Folder", "/mock/cover/test.jpg");

INSERT INTO audio (folder_id, name, url) 
VALUES(1, "Test Audio", "/mock/audios/test");

INSERT INTO genre (name) 
VALUES("Test Genre");

INSERT INTO author (name) 
VALUES("Test Author");

INSERT INTO author_folder (author_id, folder_id) 
VALUES(1, 1);

INSERT INTO genre_folder (genre_id, folder_id) 
VALUES(1, 1);