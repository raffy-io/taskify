-- name: UpsertHomepage :exec
INSERT INTO homepage (id, hero_image, hero_image_alt, hero_title, hero_text)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id) 
DO UPDATE SET 
    hero_image = EXCLUDED.hero_image,
    hero_image_alt = EXCLUDED.hero_image_alt,
    hero_title = EXCLUDED.hero_title,
    hero_text = EXCLUDED.hero_text;