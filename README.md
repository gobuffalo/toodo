# toodo - An example Buffalo App

This is an example Buffalo application using the trusted "todo" application.

# Download

```text
$ go get -u -v -t github.com/gobuffalo/toodo/...
$ cd $GOPATH/src/github.com/gobuffalo/toodo
```

## Setting up Assets

This application, like most "standard" Buffalo applications uses [NPM/Node](http://www.npmjs.com) to build and manage it's assets.

You will need to install Node and NPM, then install the dependencies for this application:

```text
$ npm install
```

# Database Setup

This application is backed by a `SQLite3` database, so there isn't much to get the database setup and running.

## Create The Databases

```text
$ buffalo db create -a -d
```

The `-a` flag will create all of the databases list in the `database.yml` file. The `-d` flag will print debugging information out so you can see what is being done.

Both flags are completely optional. If the `-a` flag is not present it will default to creating the `development` database.

## Migrate the Database

The next step is create all of the tables and schema information we need to run this application. Buffalo has a command to let us do that.

```text
$ buffalo db migrate
```

When run you should see output similar to the following:

```
v3.11.3

> 20170130231524_create_todos.up.fizz

0.0043 seconds
```

## Seed the Database

In the `grifts/seed.go` file is a `seed` task that will clear out any existing todos in the database and insert a new todo.

```text
$ buffalo task seed
```

# Run The Tests

Next, let's make sure everything works like we expect it to by running the tests:

```text
$ buffalo test
```

You should see an output similar to the following:

```text
Buffalo version 0.7.3

go test github.com/gobuffalo/toodo github.com/gobuffalo/toodo/actions github.com/gobuffalo/toodo/grifts github.com/gobuffalo/toodo/models
?   	github.com/gobuffalo/toodo	[no test files]
ok  	github.com/gobuffalo/toodo/actions	0.072s
?   	github.com/gobuffalo/toodo/grifts	[no test files]
ok  	github.com/gobuffalo/toodo/models	0.018s
```

# Running the application

Finally all that is left to do is run the application and play with it in a browser.

```text
$ buffalo dev
```

Now, head over to [http://127.0.0.1:3000](http://127.0.0.1:3000) and start playing with your new Buffalo application!

[Powered by Buffalo](http://gobuffalo.io)