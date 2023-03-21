# production-ready-api


gh repo create manoj-negi/production-ready-api --public --add-readme --clone

git init -b main

migrate create -ext sql -dir db/migrations -seq create_users_table

docker exec -it bef9 bash