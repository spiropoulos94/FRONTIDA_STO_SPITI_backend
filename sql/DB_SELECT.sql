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
    Reports_perms.LoggedUserRole = 2