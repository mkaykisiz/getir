
# Getir

Getir Case

<p>
    <a href="https://app.getpostman.com/run-collection/2ecb07d7ce3cef97e8a7"><img src="https://run.pstmn.io/button.svg" alt="getir"></a>
</p>


# Connection
**Host:**
```text
https://calm-ocean-71421.herokuapp.com/
```

**Router:**
- `[POST] /records  | https://calm-ocean-71421.herokuapp.com/records`
- `[GET, POST] /in-memory | https://calm-ocean-71421.herokuapp.com/in-memory`

#Install
create .env file or set environment variables
```text
PORT=8000
MONGO_DB_URL=
MONGO_DB_NAME=
MONGO_COLLECTION_NAME=
```

**Docker install**
```text
$ docker build -t getir .
$ docker run -d -p 8000:8000 --env-file .env getir
```

**Heroku Deploy**
```text
heroku login
heroku create
git push heroku master
heroku open
```
**View logs**
```text
heroku logs --tail
```

**Heroku set env**
```text
heroku config:set PORT=8000
heroku config:set MONGO_DB_URL=
heroku config:set MONGO_DB_NAME=
heroku config:set MONGO_COLLECTION_NAME=

```


# Services
## Records

**URL** : `/records`

**Method** : `POST`

**Body**

```json
{
  "startDate": "2016-01-21",
  "endDate": "2016-03-02",
  "minCount": 2900,
  "maxCount": 3000
}
```

### Responses
#### Success Response

**Code** : `200`

**Response Body**

```json
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "key": "asbjlafla8asf75asf",
      "createdAt": "2019-03-12T02:45:42.111+02:00",
      "totalCount": 23
    }
  ]
}
```

#### Error Response

**Code** : `400`

**Response Body**

```json
{
  "code": 1,
  "msg": "Decode Error",
  "records": []
}
```

## In-Memory Save Data 

**URL** : `/in-memory`

**Method** : `POST`

**Body**

```json
{
    "key": "active-tabs",
    "value": "getir"
}
```

### Responses
#### Success Response

**Code** : `201`

**Response Body**

```json
{
  "key": "active-tabs",
  "value": "getir"
}
```

## In-Memory Get Data

**URL** : `/in-memory?key=active-tabs`

**Method** : `POST`

**Querystring**

```text
?key=active-tabs
```

### Responses
#### Success Response

**Code** : `200`

**Response Body**

```json
{
  "key": "active-tabs",
  "value": "getir"
}
```
