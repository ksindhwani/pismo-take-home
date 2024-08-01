# pismo-take-home
Take Home Assingnment


This Application uses Docker and GoModules 

# Pre-Requisities 


System must have Docker installed 


# Setup Instructions


1. git clone the repo / Download zip .
2. Unzip if downloaded as zip else skip to next step.
3. Rename `env.Template` to `.env`
4. Open the application directory in terminal 

5. Run the application through this command 

`docker-compose -f docker-compose.yml up -d`


# Endpoints

## Create Account

```
curl --location 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data '{
"document_number": "12345678902"
}'
```

## Create Transaction

```
curl --location 'http://localhost:8081/transactions' \
--header 'Content-Type: application/json' \
--data '{
"account_id": 4,
"operation_type_id": 3,
"amount": 123.45
}'
```

## Get Account

```
curl --location 'http://localhost:8081/accounts/12'
```


