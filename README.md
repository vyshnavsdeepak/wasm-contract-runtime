# Contract Execution Runtime with Server

## Directory Structure

```
mkdir -p data/contracts
mkdir -p data/state
```

```
data
├── contracts
│   ├── contract1.wasm
│   └── contract2.wasm
├── state
│   ├── contract1.json
│   └── contract2.json
```

Move contracts to `data/contracts` directory.

## Run server
```
go run server.go
```

## Get state API call

```
curl --location --request GET 'http://localhost:8304/{contract_id}/state'
```
###Example
```
curl --location --request GET 'http://localhost:8304/123/state'
```

## Post events API call
```
curl --location --request POST 'http://localhost:8304/{contract_id}/events' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "blockID": 1,
        "function": "{functionName}",
        "args": [arg1, arg2, ...]
    }
]'
```

### Example
```
  curl --location --request POST 'http://localhost:8304/123/events' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "blockID": 1,
        "function": "vote",
        "args": ["blue"]
    }
]'
```