# Tiny Coin
This is a small implementation of block chain technology in Go based loosely on the following article:

https://medium.com/crypto-currently/lets-build-the-tiniest-blockchain-e70965a248b

## Building
`go build`

`./tinycoin`

### Create a transaction
You can delete a book using an API call:
`POST /txion`
#### Example
```bash
curl -i -H "Content-Type: application/json" -X POST -d '{
  "from": "71238uqirbfh894-random-public-key-a-alkjdflakjfewn204ij",
  "to": "93j4ivnqiopvh43-random-public-key-b-qjrgvnoeirbnferinfo",
  "amount": 3
}' http://localhost:80/txion
```
#### Returns:
A saved transaction...
```javascript
{
  "from": "71238uqirbfh894-random-public-key-a-alkjdflakjfewn204ij",
  "to": "93j4ivnqiopvh43-random-public-key-b-qjrgvnoeirbnferinfo",
  "amount": 3
}
```

### Mine
You can mine a new coin using an API call:
`GET /mine/{minerAddress}`
#### Example:
```bash
curl -i -H "Content-Type: application/json" -X GET http://localhost:80/mine/93j4ivnqiopvh43-random-public-key-b-qjrgvnoeirbnferinfo
```
#### Returns:
The mined coin data
```javascript
{
    "Index": 1,
    "TimeStamp": "2018-04-15T21:07:09.286704-06:00",
    "Data": {
        "ProofOfWork": 9,
        "Transactions": [
            {
                "to": "93j4ivnqiopvh43-random-public-key-b-qjrgvnoeirbnferinfo",
                "from": "Network",
                "amount": 1
            }
        ]
    },
    "PreviousHash": [
        35,
        114,
        170,
        138,
        241,
        128,
        235,
        159,
        13,
        64,
        196,
        22,
        234,
        66,
        78,
        15,
        160,
        64,
        252,
        46,
        114,
        232,
        50,
        191,
        52,
        15,
        95,
        65,
        193,
        75,
        197,
        39
    ],
    "Hash": [
        176,
        109,
        197,
        236,
        187,
        15,
        54,
        149,
        110,
        219,
        14,
        41,
        4,
        131,
        207,
        108,
        36,
        45,
        186,
        15,
        61,
        246,
        0,
        241,
        100,
        0,
        151,
        167,
        53,
        204,
        162,
        183
    ]
}
```