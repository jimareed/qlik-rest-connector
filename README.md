# qlik-rest-connector

A simple REST connector example for Qlik Sense Enterprise.

A simple go service that serves up xlsx files as a REST datasource.  

## Steps to import example REST data

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

##### 1. Create temp folder and copy data to that folder
```
$ mkdir /tmp/qlik-rest-connector
$ cp *.xslx /tmp/qlik-rest-connector
```

##### 2. Install Kubernetes service
```
$ helm install --name qlik-rest-connector ./qlik-rest-connector
$ kubectl get pods | grep qlik-rest-connector
qlik-rest-connector-7f4965f4b6-swb96                             1/1       Running   0          14s
$ kubectl port-forward qlik-rest-connector-7f4965f4b6-swb96 8080
Forwarding from 127.0.0.1:8080 -> 8080
```

##### 3. Update REST connector in App to point to new datasource

> browse to https://localhost:8080/data/<filename>

##### 4. Shutdown Kubernetes service when done

```
$ helm del --purge qlik-rest-connector
```

## Steps to test service
##### 1. Build and run REST connector
```
$ docker build --tag qlik-rest-connector .
$ docker-compose up or (or docker run -p 8080:8080 -d qlik-rest-connector)
$ docker-compose down (to shutdown connector)
```
