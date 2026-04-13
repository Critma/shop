# SHOP API

## Docker run

1. rename **config/.env.example** to **config/.env**
2. run command (make):
```shell
make compose
```
or
```
docker-compose -p "shop-api" -f deploy/docker-compose.yml --env-file ./config/.env up -d
```


## Local run
1. rename **.env.example** to **.env**
1. run *postgres* service in compose file
    ```
    make compose-postgres
    ```
1. run command:
    ```
    make run
    ```