Command to run the docker container for MySQL
docker run -d --name mysql1 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=13BCA006@shiats -e MYSQL_ROOT_HOST=% mysql/mysql-cluster mysqld
