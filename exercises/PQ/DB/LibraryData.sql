use dev_library;

-- materialType: 0: book, 1: newspaper, 2: magazine
INSERT INTO Material (uniqueCode, name, dateOfEmission, numberOfPages, materialType)
VALUES
       ('1234', 'name 1234',  '2018-01-09', 534, 0),
       ('3432', 'name 3432',  '2017-02-12', 436, 0),
       ('6785', 'name 6785',  '2011-09-26', 456, 1),
       ('458', 'name 458',  '2088-04-03', 234, 0),
       ('24', 'name 24',  '2000-10-14', 84, 2),
       ('9766', 'name 9766',  '1994-06-28', 765, 1),
       ('8445', 'name 8445',  '2005-12-11', 125, 0),
       ('945', 'name 945',  '2002-02-07', 824, 2),
       ('90', 'name 90',  '1985-11-03', 342, 1),
       ('943', 'name 943',  '2019-08-21', 642, 0),
       ('894', 'name 894',  '2011-04-14', 55, 2),
       ('98', 'name 98',  '2013-07-04', 64, 2);

INSERT INTO Book (uniqueCode, authorName, genre)
VALUES
       ('1234', 'authorName 1234',  'genre 1234'),
       ('3432', 'authorName 3432',  'genre 3432'),
       ('458', 'authorName 458',  'genre 458'),
       ('8445', 'authorName 8445',  'genre 8445'),
       ('943', 'authorName 943',  'genre 943');

INSERT INTO Newspaper (uniqueCode, url)
VALUES
       ('6785', 'http://link6785.com/'),
       ('9766', 'http://link9766.com/'),
       ('90', 'http://link90.com/');

INSERT INTO Magazine (uniqueCode, url)
VALUES
       ('24', 'http://link24.com/'),
       ('945', 'http://link945.com/'),
       ('894', 'http://link894.com/'),
       ('98', 'http://link98.com/');


INSERT INTO Section (uniqueCode, code, content)
VALUES
       ('24', 'code 24 1', 'content 98 1'),
       ('24', 'code 24 2', 'content 98 2'),
       ('24', 'code 24 3', 'content 98 3'),
       ('945', 'code 945 1', 'content 98 1'),
       ('945', 'code 945 2', 'content 98 2'),
       ('945', 'code 945 3', 'content 98 3'),
       ('945', 'code 945 4', 'content 98 4'),
       ('894', 'code 894 1', 'content 98 1'),
       ('894', 'code 894 2', 'content 98 2'),
       ('894', 'code 894 3', 'content 98 3'),
       ('98', 'code 98 1', 'content 98 1'),
       ('98', 'code 98 2', 'content 98 2'),
       ('98', 'code 98 3', 'content 98 3'),
       ('98', 'code 98 4', 'content 98 4'),
       ('98', 'code 98 5', 'content 98 5');
