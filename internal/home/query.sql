-- name: UpsertHomepage :exec
INSERT INTO homepage (id, image)
VALUES ($1, $2)
ON CONFLICT (id) 
DO UPDATE SET 
    image = EXCLUDED.image;