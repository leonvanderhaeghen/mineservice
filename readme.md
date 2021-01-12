# mineservice

**mineservice** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started
[golang > 1.13.0](https://golang.org/doc/install) needed

```
npm i -g @tendermint/starport
```
`npm` command installs tendermint starport

- [starport install documentation](https://github.com/tendermint/starport/blob/develop/docs/1%20Introduction/2%20Install.md)
```
cd nameservice
```
`cd` cd to the folder you want to serve
```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

starport `launchpad` sdk version used to develop

Developed using `gitpod` IDE on github

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Routes api

### `GET`
    - /mineservice/player
    - /mineservice/player/{key}
    - /mineservice/resource
    - /mineservice/resource/{key}
    - /mineservice/mine
    - /mineservice/mine/{key}
### `CREAT PLAYER` (post)
    - mineservice/player
    {
         "base_req" : "sign with js library"
         "name": " "
    }
###  `CREAT RESOURCE` (post)
    - mineservice/resource
     {
         "base_req" : "sign with js library"
         "name": " "
         "amount":" "
         "mineID":" "
    }
### `MOVE RESOURCE` (post)
      - mineservice/resource/move
        {
            "base_req" : "sign with js library"
            "id" : " "
            "amount" : " "
            "playerid" : " "
         }
### `CREAT MINE` (post)
             - mineservice/mine
            {
                "base_req" : "sign with js library"
              	"owner" : " "
              	"name" : " "
              	"price" : "200token"
              	"selling" : "true"
              	"efficiency" : " "
              	"resources" : ["bronze"}
              	"uraniumCost" : " "
            }
### `SELL MINE` (post)
            - mineservice/mine/sell
            {
                "base_req" : "sign with js library"
                "id" : " "
                "price" : "300token"
            }
### `BUY MINE` (post)
        - mineservice/mine/buy
            {
                "base_req" : "sign with js library"
                "buyerid": " "
                "sellerid": " "

                "price": "300token"

            }
### `DELETE MINE` (delete)
            - mineservice/mine
            {
                "base_req" : "sign with js library"
                "id" : " "
            }
### `DELETE RESOURCE`(delete)
            - mineservice/resource
            {
                "base_req" : "sign with js library"
                "id" : " "
            }
### `DELETE PLAYER`(delete)
             - mineservice/player
             {
                 "base_req" : "sign with js library"
                 "id" : " "
             }                      
## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)