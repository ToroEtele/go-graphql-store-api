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
