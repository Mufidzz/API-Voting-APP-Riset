# Documentation

## User API
###### Creating New User

> Request

Request Method  : POST

```
{
    "FirstName" : "New",
    "LastName"  : "User",
    "Username"  : "NewUser",
    "Email"     : "email@newuser.com",
    "Password"  : "the password"
}
```

> Response

```
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

###### Read Specific User
###### Update User Data
###### Delete User

