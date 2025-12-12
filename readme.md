API EndPoints :

POST /auth/register - Register new user

POST /auth/login - Login user, get JWT token

GET /shops/me - Get all user's shops

POST /shops - Create new shop
Template Body (JSON):
{
"name": "User1' shops",
"address": "Rangkah"
}

PUT /shops/:id - Update shop (ownership verified)
{
    "name": "User1' shops",
    "address": "Rangkah"
}

DELETE /shops/:id - Delete shop (ownership verified)

GET /products/:id - Get product by ID (public)

POST /products - Create product (shop ownership verified):

PUT /product/:id - Update product (shop ownership verified)

Delete /product/:id - Delete product (shop ownership verified)
