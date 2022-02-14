#### Init some Characters

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "John",
    "class_name": "Warrior"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "Mary",
    "class_name": "Mage"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "Oswaldu",
    "class_name": "Thief"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "Antonieu",
    "class_name": "Warrior"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "Arboratoire",
    "class_name": "Mage"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "mrs_stronger",
    "class_name": "Thief"
}'

curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "steve_rogers",
    "class_name": "Warrior"
}'