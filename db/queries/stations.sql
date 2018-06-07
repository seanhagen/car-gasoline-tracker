-- name: find-station-by-latlng
select
       id,
       name,
       address,
       location,
       visits,
       ST_DistanceSphere(location, ST_GeomFromText('POINT(:x :y)', 4326)) distance
from stations
     where ST_DistanceSphere(location, ST_GeomFromText('POINT(:x :y)', 4326)) < 70
     order by ST_Distance(location, ST_GeomFromText('POINT(:x :y)', 4326))
     limit 1;

-- name: create-station
insert into stations (id, name, address, location, visits)
values (:id, :name, :address, :location, :visits)
