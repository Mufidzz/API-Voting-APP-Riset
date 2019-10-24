# Documentation

## User API
###### Creating New User

> Request

URL             : http://url/user
Request Method  : POST

```json
{
    "FirstName" : "New",
    "LastName"  : "User",
    "Username"  : "NewUser",
    "Email"     : "email@newuser.com",
    "Password"  : "the password"
}
```

> Response

```json
{
    "data": {
        "ID": 2,
        "CreatedAt" : "2019-10-24T21:24:05.4949625+07:00",
        "UpdatedAt" : "2019-10-24T21:24:05.4949625+07:00",
        "DeletedAt" : null,
        "FirstName" : "New",
        "LastName"  : "User",
        "Username"  : "NewUser",
        "Email"     : "email@newuser.com",
        "Password"  : "$2a$10$3Sk5VwZwV408nW6CtLZ7yew9wOWUEn3IFxfxKQ2t5Y0dVSDKS0HXS"
    },
    "message": "Success Creating User",
    "status": 200
}
```

###### Read All User
> Request 

URL             : http://url/user/:id
Request Method  : POST

> Response

```json
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2019-10-24T00:00:00+07:00",
            "UpdatedAt": "2019-10-24T00:00:00+07:00",
            "DeletedAt": null,
            "FirstName": "a",
            "LastName": "a",
            "Username": "a",
            "Email": "a",
            "Password": "awdaw"
        },
        {
            "ID": 2,
            "CreatedAt": "2019-10-24T21:24:05+07:00",
            "UpdatedAt": "2019-10-24T21:24:05+07:00",
            "DeletedAt": null,
            "FirstName": "ax",
            "LastName": "as",
            "Username": "aq",
            "Email": "asd",
            "Password": "$2a$10$3Sk5VwZwV408nW6CtLZ7yew9wOWUEn3IFxfxKQ2t5Y0dVSDKS0HXS"
        },
        {
            "ID": 3,
            "CreatedAt": "2019-10-24T22:10:21+07:00",
            "UpdatedAt": "2019-10-24T22:10:21+07:00",
            "DeletedAt": null,
            "FirstName": "ax",
            "LastName": "as",
            "Username": "aq",
            "Email": "asd",
            "Password": "$2a$10$EBwBewk/5UD4YAbrznDF5emvUdn3XKb0NCs4Cj4WLhgf0/yuXeqWS"
        },
        ....
    ],
    "data_count": 6,
    "message": "Success Retrieving User Data",
    "status": 200
}
```

###### Read Specific User
###### Update User Data
###### Delete User

