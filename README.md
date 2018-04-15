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
