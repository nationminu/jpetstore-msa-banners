# jpetstore microservices banner API

## docker build
```
docker build -t petstore/banners .

docker run -p "8080:8080" petstore/banners
```

## curl test
```
curl http://localhost:8080/banners
[{"favcategory":"FISH","bannername":"\u003cimage src=\"../images/banner_fish.gif\"\u003e","descn":"Saltwater, Freshwater","image":"\u003cimage src=\"../images/fish_icon.gif\"\u003e"},{"favcategory":"CATS","bannername":"\u003cimage src=\"../images/banner_cats.gif\"\u003e","descn":"Various Breeds","image":"\u003cimage src=\"../images/dogs_icon.gif\"\u003e"},{"favcategory":"DOGS","bannername":"\u003cimage src=\"../images/banner_dogs.gif\"\u003e","descn":"Various Breeds, Exotic Varieties","image":"\u003cimage src=\"../images/reptiles_icon.gif\"\u003e"},{"favcategory":"REPTILES","bannername":"\u003cimage src=\"../images/banner_reptiles.gif\"\u003e","descn":"Lizards, Turtles, Snakes","image":"\u003cimage src=\"../images/cats_icon.gif\"\u003e"},{"favcategory":"BIRDS","bannername":"\u003cimage src=\"../images/banner_birds.gif\"\u003e","descn":"Exotic Varieties","image":"\u003cimage src=\"../images/birds_icon.gif\"\u003e"}]

curl localhost:8080/banners/FISH
{"favcategory":"FISH","bannername":"\u003cimage src=\"../images/banner_fish.gif\"\u003e","descn":"Saltwater, Freshwater","image":"\u003cimage src=\"../images/fish_icon.gif\"\u003e"}
``# jpetstore-msa-banners
