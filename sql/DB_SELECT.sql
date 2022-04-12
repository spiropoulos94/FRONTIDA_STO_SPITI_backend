SELECT
    *
from
    Reports_permissions
    LEFT JOIN Roles ON Reports_permissions.Permission_id = Roles.Role_id
    LEFT JOIN Permissions ON Reports_permissions.Permission_id = Permissions.Permission_id
    LEFT JOIN Roles on Roles.Role_id = Reports_permissions.Report_author_id;