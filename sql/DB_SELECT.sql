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