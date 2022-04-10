DROP TABLE IF EXISTS Users;

DROP TABLE IF EXISTS Roles;

DROP TABLE IF EXISTS Services;

DROP TABLE IF EXISTS Addresses;

DROP TABLE IF EXISTS Permissions;

DROP TABLE IF EXISTS Daily_Reports;

-- DROP TABLE IF EXISTS Phones;
-- DROP TABLE IF EXISTS Clients;
-- DROP TABLE IF EXISTS Addresses;
-- DROP TABLE IF EXISTS Addresses;
CREATE TABLE Roles (
    Role_id int AUTO_INCREMENT,
    Title varchar(255),
    PRIMARY KEY (Role_id)
);

CREATE TABLE Users (
    User_id int AUTO_INCREMENT,
    Name varchar(255),
    Surname varchar(255),
    AFM int NOT NULL UNIQUE,
    AMKA int NOT NULL UNIQUE,
    Email varchar(255),
    Password varchar(255),
    Role_id int,
    PRIMARY KEY (User_id),
    FOREIGN KEY (Role_id) REFERENCES Roles(Role_id)
);

-- every service belongs to a role/profession 
CREATE TABLE Services (
    Service_id int AUTO_INCREMENT,
    Title varchar(255),
    Role_id int,
    PRIMARY KEY (Service_id),
    FOREIGN KEY (Role_id) REFERENCES Roles(Role_id)
);

CREATE TABLE Permissions (
    Permission_id int AUTO_INCREMENT,
    Name varchar(255),
    PRIMARY KEY (Permission_id)
);

CREATE TABLE Roles_permissions (
    Role_id int,
    Permission_id int,
    FOREIGN KEY (Role_id) REFERENCES Roles(Role_id),
    FOREIGN KEY (Permission_id) REFERENCES Permissions(Permission_id)
);

CREATE TABLE Addresses(
    Address_id int AUTO_INCREMENT,
    Street varchar(255),
    Number int,
    City varchar(255),
    Postal_code int,
    PRIMARY KEY (Address_id)
);

CREATE TABLE Daily_Reports (
    Report_id int UNIQUE,
    User_id int NOT NULL,
    Patient_fullname varchar(255),
    Patient_AMKA int,
    Patient_health_security boolean,
    Patient_address int,
    Report_Date date,
    Arrival_Time datetime,
    Departure_Time datetime,
    Absence_Status boolean,
    PRIMARY KEY (Report_id),
    FOREIGN KEY (User_id) REFERENCES Users(User_id),
    FOREIGN KEY (Patient_address) REFERENCES Addresses(Address_id)
);

-- CREATE TABLE RoomCategories (
--     RoomCategoryID int AUTO_INCREMENT,
--     Name varchar(255) UNIQUE,
--     RoomPrice int,
--     RoomCount int,
--     PRIMARY KEY (RoomCategoryID)
-- );
-- CREATE TABLE Rooms (
--     RoomID int UNIQUE,
--     RoomCategory int NOT NULL,
--     RoomNumber int,
--     FloorNumber int,
--     SquareMeters int,
--     BookTimes int,
--     PRIMARY KEY (RoomID),
--     FOREIGN KEY (RoomCategory) REFERENCES RoomCategories(RoomCategoryID)
-- );
-- ALTER TABLE
--     Rooms
-- ADD
--     CONSTRAINT RoomNumber UNIQUE NONCLUSTERED(RoomNumber, RoomCategory);
-- CREATE TABLE Bookings (
--     BookID int AUTO_INCREMENT,
--     BookInDate DATE,
--     BookOutDate DATE,
--     Bill int,
--     ClientID int,
--     PRIMARY KEY (BookID),
--     FOREIGN KEY (ClientID) REFERENCES Clients(ClientID)
-- );
-- ALTER TABLE
--     Bookings
-- ADD
--     CONSTRAINT ClientID UNIQUE NONCLUSTERED(ClientID, BookInDate, BookOutDate);
-- CREATE TABLE Facilities (
--     FacilityID int AUTO_INCREMENT,
--     Name varchar(255) UNIQUE,
--     DailyPrice int,
--     PRIMARY KEY (FacilityID)
-- );
-- CREATE TABLE Rooms_Facilities (
--     RoomID int,
--     FacilityID int,
--     FOREIGN KEY (RoomID) REFERENCES Rooms(RoomID),
--     FOREIGN KEY (FacilityID) REFERENCES Facilities(FacilityID)
-- );