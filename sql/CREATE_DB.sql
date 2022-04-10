DROP TABLE IF EXISTS Users;

DROP TABLE IF EXISTS Roles;

-- DROP TABLE IF EXISTS Phones;
-- DROP TABLE IF EXISTS Clients;
-- DROP TABLE IF EXISTS Addresses;
-- DROP TABLE IF EXISTS Phones;
-- DROP TABLE IF EXISTS Clients;
-- DROP TABLE IF EXISTS Addresses;
-- DROP TABLE IF EXISTS Addresses;
CREATE TABLE Roles (
    ID int AUTO_INCREMENT,
    Title varchar(255),
    PRIMARY KEY (ID)
);

CREATE TABLE Users (
    ID int AUTO_INCREMENT,
    Name varchar(255),
    Surname varchar(255),
    AFM int NOT NULL UNIQUE,
    AMKA int NOT NULL UNIQUE,
    Email varchar(255),
    Password varchar(255),
    Role_id int,
    PRIMARY KEY (ID),
    FOREIGN_KEY (Role_id) REFERENCES Roles(ID)
);

-- CREATE TABLE Phones (
--     PhoneID int AUTO_INCREMENT,
--     PhoneNumber int UNIQUE,
--     ClientID int,
--     PRIMARY KEY (PhoneID),
--     FOREIGN KEY (ClientID) REFERENCES Clients(ClientID)
-- );
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