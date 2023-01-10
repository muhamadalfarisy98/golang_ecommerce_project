# Final Project
Final project: **Golang Backend Development** :wave:

## Informasi Project
- Tema : E-commerce
- Judul : E-commerce database system management 
- Creator : Muhamad Alfarisy (muhamadalfarisy9d@yahoo.com)

## Description
Berisikan implementasi aplikasi _**Backend**_ pada bahasa pemrograman Golang dan juga _**microservice**_ dalam bentuk REST API. Projek ini sendiri berupa database management sistem untuk jual beli online toko pakaian. 

## 1. Link deploy/hosting Heroku (deprecated)
> https://eclothestore-faris.herokuapp.com/swagger/index.html#/

## 2. register, login, change password
- silahkan buka dokumentasi pada point 1 untuk melihat enpoint fitur ini disertai dengan keamaan JWT Auth

## 3. Interaksi User sesuai dengan rolenya
- pada project ini, table users memiliki field 'role' yang bertipe boolean. value 1 berarti role sebagai admin,
dan value 0 sebagai user biasa (customer).
- interaksi user biasa atau customer, dapat memberikan review product yang telah dipesan.
- user juga dapat berinteraksi untuk memperbaharui data dirinya sendiri untuk keperluan informasi transaksi

## 4. Project ini memiliki sekitar 8 Tabel yang digunakan
- untuk keterangan tabel yang digunakan dapat melihat di dokumentasi API Heroku pada point 1

## 5. Dokumentasi API 
- dapat dilihat lebih detail pada link point 1

## ERD Sistem (Skema Database)
![erd](https://user-images.githubusercontent.com/23287190/147418329-e3797764-acd1-4f21-890c-7fbba364fdf2.png)

## File dump sql
> cd /ecommerce/dump_sql


