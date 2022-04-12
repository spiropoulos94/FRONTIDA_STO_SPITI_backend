INSERT
    DUMMY DATA
INSERT INTO
    Addresses (Street, Number, City, Postal_Code)
VALUES
    ('Main street', '1', 'Αθήνα', '1010'),
    ('Katexaki', '2', 'Αθήνα', '1010'),
    ('Syggrou', '13', 'Αθήνα', '1010');

INSERT INTO
    Roles (Title)
VALUES
    ('Admin'),
    ('Doctor'),
    ('Nurse'),
    ('Family Helper'),
    ('Social Worker/ Psychologist'),
    ('Physiotherapist'),
    ('Patient');

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

INSERT INTO
    Services (Title, Role_id) Nurse services
VALUES
    ('Measurement of vital points', 3),
    ('Body wash', 3),
    ('Local ministration', 3),
    ('Intramuscular injections', 3),
    ('Sores - Injury treatment', 3),
    ('Catheter placement', 3),
    ('Enema', 3),
    ('Alimentation with Levin', 3),
    ('Prescription', 3),
    ('Medicine purchase', 3),
    ('Medical appointment', 3),
    Social Worker / Psychologist ('Psychological, Social Support', 5),
    ('Orientation to elders for their rights', 5),
    ('Support contacting the appropriate agency', 5),
    ('EFKA medical documentation submission', 5),
    Family Helper ('Yard Cleaning', 4),
    ('Sweeping/Mopping', 4),
    ('Meal preparation', 4),
    ('Food Supply', 4),
    ('Other', 4),
    Physiotherapist ('Physiotherapy', 6),
    ('Kinesiotherapy', 6),
    Doctor ('Prescription', 2),
    ('Clinical Examination', 2),
    ('Catheteriza', 2);

INSERT INTO
    Permissions (Name)
VALUES
    ('ADMIN'),
    ('READ'),
    ('WRITE');

INSERT INTO
    Patients (
        Fullname,
        Patient_AMKA,
        Health_security,
        Address_id
    )
VALUES
    (
        'Kwstas Makrakis',
        '212121',
        TRUE,
        1
    ),
    (
        'Panagiwths Stravaleksis',
        '212123',
        FALSE,
        2
    ),
    (
        'Kleomenhs Enas',
        '212127',
        TRUE,
        3
    ),
;

INSERT INTO
    Reports_permissions (LoggedUserRole, Permission_id, Report_author_id)
VALUES
    Admin permissions (2 :READ, 3 :WRITE) (1, 1),
    (2, 2, 2),
    (2, 3, 2),
    (2, 2, 5),
    (2, 2, 3),
    (2, 2, 6),
    (4, 2, 4),
    (4, 3, 4),
    (3, 2, 3),
    (3, 3, 3),
    (3, 2, 2),
    (3, 2, 5),
    (3, 2, 6),
    (6, 2, 6),
    (6, 3, 6),
    (6, 2, 5),
    (6, 2, 3),
    (6, 2, 2),
    (5, 2, 5),
    (5, 3, 5),
    (5, 2, 2),
    (5, 2, 4),
    (5, 2, 6),
    (5, 2, 3);

-- insert daily report
INSERT INTO
    Daily_Reports(
        User_id,
        Patient_id,
        Report_content,
        Report_date_ts,
        Arrival_Time_ts,
        Departure_Time_ts,
        Absence_Status
    )
VALUES
    (
        2,
        2,
        'Everything went as expected. Will meet again.',
        1618267056,
        1618612656,
        1618616256,
        FALSE
    )