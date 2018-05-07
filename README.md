# Ethpay
This is an application with REST API for making ether payments to an Ethereum address.

## Instruction to run the application

### Dependencies

1. Postgres database: The postgres schema is available in the file `schema.sql`
2. Ethereum full node with account unlocked.

The server can be build and started with the following command:

```
go run main.go http://127.0.0.1:8443 "$POSGRES_CONNECTION_STRING" "$ACCOUNT_NAME"
```

## Production Checklist

1. Tests!!!
2. Better error handling
3. ~~Refactor the RPC stuff~~
4. Gas pricing calculation
5. Better Geth Integeration
6. Better Vendoring of Dependencies
7. Handle insufficient funds condition
8. Handle nonce generation

