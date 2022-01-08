/* CREATE USER 'contact'@'%' IDENTIFIED BY ''; */
/* DROP USER 'contact'; */

REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'contact'@'%';

GRANT SELECT ON contact.* TO 'contact'@'%';
GRANT INSERT ON contact.messages TO 'contact'@'%';

FLUSH PRIVILEGES;