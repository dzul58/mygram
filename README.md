# MyGram

## Deskripsi

MyGram adalah aplikasi yang memungkinkan pengguna untuk menyimpan foto dan membuat komentar untuk foto orang lain. Aplikasi ini dilengkapi dengan proses CRUD dengan table dan alur yang akan dijelaskan berikut ini.

## Teknologi yang Digunakan

1. Framework: Gin Gonic dan ORM Gorm
2. Library yang wajib digunakan:
   - github.com/dgrijalva/jwt-go
   - golang.org/x/crypto

## Struktur Database

Aplikasi ini memerlukan 4 tabel utama:

### 1. User

- id (Primary key)
- username (string)
- email (string)
- password (string)
- age (integer)
- created_at (Date)
- updated_at (Date)

### 2. Photo

- id (Primary key)
- title (string)
- caption (string)
- photo_url (string)
- user_id (Foreign Key Of User Table)
- created_at (Date)
- updated_at (Date)

### 3. Comment

- id (Primary Key)
- user_id (Foreign Key Of User Table)
- photo_id (Foreign Key Of Photo Table)
- message (string)
- created_at (Date)
- updated_at (Date)

### 4. SocialMedia

- id (Primary Key)
- name (String/varchar)
- social_media_url (String/Text)
- UserId (Foreign Key Of User Table)
- created_at (Date)
- updated_at (Date)

## Validasi

### A. Validasi untuk table User

1. Field email:

   - Validasi pengecekan format email yang valid
   - Validasi agar dapat menjadi unique index
   - Validasi agar field email tidak boleh kosong atau harus terisi

2. Field username:

   - Validasi agar dapat menjadi unique index
   - Validasi agar field username tidak boleh kosong atau harus terisi

3. Field password:

   - Validasi agar field password tidak boleh kosong
   - Validasi agar field password minimal memiliki panjang sebanyak 6 karakter

4. Field age:
   - Validasi agar field age tidak boleh kosong atau harus terisi
   - Validasi agar field age minimal memiliki nilai diatas 8

### B. Validasi Untuk Table Photo

1. Field title:

   - Validasi agar field title tidak boleh kosong

2. Field photo_url:
   - Validasi agar field photo_url tidak boleh kosong atau harus terisi

### C. Validasi Untuk Table SocialMedia

1. Field name:

   - Validasi agar field name tidak boleh kosong atau harus terisi

2. Field social_media_url:
   - Validasi agar field social_media_url tidak boleh kosong atau harus terisi

### D. Validasi untuk Table Comment

1. Field message:
   - Validasi agar field message tidak boleh kosong atau harus terisi

## Endpoints API

### Users

#### 1. Register User

```http
POST /users/register
```

Request Body:

```json
{
  "age": "integer",
  "email": "string",
  "password": "string",
  "username": "string"
}
```

Response (201):

```json
{
  "age": "integer",
  "email": "string",
  "id": "integer",
  "username": "string"
}
```

Note: Password user harus di hash menggunakan package Bcrypt sebelum di simpan ke database.

#### 2. Login User

```http
POST /users/login
```

Request Body:

```json
{
  "email": "string",
  "password": "string"
}
```

Response (200):

```json
{
  "token": "jwt string"
}
```

Note: Pada endpoint ini, wajib melakukan logika user login dengan pengecekan email dan password user. Pengecekan password wajib dilakukan dengan bantuan library/package Bcrypt.

#### 3. Update User

```http
PUT /users
```

Request:

- headers: Authorization (Bearer token string)
- params: userId (integer)
- body:

```json
{
  "email": "string",
  "username": "string"
}
```

Response (200):

```json
{
  "id": "integer",
  "email": "string",
  "username": "string",
  "age": "integer",
  "updated_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 4. Delete User

```http
DELETE /users
```

Request:

- headers: Authorization (Bearer token string)

Response (200):

```json
{
  "message": "Your account has been successfully deleted"
}
```

Note: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

### Photos

#### 1. Create Photo

```http
POST /photos
```

Request:

- headers: Authorization (Bearer token string)
- body:

```json
{
  "title": "string",
  "caption": "string",
  "photo_url": "string"
}
```

Response (201):

```json
{
  "id": "integer",
  "title": "string",
  "caption": "string",
  "photo_url": "string",
  "user_id": "integer",
  "created_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 2. Get Photos

```http
GET /photos
```

Request:

- headers: Authorization (Bearer token string)

Response (200):

```json
[
  {
    "id": "integer",
    "title": "string",
    "caption": "string",
    "photo_url": "string",
    "user_id": "integer",
    "created_at": "date",
    "updated_at": "date",
    "User": {
      "email": "string",
      "username": "string"
    }
  }
]
```

Note: Endpoint ini perlu melewati proses autentikasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 3. Update Photo

```http
PUT /photos/:photoId
```

Request:

- headers: Authorization (Bearer token string)
- params: photoId (integer)
- body:

```json
{
  "title": "string",
  "caption": "string",
  "photo_url": "string"
}
```

Response (200):

```json
{
  "id": "integer",
  "title": "string",
  "caption": "string",
  "photo_url": "string",
  "user_id": "integer",
  "updated_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data photo miliknya sendiri.

#### 4. Delete Photo

```http
DELETE /photos/:photoId
```

Request:

- headers: Authorization (Bearer token string)
- params: photoId (integer)

Response (200):

```json
{
  "message": "Your photo has been successfully deleted"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh menghapus data photo miliknya sendiri.

### Comments

#### 1. Create Comment

```http
POST /comments
```

Request:

- headers: Authorization (Bearer token string)
- body:

```json
{
  "message": "string",
  "photo_id": "integer"
}
```

Response (201):

```json
{
  "id": "integer",
  "message": "string",
  "photo_id": "integer",
  "user_id": "integer",
  "created_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 2. Get Comments

```http
GET /comments
```

Request:

- headers: Authorization (Bearer token string)

Response (200):

```json
[
  {
    "id": "integer",
    "message": "string",
    "photo_id": "integer",
    "user_id": "integer",
    "updated_at": "date",
    "created_at": "date",
    "User": {
      "id": "integer",
      "email": "string",
      "username": "string"
    },
    "Photo": {
      "id": "integer",
      "title": "string",
      "caption": "string",
      "photo_url": "string",
      "user_id": "integer"
    }
  }
]
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 3. Update Comment

```http
PUT /comments/:commentId
```

Request:

- headers: Authorization (Bearer token string)
- params: commentId (integer)
- body:

```json
{
  "message": "string"
}
```

Response (200):

```json
{
  "id": "integer",
  "title": "string",
  "caption": "string",
  "photo_url": "string",
  "user_id": "integer",
  "updated_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data comment miliknya sendiri.

#### 4. Delete Comment

```http
DELETE /comments/:commentId
```

Request:

- headers: Authorization (Bearer token string)
- params: commentId (integer)

Response (200):

```json
{
  "message": "Your comment has been successfully deleted"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh menghapus data comment miliknya sendiri.

### SocialMedias

#### 1. Create Social Media

```http
POST /socialmedias
```

Request:

- headers: Authorization (Bearer token string)
- body:

```json
{
  "name": "string",
  "social_media_url": "string"
}
```

Response (201):

```json
{
  "id": "integer",
  "name": "string",
  "social_media_url": "string",
  "user_id": "integer",
  "created_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 2. Get Social Medias

```http
GET /socialmedias
```

Request:

- headers: Authorization (Bearer token string)

Response (200):

```json
{
  "social_media": [
    {
      "id": "integer",
      "name": "string",
      "social_media_url": "string",
      "UserId": "integer",
      "createdAt": "date",
      "updatedAt": "date",
      "User": {
        "id": "integer",
        "username": "string",
        "profile_image_url": "string"
      }
    }
  ]
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken.

#### 3. Update Social Media

```http
PUT /socialmedias/:socialMediaId
```

Request:

- headers: Authorization (Bearer token string)
- params: socialMediaId (integer)
- body:

```json
{
  "name": "string",
  "social_media_url": "string"
}
```

Response (200):

```json
{
  "id": "integer",
  "name": "string",
  "social_media_url": "string",
  "user_id": "integer",
  "updated_at": "date"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh mengupdate data social media miliknya sendiri.

#### 4. Delete Social Media

```http
DELETE /socialmedias/:socialMediaId
```

Request:

- headers: Authorization (Bearer token string)
- params: socialMediaId (integer)

Response (200):

```json
{
  "message": "Your social media has been successfully deleted"
}
```

Note: Endpoint ini perlu melewati proses autentikasi dan autorisasi terlebih dahulu. Proses autentikasi wajib dilakukan dengan bantuan package/library JsonWebToken. Dan alur proses autorisasinya adalah user hanya boleh menghapus data social media miliknya sendiri.

## Catatan Penting

- Endpoint-endpoint untuk mengakses data pada tabel SocialMedia, Photo, dan Comment harus melalui proses autentikasi terlebih dahulu, dan proses autentikasinya wajib menggunakan JsonWebToken.
- Untuk endpoint-endpoint yang berguna untuk memodifikasi data kepemilikan seperti Update atau delete maka harus melalui proses autorisasi.

## Notes (Wajib)

- Seluruh routing endpoint diatas harus diikuti dengan betul
- Seluruh request body, headers maupun request params harus diikuti dengan betul.
- Seluruh response status, maupun response data nya harus diikuti dengan betul.
- Mohon untuk diperhatikan notes-notes yang telah diberikan diatas seperti endpoint-endpoint yang harus melewati proses autentikasi dan yang tidak. Begitu juga dengan proses autorisasinya.
- Perlu diingat disini bahwa proses autorisasi dilakukan setelah proses autentikasi, bukan sebaliknya.

## Struktur Aplikasi

```
├── cmd
│   └── main.go                 # Entry point aplikasi
├── config
│   └── database.go             # Konfigurasi database
├── controllers                 # Handler untuk setiap request
│   ├── user_controller.go
│   ├── photo_controller.go
│   ├── comment_controller.go
│   └── socialmedia_controller.go
├── middlewares                # Middleware untuk auth dan validasi
│   ├── authentication.go
│   └── authorization.go
├── models                     # Definisi struct dan validasi
│   ├── user.go
│   ├── photo.go
│   ├── comment.go
│   └── socialmedia.go
├── repositories              # Layer untuk interaksi dengan database
│   ├── user_repository.go
│   ├── photo_repository.go
│   ├── comment_repository.go
│   └── socialmedia_repository.go
├── services                  # Business logic
│   ├── user_service.go
│   ├── photo_service.go
│   ├── comment_service.go
│   └── socialmedia_service.go
├── routes                    # Definisi routing
│   └── router.go
├── helpers                   # Helper functions
│   ├── jwt.go
│   ├── bcrypt.go
│   └── response.go
├── dto                      # Data Transfer Objects
│   ├── user_dto.go
│   ├── photo_dto.go
│   ├── comment_dto.go
│   └── socialmedia_dto.go
├── go.mod
├── go.sum
└── README.md
```

### Penjelasan Struktur:

1. **cmd/**

   - Berisi file main.go sebagai entry point aplikasi
   - Inisialisasi database, router, dan menjalankan server

2. **config/**

   - Konfigurasi database dan environment variables
   - Setup koneksi database menggunakan GORM

3. **controllers/**

   - Menangani HTTP requests
   - Validasi input
   - Memanggil service layer untuk business logic
   - Mengembalikan response ke client

4. **middlewares/**

   - authentication.go: Validasi JWT token
   - authorization.go: Cek kepemilikan resource

5. **models/**

   - Definisi struct untuk entities
   - Validasi menggunakan struct tags
   - Relasi antar model

6. **repositories/**

   - Interface dan implementasi untuk akses database
   - Menggunakan GORM untuk operasi database
   - Implementasi CRUD operations

7. **services/**

   - Business logic aplikasi
   - Koordinasi antara repository dan controller
   - Implementasi validasi bisnis

8. **routes/**

   - Setup router menggunakan Gin
   - Grouping routes
   - Penerapan middleware

9. **helpers/**

   - jwt.go: Fungsi untuk generate dan validate JWT
   - bcrypt.go: Hash dan verify password
   - response.go: Format response API

10. **dto/**
    - Struct untuk request/response data
    - Validasi input
    - Transform data antara model dan response

### Contoh Implementasi Route:

```go
// routes/router.go
func SetupRouter() *gin.Engine {
    router := gin.Default()

    // User routes
    userGroup := router.Group("/users")
    {
        userGroup.POST("/register", userController.Register)
        userGroup.POST("/login", userController.Login)
        userGroup.PUT("/", middleware.Authentication(), userController.Update)
        userGroup.DELETE("/", middleware.Authentication(), userController.Delete)
    }

    // Photo routes
    photoGroup := router.Group("/photos")
    photoGroup.Use(middleware.Authentication())
    {
        photoGroup.POST("/", photoController.Create)
        photoGroup.GET("/", photoController.GetAll)
        photoGroup.PUT("/:photoId", middleware.PhotoAuthorization(), photoController.Update)
        photoGroup.DELETE("/:photoId", middleware.PhotoAuthorization(), photoController.Delete)
    }

    // Comment routes
    commentGroup := router.Group("/comments")
    commentGroup.Use(middleware.Authentication())
    {
        commentGroup.POST("/", commentController.Create)
        commentGroup.GET("/", commentController.GetAll)
        commentGroup.PUT("/:commentId", middleware.CommentAuthorization(), commentController.Update)
        commentGroup.DELETE("/:commentId", middleware.CommentAuthorization(), commentController.Delete)
    }

    // Social Media routes
    socialMediaGroup := router.Group("/socialmedias")
    socialMediaGroup.Use(middleware.Authentication())
    {
        socialMediaGroup.POST("/", socialMediaController.Create)
        socialMediaGroup.GET("/", socialMediaController.GetAll)
        socialMediaGroup.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.Update)
        socialMediaGroup.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), socialMediaController.Delete)
    }

    return router
}
```

Struktur ini mengikuti prinsip:

- Separation of Concerns
- Clean Architecture
- Dependency Injection
- Repository Pattern
- Middleware Pattern

Setiap layer memiliki tanggung jawab yang jelas dan terpisah, memudahkan untuk testing dan maintenance.
