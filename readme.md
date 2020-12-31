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
    - /mineservice/mine
    - /mineservice/mine/{key}
### `POST`
    - mineservice/mine
    - /mineservice/mine/sell
    - /mineservice/mine/buy
### `PUT`
    - /mineservice/mine
### `DELETE`
    - /mineservice/mine

## Routes cli
 ### `transaction commands`
    - create-mine [name] [price] [selling] [efficiency] [invetory] [uraniumCost] [resources]
    - sell-mine [id] [price]
    - set-mine [id]  [name] [price] [selling] [efficiency] [invetory] [uraniumCost] [resources]
    - delete-mine [id]
    - buy-mine [id] [price]
 ### `query commands`
    - list-mine
    - get-mine [key]
## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)