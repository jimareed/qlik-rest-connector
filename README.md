# qlik-rest-connector

A simple REST connector example for Qlik Sense Enterprise.

## Steps to import REST data

##### 1. Create new app and select add data

![Create new app](./doc/add-data.png)

##### 2. Select REST data source

![Connect to data source](./doc/connect-to-data-source.png)

##### 3. Enter REST URL and connection name
> Set URL to https://raw.githubusercontent.com/jimareed/qlik-rest-connector/master/example.json and enter a connection name.

![Create Connection](./doc/create-connection.png)

##### 4. Set Authentication to Anonymous and select Create

![Create Connection authentication](./doc/create-connection-authentication.png)


##### 5. Select 'data' and then select Add Data to import data

![Import data](./doc/import-data.png)

##### 6. Build Qlik Sense app
![Bug report app](./doc/bug-report-app.png)


## Steps to refresh REST data
##### 1. Build and run REST connector
```
$ docker build --tag qlik-rest-connector .
$ docker-compose up
$ docker-compose down (to shutdown connector)
```
