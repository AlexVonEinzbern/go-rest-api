# go-rest-api
ecommerce api-rest with basic CRUD operations (Create, Read, Update and Delete) products, suppliers, etc. Manage automatically invoice and shipping operations
## How to use this repo?

### Do it by yourself
- Clone the repo
```
git clone https://github.com/AlexVonEinzbern/go-rest-api.git
```
- Add an `.env` file in the main root of the project containing the environment variables with the database credentials. E.g. :
```
DB_HOST=host.docker.internal
DB_USER=voneinzbern
DB_PASSWORD=bacbabaBDBWUOB3242obf
DB_NAME=go-rest-api
```
- Uncomment the followig lines on `db/connection.go` (Line 8):
```
//"github.com/joho/godotenv"
```
(Lines 15-19) 
```
//err := godotenv.Load(".env")

//if err != nil {
//	log.Fatal("Error loading .env file")
//}
```
- Uncomment the following line on `Dockerfile` (Line 27):
```
#COPY .env /
```

- Build the project using `Docker`:
```
docker build -t go-rest-api .
```
- Run the container:

&emsp; On `Mac`:
```
docker run -it -p 8080:8080 --env-file=".env" go-rest-api
```
&emsp; On `Linux` it is necessary to use the flag `--add-host=host.docker.internal:host-gateway`
```
docker run --add-host=host.docker.internal:host-gateway -it -p 8080:8080 --env-file=".env" go-rest-api
```
### Use the API
[https://go-rest-api.onrender.com](https://go-rest-api.onrender.com/api/v1)

## Documentation
After the project is running, an interactive documentation created using `Swagger` can be visited on [http://localhost:8080/swagger/index.html#/](http://localhost:8080/swagger/index.html#/)

**Note:** The online documentation can be found [here](https://go-rest-api.onrender.com/swagger/index.html#/).

&emsp; Some examples:

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/4fFx9ZP/swagger1.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 1. Documentation preview.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/V2GVYXH/swagger2.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 2. Example of GET request.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/yyKvCFP/swagger3.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 3. Example of POST request.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/cFwLXYn/swagger4.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 4. Example of PATH (UPDATE) request.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/QPJs8QP/swagger5.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 5. Example of GET request using date.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/z2b5tmh/swagger6.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 6. Example of DELETE request.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>

<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div class="swagger" align="center">
      <table>
        <tr>
         <img src="https://i.ibb.co/CwLRwP4/swagger7.png" aling="center"> 
        </tr>
        <tr>
          <p aling="center"><b>Fig 7. Example of model.</b></p>
        </tr>
      </table>
    </div>
  </body>
</html>
