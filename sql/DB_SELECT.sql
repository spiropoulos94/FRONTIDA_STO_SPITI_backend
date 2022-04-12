-- check which roles can read or write which author
SELECT
    roles1.Title,
    perms.Name,
    roles2.Title
from
    Reports_permissions rp
    LEFT JOIN Roles roles1 ON rp.LoggedUserRole = roles1.Role_id
    LEFT JOIN Permissions perms ON rp.Permission_id = perms.Permission_id
    LEFT JOIN Roles roles2 ON rp.Report_author_id = roles2.Role_id;