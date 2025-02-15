# Minimal setup with Docker, MySQL and DBeaver

- `docker run -p 3307:3306 --name code-with-mosh-sql-course -e MYSQL_ROOT_PASSWORD=my-secret-pw mysql:8.0` with port 3307 to avoid duplicating ports with other containers
- Run SQL scripts inside WSL to create databases and tables
- Connect to the database via DBeaver with port 3307