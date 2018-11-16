# WebBitcoinEscrow
Minimalist and secure web based escrow service for Bitcoin powered by blockchain.info written in Go.
  - Very secure due to secret based system rather than using authentication
  - Written in Go with Echo, thus highly performant and low in RAM usage
  - Easy to deploy
  - Adminpanel to deal with escalated escrows
  - Secured against Denial-of-service attacks
  - leverages blockchain.info API
  
### Installation
WebBitcoinEscrow requires [Go 1.9](https://golang.org/dl/)+ to run.

Install the dependencies and build with Go.

```sh
go get github.com/didip/tollbooth
go get github.com/didip/tollbooth_echo
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/sqlite
go get github.com/labstack/echo
go get github.com/rs/xid
go get golang.org/x/crypto/bcrypt
```
### Screenshots
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/1.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/2.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/3.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/4.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/5.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/6.jpg)
![pic1](https://raw.githubusercontent.com/tegk/WebBitcoinEscrow/master/7.jpg)
