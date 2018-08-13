Simple Golang Rest Api
===================
## Description

Simple RESTful API to create, read, update and delete books. 

No database implementation yet, using hard coded data.


## Quick Start
Make sure you have correctly setup go and created a project in `$GOPATH`
``` bash
# Install mux router
go get -u github.com/gorilla/mux
# compile
go build
./simple_rest_api
```

## API Endpoints

### Get All Books
``` bash
GET api/books
```
### Get Book by Id
``` bash
GET api/books/{id}
```

### Delete Book by Id
``` bash
DELETE api/books/{id}
```

### Create Book
``` bash
POST api/books

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Book Three",
#   "author":{"firstname":"Jane",  "lastname":"Doe"}
# }
```

### Update Book by Id
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"4545454",
#   "title":"Updated Title",
#   "author":{"firstname":"John",  "lastname":"Doe"}
# }

```

## Acknowlegdement
* [Golang](https://golang.org)
* [Brad Traversy](https://www.youtube.com/watch?v=SonwZ6MF5BE&t=2045s)
* [Mux Router](https://github.com/gorilla/mux)


## License ([MIT License](http://choosealicense.com/licenses/mit/))
This project is licensed under the MIT Open Source License.