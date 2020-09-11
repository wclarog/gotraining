use dev_library;

DROP TABLE IF EXISTS Section;
DROP TABLE IF EXISTS Book;
DROP TABLE IF EXISTS Newspaper;
DROP TABLE IF EXISTS Magazine;
DROP TABLE IF EXISTS Material;


CREATE TABLE IF NOT EXISTS Material (
  uniqueCode VARCHAR(64),
  name VARCHAR(256),
  dateOfEmission DATE,
  numberOfPages INT,
  materialType TINYINT,
  
  PRIMARY KEY (uniqueCode)
);

CREATE TABLE IF NOT EXISTS Book (
  uniqueCode VARCHAR(64),
  authorName VARCHAR(256),
  genre VARCHAR(256),
  
  PRIMARY KEY (uniqueCode)
);

ALTER TABLE Book
ADD CONSTRAINT FK_book_material_uniqueCode FOREIGN KEY (uniqueCode)
REFERENCES Material (uniqueCode) ON DELETE NO ACTION;

CREATE TABLE IF NOT EXISTS Newspaper (
  uniqueCode VARCHAR(64),
  url VARCHAR(1024),
  
  PRIMARY KEY (uniqueCode)
);

ALTER TABLE Newspaper
ADD CONSTRAINT FK_newspaper_material_uniqueCode FOREIGN KEY (uniqueCode)
REFERENCES Material (uniqueCode) ON DELETE NO ACTION;

CREATE TABLE IF NOT EXISTS Magazine (
  uniqueCode VARCHAR(64),
  url VARCHAR(1024),
  
  PRIMARY KEY (uniqueCode)
);

ALTER TABLE Magazine
ADD CONSTRAINT FK_magazine_material_uniqueCode FOREIGN KEY (uniqueCode)
REFERENCES Material (uniqueCode) ON DELETE NO ACTION;

CREATE TABLE IF NOT EXISTS Section (
  uniqueCode VARCHAR(64),
  code VARCHAR(64),
  content VARCHAR(1024),
  
  PRIMARY KEY (uniqueCode, code)
);

ALTER TABLE Section
ADD CONSTRAINT FK_section_magazine_uniqueCode FOREIGN KEY (uniqueCode)
REFERENCES Magazine (uniqueCode) ON DELETE NO ACTION;



/*
-- column_name data_type(length) [NOT NULL] [DEFAULT value] [AUTO_INCREMENT] column_constraint;

CREATE TABLE table1
( number INT(11) AUTO_INCREMENT,
name VARCHAR(32) NOT NULL,
city VARCHAR(32),
age VARCHAR(7),
CONSTRAINT key1 PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS Material (
  name VARCHAR(20),
  owner VARCHAR(20),
  species VARCHAR(20),
  sex CHAR(1),
  birth DATE,
  death DATE
);
*/
