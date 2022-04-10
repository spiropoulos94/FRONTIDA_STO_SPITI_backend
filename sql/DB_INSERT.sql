-- INSERT DUMMY DATA
-- INSERT INTO
--     Addresses (Street, Number, City, Postal_Code)
-- VALUES
--     ('Main street', '1', 'Αθήνα', '1010'),
--     ('Katexaki', '2', 'Αθήνα', '1010'),
--     ('Syggrou', '13', 'Αθήνα', '1010');
-- INSERT INTO
--     Roles (Title)
-- VALUES
--     ('Admin'),
--     ('Doctor'),
--     ('Nurse'),
--     ('Fmily Helper'),
--     ('Social Worker/ Psychologist'),
--     ('Physiotherapist'),
--     ('Patient');
INSERT INTO
    Users (Name, Surname, AFM, AMKA, Role_id)
VALUES
    ('John', 'Doe', '1111111', '88888', 1),
    ('James', 'Brown', '2222222', '99999', 2),
    ('Oscar', 'Scoffield', '3333333', '10101010', 3),
    ('Michael', 'Daglas', '444444', '12121212', 4),
    ('Vanessa', 'Smith', '555555', '13131313', 5),
    ('Joanna', 'Downing', '6666666', '141414', 6),
    (
        'Hailey',
        'Smurf',
        '12345677',
        '14141414',
        7
    );