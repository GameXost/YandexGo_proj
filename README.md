# Creating a PostgreSQL Database in pgAdmin

To create a PostgreSQL database named "users" with the specified parameters, follow these steps:

1. **Open pgAdmin**:
   - Launch pgAdmin from your applications.

2. **Connect to the PostgreSQL Server**:
   - In the pgAdmin interface, expand the "Servers" node in the left sidebar.
   - Right-click on your PostgreSQL server and select "Connect".

3. **Create a New Database**:
   - Right-click on the "Databases" node under your server.
   - Select "Create" and then "Database...".

4. **Configure Database Settings**:
   - In the "General" tab, enter the following:
     - **Database**: `users`
   - Click on the "Save" button to create the database.

5. **Set Up User Roles (if needed)**:
   - If you need to create a user with the username "test" and password "0000":
     - Expand the "Login/Group Roles" node under your server.
     - Right-click and select "Create" > "Login/Group Role...".
     - Enter the following:
       - **Role name**: `test`
       - **Password**: `0000`
     - In the "Definition" tab, ensure the "Can login?" option is checked.
     - Click "Save".

6. **Grant Permissions**:
   - After creating the user, you may need to grant permissions to the "users" database:
     - Right-click on the "users" database and select "Properties".
     - Go to the "Privileges" tab.
     - Click on the "Add" button and select the user `test`.
     - Set the appropriate privileges (e.g., CONNECT, CREATE, TEMPORARY).
     - Click "Save".

7. **Connect to the Database**:
   - You can now connect to the "users" database using the credentials:
     - **Username**: `test`
     - **Password**: `0000`
     - **Host**: `localhost`
     - **Port**: `5432`

Now you have successfully created a PostgreSQL database named "users" with the specified parameters.
