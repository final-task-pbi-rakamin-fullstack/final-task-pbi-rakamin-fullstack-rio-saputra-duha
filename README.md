# GoLang API #
This API provides basic authentication for user and photo management. Only authorized users can create, edit, and delete their own photos. Each user can only manage their own photos, preventing others from making changes.
The project uses UUIDs for user and photo IDs, making them difficult to guess. It also employs JSON Web Tokens for authorization, enhancing security by preventing unauthorized users from modifying data.

## Features ##
<details>
        <summary>1. Json Web token</summary>
        <p><a href="https://jwt.io/">JWT</a> for users auth purposes</p>
</details>

<details>
        <summary>2. Validation</summary>
        <p><a href="https://github.com/asaskevich/govalidator">GoValidator</a> for validate user register</p>
</details>
<details>
        <summary>3. GodotEnv</summary>
        <p><a href="https://github.com/joho/godotenv">GodotEnv</a> used for store any sensitive data to local variable in <code>.env</code></p>
</details>
<details>
        <summary>4. Swagger</summary>
        <p> You can see all the API documentation fromated in <a href="https://swagger.io/">Swagger</a> at <a href="/docs/APIDocs.yaml">docs/APIDocs.yaml</a></p>
</details>
<details>
        <summary>5. UUID</summary>
        <p>Instead of using simple increment ID , using UUID is much simpler, and using <a href="https://github.com/jaswdr/faker">Faker</a> to automatically generate the UUID</p>
</details>
<details>
        <summary>6. Salted Secret Key</summary>
        <p>The secret key is salted so it much safer</p>
</details>

### Preparations ###
1. Make sure there is .env is present in the project, if there is no `.env` file, rename `env` file to `.env`
2. Setup the database `username` in `DBUSER` variable and `password` in `DBPASSWORD` in the `.env` file
4. Make sure there is Mysql database named `db_goapi` in your local machine or you can edit the `DBNAME` inside `.env` file

### Installing the project ###
#### 1. First clone this project ####
```
git clone https://github.com/Cherno60/final-task-pbi-rakamin-fullstack-Rafly-Andrian-Wicaksana.git
```
#### 2. Navigate to inside the project ####
```bash
cd final-task-pbi-rakamin-fullstack-Rafly-Andrian-Wicaksana
```
#### 3. Install all the dependency ####
Run this cmd prompt to install all the dependency needed
```go
go mod tidy
```
#### 4. Run the project ####
Close any `9090` port first if you have it running from another project, or you can modify the port inside the `.env` file 
in `PORT` variable
,run this cmd prompt to run the project and start localserver
```golang
go run main.go
```
after you run it, you can access it in 
```bash
http://localhost:9090
```
#### 5. Test the API ####
You can test the API using [Postman](https://www.postman.com/) by clicking this button

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/9886572-874213a2-d42c-4050-8694-61bed3fa78f8?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D9886572-874213a2-d42c-4050-8694-61bed3fa78f8%26entityType%3Dcollection%26workspaceId%3D003a7aa7-8e1c-48bf-bac7-04d702568a60)
## Endpoint ##
Any endpoint of this API can be accessed in here

## User API ##


### User Registration ###
**Endpoint : POST**

URL : /users/register

Request Body:
```json
{
  "username": "username",
  "email": "example@gmail.com",
  "password": "password"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "User Created",
    "errors": null
}
```

### User Login ###
**Endpoint : POST**

URL : /users/login

Request Body:
```json
{
	"email": "email@email.com",
	"password": "password"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "Logined successfully!, Hello {Username}"
}
```

### User Update ###
**Endpoint : PUT**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/edit/{userid}

Request Body:
```json
{
  "username": "newUsername",
  "email": "new@gmail.com",
}
```
Response Body success :
```json
{
    "Errors": null,
    "Messages": {
        "Username": "newUsername",
        "Email": "new@gmail.com",
        "Password": ""
    },
    "Status": 200
}
```

### User Delete ###
**Endpoint : DELETE**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/delete/{userid}
Request Parameter: 
`{userid}`
Response Body success :
```json
{
    "status": 200,
    "messages": "User Deleted",
    "errors": null
}{
    "Message": "User Logouted Successfully"
}
```

### User Logout ###
**Endpoint : GET**

URL : /users/logout

Response Body success :
```json
{
    "Message": "User Logouted Successfully"
}
```

## Photo API ##


### Photo Index ###
**Endpoint : GET**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos

Response Body success :
```json
{
    "Errors": null,
    "Photo Data": [
        {
            "uuid": "f4abf47c-07f1-4a7e-b250-9af9420af671",
            "title": "Photos",
            "caption": "loresdsadmsd",
            "photo_url": "google.com",
            "user_id": "38157368-9e2d-4889-9c6f-1c6f8a74a6d5",
            "User": null,
            "created_at": "2024-06-01T21:53:18.766+07:00",
            "updated_at": "2024-06-01T21:53:18.766+07:00"
        }
    ],
    "Status": 200
}
```

### Photo Create ###
**Endpoint : POST**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos/add

Request Body:
```json
{
   "title" : "Photos",
  "caption": "loresdsadmsd",
  "photo_url": "google.com"
}
```
Response Body success :
```json
{
    "status": 200,
    "messages": "Photo Created",
    "errors": null
}
```

### Photo Edit ###
**Endpoint : PUT**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /photos/edit/{photoid}

Request Body:
```json
{
  "title" : "newTitle",
  "caption": "CaptionNew",
  "photoUrl": "newUrl.com"
}
```
Response Body success :
```json
{
    "Errors": null,
    "Messages": {
        "Title": "newTitle",
        "Caption": "CaptionNew",
        "PhotoUrl": "newUrl.com"
    },
    "Status": 200
}
```


### Photo Delete ###
**Endpoint : DELETE**
<br>
JWT using cookie to save the Token
Cookie :
- UserData: token

URL : /users/delete/{userid}
Request Parameter: 
`{photoid}`
Response Body success :
```json
{
    "Errors": null,
    "Messages": "Photo with caption CaptionNew has been deleted",
    "Status": 200
}
```


