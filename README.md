## Simple service implementing table booking functionality
 
## Install
1. Clone repo outside of GOPATH
2. Run:

       go build
     
## Operation
The service will build into `tablebooker` binary.

### Database initialization
    ./tablebooker initdb 
Initialization will create an SQLite db file `tables.db` into same directory, 
which will be used to store the data. Note that it will drop all existing data
if exists

### Starting the service
    ./tablebooker server
    
### API reference
When the server is running, it implements the following API:
#### Get the list of all tables
Example request

    GET /table/list
Example response
```json5
{
    "head": {
        "timestamp": "2019-09-18T08:25:06.646617Z" // request timestampp
    },
    "body": [
    // list of all tables below
        {
            "id": "77146dc9-df4b-46a9-a52a-3d36cc0a22a4",
            "number": 1,
            "capacity": 3
        },
        {
            "id": "e4d17aef-fc56-4679-a905-bf02c09eee2e",
            "number": 2,
            "capacity": 9
        }
     ]
}
```

#### Book the table by its id
Example request

    POST /table/book/{table_id}    
```json5
{
    "bookDate": "2019-10-07",
    "guestName":"Alex",
    "guestContact":"Heyhey"
}
```

Example success response
```json5
{
    "head": {
        "timestamp": "2019-09-18T07:45:06.287185Z"
    },
    "body": {
        "id": "93b2fba1-8706-48b0-ac23-98f32b59df98",
        "tableId": "2d06220d-962f-4ab0-8847-78e9e3b45d83",
        "bookDate": "2019-10-07",
        "guestName": "Alex",
        "guestContact": "Heyhey",
        "code": "ae134cf7-c4e5-4b34-8f2a-2c5ba044af2d"
    }
}
```

Example error response
    
    409 Conflict
    table already booked 

#### Unbook table using the booking copde
Example request

    DELETE /table/book/{code}

Example success response
```json5
{
    "head": {
        "timestamp": "2019-09-18T08:30:55.644159Z"
    },
    "body": "OK"
}
```

Example error response

    404 Not Found
    no booking with code found