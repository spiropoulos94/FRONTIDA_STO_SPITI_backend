-- check which roles can read or write which author
SELECT
    loggedUsers.Title,
    perms.Name,
    report_authors.Title
from
    Reports_permissions rp
    LEFT JOIN Roles loggedUsers ON rp.LoggedUserRole = loggedUsers.Role_id
    LEFT JOIN Permissions perms ON rp.Permission_id = perms.Permission_id
    LEFT JOIN Roles report_authors ON rp.Report_author_id = report_authors.Role_id;

-- check doctors (2) permissions 
select
    loggedUsers.Title,
    Permissions.Name,
    reportAuthors.Title
from
    Reports_permissions reports_perms
    left join Roles loggedUsers on reports_perms.LoggedUserRole = loggedUsers.Role_id
    left join Permissions on reports_perms.Permission_id = Permissions.Permission_id
    left join Roles reportAuthors on reports_perms.Report_author_id = reportAuthors.Role_id
where
    Reports_perms.LoggedUserRole = 2;

-- select available services for a report
select
    Daily_Reports.Report_id,
    Users.Role_id,
    Roles.Title,
    Services.Title
from
    Daily_Reports
    left join Users on Daily_Reports.User_id = Users.User_id
    left join Roles on Users.Role_id = Roles.Role_id
    left join Services on Daily_Reports.User_id = Services.Role_id;

-- select available services for a role
select
    Roles.Title,
    Services.Title
from
    Roles
    left join Services on Roles.Role_id = Services.Role_id;

-- select available permissions from a user to a report author
SELECT
    Permissions.Permission_id,
    Permissions.Name
FROM
    Reports_permissions
    LEFT JOIN Permissions ON Reports_permissions.Permission_id = Permissions.Permission_id
WHERE
    Reports_permissions.LoggedUserRole = 2
    AND Reports_permissions.Report_author_id = 3;