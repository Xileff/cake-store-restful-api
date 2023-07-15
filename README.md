
# Cake Store RESTful API
This is a simple RESTful API for creating, reading, updating, and deleting cakes data. I built it using Golang and MySQL database. It will run on **port 5000**, so make sure the port is not in use by any other process.

## Installation
* Clone the project
    ```bash
        git clone https://github.com/Xileff/cake-store-restful-api.git
    ```

* Install dependencies
    ```bash
        go mod download
    ```
* Import the MySQL database in the *```db_export.sql```* file

* Create a *`.env`* file in this project's root directory. It must contain these values
    ```bash
        DB_HOST=<localhost or your_db_instance>
        DB_PORT=<3306 or your_db_port>
        DB_USER=<your_db_user>
        DB_PASS=<your_db_password>
    ```

* Run the project
    ```bash
        go run main.go
    ```

* or... run the unit tests
    ```
        go test .\test\cake_controller_test.go -v
    ```

## API Reference
### 1. List cakes
```http
    GET /cakes
```
#### Description
Get a list of all cakes in JSON format, sorted by rating (highest to lowest) and title alphabetically

### 2. Detail of cake
```http
    GET /cakes/{id}
```
#### Description
Get a cake data by the specified id

#### Query Parameter
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to fetch |

### 3. Add cake
```http
    POST /cakes
```
#### Description
Add a new cake to the app

#### Request Payload (JSON)
| Key | Type     | Value                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Title of the cake |
| `description`      | `string` | **Required**. Description of the cake |
| `rating`      | `float` | **Required**. Rating of the cake (1-10) |
| `image`      | `string` | **Required**. Image URL of the cake |

### 4. Update the cake
```http
    PUT /cakes/{id}
```
#### Description
Update a cake data by the specified id

#### Query Parameter
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to update |

#### Request Payload (JSON)
| Key | Type     | Value                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. Title of the cake |
| `description`      | `string` | **Required**. Description of the cake |
| `rating`      | `float` | **Required**. Rating of the cake (1-10) |
| `image`      | `string` | **Required**. Image URL of the cake |

### 5. Delete the cake
```http
    DELETE /cakes/{id}
```
#### Description
Delete a cake data by the specified id

#### Query Parameter
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to update |