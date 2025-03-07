# Go Texter

A simple server texting app that supports live updates.

Eventually, the app will have a basic front end that will have a chat room that can then enter and write messages for (other users should be able to see this in real time).

- [ ] Index page that shows all chat rooms.
- [ ] Able to sign up for an account through basic username and password.
- [ ] Able to create chat rooms (only available for users logged in).
- [ ] Anyone can view and enter chat rooms.
- [ ] Users logged in should be able to write new messages.
- [ ] Implement Websockets to showcase live updates.

## Dependencies

- direnv: This will allow the app to load in env variables to the project.
- docker-compose and docker: help build sub-resources needed to run the project (psql)

## Current Dev Setup

Before anything, run `direnv allow` to allow the environments on to this project.

### Database

You can either have one locally installed or use the [docker-compose.yml](/docker-compose.yml) file to spin one up for you.

Afterwards, you'll need to run the sql file [create_user.sql](/config/create_user.sql) to create the user.

Then you should be able to run this go command:

```bash
go run config/db_schema.go
```

This will build the tables needed for the app.

#### TODOs

- [ ] Need to put the create_user sql file into a go file. This way, there can be a single command to create this.
    - Or, I can make the Dockerfile for psql just use the gtexter variables.
- [ ] Need to configure this to be more of a migration setting. Right now, it will create new tables (but you'll need to manually drop them and recreate the table if needed).