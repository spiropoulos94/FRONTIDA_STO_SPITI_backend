DROP TABLE IF EXISTS Users;

DROP TABLE IF EXISTS Roles;

DROP TABLE IF EXISTS Services;

DROP TABLE IF EXISTS Addresses;

DROP TABLE IF EXISTS Permissions;

DROP TABLE IF EXISTS Daily_Reports;

DROP TABLE IF EXISTS Reports_services;

DROP TABLE IF EXISTS Patients;

-- DROP TABLE IF EXISTS Clients;
-- DROP TABLE IF EXISTS Addresses;
-- DROP TABLE IF EXISTS Addresses;
CREATE TABLE Roles (
    Role_id int AUTO_INCREMENT,
    Title varchar(255) UNIQUE,
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

CREATE TABLE Patients (
    Patient_id int AUTO_INCREMENT,
    Fullname varchar(255),
    Patient_AMKA int,
    Health_security boolean,
    Address_id int,
    PRIMARY KEY (Patient_id),
    FOREIGN KEY (Address_id) REFERENCES Addresses(Address_id)
);

CREATE TABLE Daily_Reports (
    Report_id int UNIQUE,
    User_id int NOT NULL,
    Patient_id int,
    Report_content longtext,
    Report_Date_ts int,
    Arrival_Time_ts int,
    Departure_Time_ts int,
    Absence_Status boolean,
    PRIMARY KEY (Report_id),
    FOREIGN KEY (User_id) REFERENCES Users(User_id),
    FOREIGN KEY (Patient_id) REFERENCES Patients(Patient_id)
);

CREATE TABLE Reports_services (
    Report_id int,
    Service_id int,
    PRIMARY KEY (Report_id, Service_id),
    FOREIGN KEY (Report_id) REFERENCES Daily_Reports(Report_id),
    FOREIGN KEY (Service_id) REFERENCES Services(Service_id)
);