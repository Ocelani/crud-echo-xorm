#!/bin/bash


#* User *#

# GET /users

# GET /users/:id

# POST /users
# curl -X POST -d 'username=celani' -d 'password=pass' localhost:1323/users

# PUT /users/:id

# DELETE /users/:id


#* Auth *#

# GET /accessible
curl -X GET localhost:1323/

# POST /login
curl -X POST -d 'username=rob' -d 'password=pike' localhost:1323/login

# GET /restricted
