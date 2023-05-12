# Tugas REAST API Kelompok 7 Pemrograman Integrative

## Table of Contents
- [Tugas REAST API Kelompok 7 Pemrograman Integrative](#tugas-reast-api-kelompok-7-pemrograman-integrative)
  - [Table of Contents](#table-of-contents)
  - [Deskripsi Singkat](#deskripsi-singkat)
  - [Endpoint](#endpoint)
  - [Struktur File](#struktur-file)
  - [Requirements](#requirements)
  - [Cara Menjalankan Program](#cara-menjalankan-program)
  - [Anggota Kelompok](#anggota-kelompok)

## Deskripsi Singkat
Repository ini berisi implementasi REST API sederhana menggunakan bahasa pemrograman Go dan framework Gin. API ini dapat digunakan untuk mengelola data pengguna, permainan, dan kepemilikan permainan dalam operasi CRUD (*Create*, *Read*, *Upload*, *Delete*)

## Endpoint
1. Bagian Users terdiri dari 5 endpoint, yaitu:
- **GET** */rest/api/users* → Menampilkan daftar seluruh pengguna.
- **GET** */rest/api/users/:id* →Menampilkan data pengguna berdasarkan ID pengguna.
- **POST** */rest/api/user* → Membuat pengguna baru.
- **PUT** */rest/api/user/:id* → Memperbarui data pengguna berdasarkan ID pengguna.
- **DELETE** */rest/api/user/:id* → Menghapus data pengguna berdasarkan ID pengguna.
Selain itu, ada juga endpoint **GET** /rest/api/password/:id yang digunakan untuk membandingkan password pengguna berdasarkan ID pengguna.

2. Bagian Games terdiri dari 5 endpoint, yaitu:
- **GET** */rest/api/games* → Menampilkan daftar seluruh game.
- **GET** */rest/api/games/:id* → Menampilkan data game berdasarkan ID game.
- **POST** */rest/api/game* → Membuat game baru.
- **PUT** */rest/api/game/:id* → Memperbarui data game berdasarkan ID game.
- **DELETE** */rest/api/game/:id* → Menghapus data game berdasarkan ID game.

3. Bagian Kepemilikan terdiri dari 4 endpoint, yaitu:
- **GET** */rest/api/kepemilikan* → Menampilkan daftar seluruh kepemilikan game.
- **GET** */rest/api/kepemilikan/:id* → Menampilkan data kepemilikan game berdasarkan ID kepemilikan.
- **POST** */rest/api/kepemilikan:* → Membuat kepemilikan game baru.
- **DELETE** */rest/api/kepemilikan/:id* → Menghapus data kepemilikan game berdasarkan ID kepemilikan.


## Struktur File

```
tugas_rest_api
├─ bin
├─ pkg
└─ src
   ├─ controllers
   │  ├─ gamescontroller
   │  │  └─ gamescontroller.go
   │  ├─ kepemilikancontroller
   │  │  └─ kepemilikancontroller.go
   │  ├─ password
   │  │  └─ password.go
   │  └─ userscontroller
   │     └─ userscontroller.go
   ├─ go.mod
   ├─ go.sum
   ├─ main.go
   └─ models
      ├─ games
      │  └─ games.go
      ├─ kepemilikan
      │  └─ kepemilikan.go
      ├─ setup.go
      └─ users
         └─ users.go

```
 
## Requirements
1. Visual Studio Code
2. Bahasa Pemrograman Go
3. MySQL (`username : 'root', password : '', port : '3306'`)

## Cara Menjalankan Program
Langkah-langkah proses setup program adalah sebagai berikut:
1. Clone repository ini
2. Buat dahulu sebuah database pada MySQL bernama `steam`
2. Arahkan terminal pada `src` lalu jalankan perintah `go run main.go'
3. REST API siap digunakan pada `localhost:8080`

## Anggota Kelompok
<table>
    <tr>
        <td>No.</td>
        <td>Nama</td>
        <td>NIM</td>
    </tr>
    <tr>
        <td>1.</td>
        <td>Ichda Maulana Ramadhan Pelis</td>
        <td>201402134</td>
    </tr>
    <tr>
        <td>2.</td>
        <td>Alessandro Siburian</td>
        <td>201402095</td>
    </tr>
    <tr>
        <td>3.</td>
        <td>Rahmad Adli Harahap</td>
        <td>201402014</td>
    </tr>
    <tr>
        <td>4.</td>
        <td>Asad Iqbal</td>
        <td>201402149</td>
    </tr>
    <tr>
        <td>5.</td>
        <td>Anthony De Rivaldo</td>
        <td>201402119</td>
    </tr>
</table>
