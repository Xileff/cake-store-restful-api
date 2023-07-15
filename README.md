
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

* Create a *`.env`* file in this project's root directory (or just modify the one provided in this repo). It must contain these values
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

#### Response example
```json
{
    "status": "success",
    "data": [
        {
            "id": 2,
            "title": "Lapis Legit Update",
            "description": "Kue lapis is an Indonesian kue, or a traditional snack of steamed colourful layered soft rice flour pudding. In Indonesian lapis means \"layers\". This steamed layered sticky rice cake or pudding is quite popular in Indonesia, Suriname and can also be found in the Netherlands through their colonial links.",
            "rating": 7.5,
            "image": "https://endeus.tv/resep/kue-lapis-legit",
            "created_at": "2023-07-15 11:14:29",
            "updated_at": "2023-07-15 11:14:42"
        },
        {
            "id": 1,
            "title": "Lemon cheesecake",
            "description": "A cheesecake made of lemon",
            "rating": 7.0,
            "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
            "created_at": "2023-07-15 11:14:18",
            "updated_at": "2023-07-15 11:14:18"
        }
    ]
}
```

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

#### Response example
```json
{
    "status": "success",
    "data": {
        "id": 1,
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7.0,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2023-07-15 11:14:18",
        "updated_at": "2023-07-15 11:14:18"
    }
}
```

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

#### Response example
```json
{
    "status": "success",
    "data": {
        "id": 2,
        "title": "Kue Lapis",
        "description": "Kue lapis is an Indonesian kue, or a traditional snack of steamed colourful layered soft rice flour pudding. In Indonesian lapis means \"layers\". This steamed layered sticky rice cake or pudding is quite popular in Indonesia, Suriname and can also be found in the Netherlands through their colonial links.",
        "rating": 9.0,
        "image": "https://endeus.tv/resep/kue-lapis-legit",
        "created_at": "2023-07-15 11:14:29",
        "updated_at": "2023-07-15 11:14:29"
    }
}
```

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

#### Response example
```json
{
    "status": "success",
    "data": {
        "id": 2,
        "title": "Lapis Legit Update",
        "description": "Kue lapis is an Indonesian kue, or a traditional snack of steamed colourful layered soft rice flour pudding. In Indonesian lapis means \"layers\". This steamed layered sticky rice cake or pudding is quite popular in Indonesia, Suriname and can also be found in the Netherlands through their colonial links.",
        "rating": 7.5,
        "image": "https://endeus.tv/resep/kue-lapis-legit",
        "created_at": "2023-07-15 11:14:29",
        "updated_at": "2023-07-15 11:14:42"
    }
}
```

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

#### Response example
```json
{
    "status": "success",
    "message": "Cake 2 deleted successfully."
}
```