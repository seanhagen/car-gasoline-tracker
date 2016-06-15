-- name: fetch-location
SELECT id,name,longitude,latitude,address,visits FROM locations WHERE address LIKE $1
