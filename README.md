# goUsermanagement
Go usermanagement group work assinment-5 , where user can signup , signin  with invitation code , generate invitation code , delete their profile &amp; update their profile.


# Steps to Run
1: git clone https://github.com/hirensavani-99/goUsermanagement

2: Navigate the the directory - "cd goUsermanagement"

3: Update the congiguration accordingly to set up the database

4: Command to build and run the server: "go run main.go"


# Configuration variables
CONTAINER_NAME="User management"
POSTGRES_USER="root"
POSTGRES_PASSWORD="Hiren123"
POSTGRES_DB="User_Management"
TABLE_CREATION_SQL="CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(255) UNIQUE NOT NULL, password_hash CHAR(60) NOT NULL); 
                     CREATE TABLE IF NOT EXISTS invitation_codes (id SERIAL PRIMARY KEY, code VARCHAR(20) UNIQUE NOT NULL, used BOOLEAN DEFAULT FALSE);"

git bash : 

open project : 

run ->  docker pull postgres
run  -> docker run --name <your-db-name> -e POSTGRES_USER=<yourusername> -e POSTGRES_PASSWORD=<yourpassword> -e POSTGRES_DB=<yourdbname> -p 5432:5432 -d postgres

run -> go get

run -> go run .