Simple Blockchain
===================
## Description

Simple Blockchain technology to keep record of heart rate 

## Quick Start

- `git clone https://github.com/jakhax/learning-golang.git`
- `cd simple_blockchain`
- `touch .env && echo "PORT=8080" > .env`
- `go run main.go http_server.go` or   `go build`

## Endpoints
- You can use a browser, curl or [Postman](https://www.getpostman.com/apps) to test the endpoints

### View Entire Chain
``` bash
GET get-blockchain
# Request sample
curl --request GET http://localhost:8080/get-blockchain
```

### Create a Block
``` bash
POST write-block
# Request sample
curl --header "Content-Type: application/json" --request POST\
 --data '{"BPM":78}' http://localhost:8080/write-block
```

## Acknowlegdement / Resources
* [Coral Health](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)
* [Golang](https://golang.org)


## License ([MIT License](http://choosealicense.com/licenses/mit/))
This project is licensed under the MIT Open Source License.