# CRUD Sample

This is a sample implementation of CRUD (Create, Read, Update, Delete) in Go.

The sample is using SQLite for persistence and provides the following API endpoints via HTTP:

- `GET /items` - To retrieve a list of all items
- `GET /items/:id` - To retrieve a item using its identifier
- `POST /items` - To create a new item
- `PUT /items/:id` - To update an existing item using its identifier
- `DELETE /items/batch` - To delete multiple items providing an array of identifiers as payload
- `DELETE /items/:id` - To delete an existing item using its identifier

Send data to `POST /items` and `PUT /items/:id` using the following structure:

```jsonc
{
    "name": "item name",
    // boolean (either true or false)
    "active": true
}
```
## Running this Sample locally 

To run the sample locally, you must provide the `migrations.sql` using the `--sqlite` flag as shown in the snippet below:

```bash
# Build the project
spin build

# Run the sample
spin up --sqlite @migrations.sql
Logging component stdio to ".spin/logs/"
Storing default SQLite data to ".spin/sqlite_db.db"

Serving http://127.0.0.1:3000
Available Routes:
  crud: http://127.0.0.1:3000 (wildcard)
```

### Fermyon Cloud

You can deploy this sample to Fermyon Cloud following the steps below:

```bash
# Authenticate
spin cloud login

# Deploy the sample to Fermyon Cloud
# This will ask if a new database should be created or an existing one should be used
# Answer the question with "create a new database"
spin deploy
Uploading crud version 0.1.0 to Fermyon Cloud...
Deploying...
App "crud" accesses a database labeled "default"
    Would you like to link an existing database or create a new database?: Create a new database and link the app to it
What would you like to name your database?
    Note: This name is used when managing your database at the account level. The app "http-crud-js-sqlite" will refer to this database by the label "default".
    Other apps can use different labels to refer to the same database.: sincere-mulberry
Creating database named 'sincere-mulberry'
Waiting for application to become ready.......... ready

View application:   https://http-crud-js-sqlite-jcmbpezb.fermyon.app/
Manage application: https://cloud.fermyon.com/app/http-crud-js-sqlite

# Ensure tables are created in the new database (here sincere-mulberry)
spin cloud sqlite execute --database sincere-mulberry @migrations.sql
```