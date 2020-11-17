# Backend

The backend is an optional piece of Retro-Carnage. The frontend makes various calls to the backend - but works fine if
those calls fail. As of now the frontend sends anonymous usage data and error reports to the backend. The backend
collects this information and stores it in a MongoDB database. Once the development of local single player and multi
player modes is completed, the backend will have a key role in organizing distributed multi player matches.

The backend is written in [Kotlin](https://kotlinlang.org/) and build upon the
[Spring Boot](https://spring.io/projects/spring-boot) framework.

## Structure

![Project structure](images/backend-structure.png)

| Component                                                           | Description                                                                                                     |
| ------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| **src/main/kotlin/net/retro-carnage/backend/configuration**         | Configuration classes for various parts of the application.                                                     |
| **src/main/kotlin/net/retro-carnage/backend/errors**                | Module used to record error reports.                                                                            |
| **src/main/kotlin/net/retro-carnage/backend/usage**                 | Module used to record anonymous usage data that gets used to generate statistics and identify usability issues. |
| **src/main/kotlin/net/retro-carnage/backend/BackendApplication.kt** | Entry point of the Spring Boot application.                                                                     |
| **pom.xml**                                                         | Maven configuration file that lists dependencies and configures the build process.                              |

## Build & Run

### Configure MongoDB connection

The backend uses a MongoDB database to store stuff. You can install & run your own MongoDB server - or just use
[MongoDB Atlas](https://www.mongodb.com/cloud/atlas). Atlas has a free tier with 500 MB of storage. That should be fine
for personal use.

When you signed up for Atlas and initialized your free MongoDB cluster, you'll get a connection string that you can use
to connect to the database. It looks something like this:

`mongodb+srv://<user>:<password>@<host>/<dbname>?retryWrites=true&w=majority`

The backend assembles this connection string and uses environment variables to replace the placeholders. Before you run
the backend you thus have to configure the following environment variables with the values from your MongoDB Atlas
cluster:

| Environment variable | Part of the connection string | Example                         |
| -------------------- | ----------------------------- | ------------------------------- |
| MONGO_DB_HOST        | host                          | myclustername.2abcd.mongodb.net |
| MONGO_DB_USER        | user                          | myusername                      |
| MONGO_DB_PASS        | password                      | supersecret102983               |

### Build & Run the source code

- Make sure to have a recent version of the JDK and Apache Maven installed. Development happens on OpenJDK 11.0.8 and
  Apache Maven 3.6.3 - so these should work fine.
- Get the latest source code from [GitHub](https://github.com/huddeldaddel/retro-carnage-backend).
- Open your command line, navigate to the project folder.
- Run `mvn clean package` to create a production build of the source code. This will create a deployable file located at
  **./target/retro-carnage-backend-<version>.jar**. This is the file you would deploy to services like
  [AWS Beanstalk](https://aws.amazon.com/en/elasticbeanstalk/) to run your own backend server.
- If you want to try out the code locally or want to develop, just import the project folder or the **pom.xml** in your
  favorite Java IDE. Kotlin support is best on [IntelliJ](https://www.jetbrains.com/idea/).

## Hosting

The backend is prepared for easy use with [Heroku](https://heroku.com) and
[AWS Beanstalk](https://aws.amazon.com/en/elasticbeanstalk/). If you're new to running services in the cloud then you
should give Heroku a try - it's easy to set up and has a free tier.
