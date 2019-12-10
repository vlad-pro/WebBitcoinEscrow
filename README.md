# WebBitcoinEscrow ![logo](https://en.bitcoin.it/w/images/en/f/fd/BC_nBG_64px.png)
Minimalist and secure web-based escrow service for Bitcoin powered by blockchain.info written in Go.
Very secure due to secret based system rather than using authentication
- Written in Go with Echo, thus highly performant and low in RAM usage
-	Easy to deploy
-	Admin Panel to deal with escalated escrows
-	Secured against Denial-of-service attacks
-	Leverages blockchain.info API

Example:

Alice wants to buy 2.6 Bitcoin from Bob for $5k but wants to be sure that she gets the Bitcoin. Bob also wants to be sure that Alice just receives the Bitcoin when money appears on his bank account. Alice generates an escrow with her details eg BTC Wallet. Sends a link that was generated to Bob. Bob uses the onetime secret-based URL to generate another secret based URL to release the trade once he sees the money on his bank statement. It’s not possible to generate the secrets again! There is a fee of 5% of the transaction value but this can be changed in the admin panel. If there is a conflict the user can escalate the escrow and the admin can resolve this in the panel.



### Installation
WebBitcoinEscrow requires [Go 1.9](https://golang.org/dl/)+ to run.

The Blockchain Wallet API V2 has to be installed and must be running.
https://github.com/blockchain/service-my-wallet-v3#installation
```sh
npm install -g blockchain-wallet-service
blockchain-wallet-service start --port 3000
```

Install the dependencies and build with "Go build main.go".

```sh
go get -u github.com/tegk/WebBitcoinEscrow

You might have to install them:

go get github.com/didip/tollbooth
go get github.com/didip/tollbooth_echo
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
go get github.com/labstack/echo
go get github.com/rs/xid
go get golang.org/x/crypto/bcrypt
go get -u github.com/tidwall/gjson
```
Run with admin rights as this binds port 80:
```sh
sudo ./main
```
To configure the application login into the admin panel at http://localhost/admin with user and password "Admin" and change the settings.

### Screenshots
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/1.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/2.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/3.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/4.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/5.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/6.jpg)
![pic1](https://github.com/tegk/WebBitcoinEscrow/blob/master/screenshots/7.jpg)
