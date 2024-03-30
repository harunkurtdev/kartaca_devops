#!/bin/bash

sleep 5
curl -XPUT 'http://localhost:9200/cities' -H 'Content-Type: application/json' -d '{
  "settings": {},
  "mappings": {
    "properties": {
      "il": {"type": "text"},
      "nufus": {"type": "integer"},
      "ilceler": {"type": "keyword"}
    }
  }
}'

for i in {1..10}; do
  curl -XPOST 'http://localhost:9200/cities/_doc' -H 'Content-Type: application/json' -d '{
    "il": "City'$i'",
    "nufus": '${i}',
    "ilceler": ["District'$i'_1", "District'$i'_2"] 
  }'
done

# curl -XPUT 'http://localhost:9200/countries' -H 'Content-Type: application/json' -d '{
#   "settings": {},
#   "mappings": {
#     "properties": {
#       "name": {"type": "text"}
#     }
#   }
# }'

# for i in {1..10}; do
#   curl -XPOST 'http://localhost:9200/countries/_doc' -H 'Content-Type: application/json' -d '{
#     "name": "Country'$i'"
#   }'
# done
