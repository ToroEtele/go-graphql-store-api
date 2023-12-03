-- name: GetAllProducts :many
SELECT * FROM product;

-- name: GetProductById :one
SELECT * FROM product WHERE id = ?;

-- name: GetProductByName :one
SELECT * FROM product WHERE name = ?; 

-- name: CreateProduct :exec
INSERT INTO product (id, name, price, description, image, stock, updatedAt) 
VALUES (UUID(), ?, ?, ?, ?, ?, NOW());

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = ?;

-- name: UpdateProduct :exec
UPDATE product 
SET name = COALESCE(sqlc.narg(name), name),
    price = COALESCE(sqlc.narg(price), price),
    description = COALESCE(sqlc.narg(description), description),
    image = COALESCE(sqlc.narg(image), image),
    stock = COALESCE(sqlc.narg(stock), stock),
    updatedAt = NOW()
WHERE id = ?;