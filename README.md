# BOOKCASE API

This api was created in Go for a school project, in this app you can find book and add .
The api use JWT Token to login for create article and moderate comments.

## Summary

Libraries used for the API :
* [gorm](https://github.com/jinzhu/gorm)
* [elastic-search](https://github.com/dgrijalva/jwt-go)
* [mux](https://github.com/gorilla/mux)

## Use

Launch project with this command in the repo

```bash
go install -v ./...
```

## Get Started

* To start the API, use the **go run main.go** command.
* You can execute all requests in non-user-restricted endpoints without a token.
* To find a book, use the route /api/book.


# DOCUMENTATION

## Open Endpoints

# Books

Used to collect a Token for a registered User.

**URL** : `/api/book/`

**Method** : `GET`

**Auth required** : NO

**Data constraints**

```json
{
}
```
If you use postman, you must complete the request in json in body>raw


## Success Response

**Condition** : 
**Code** : `200 OK`

**Content example**

```json
{
   {"author":"Isabel Allende",
   "title":"La casa de los espíritus",
   "year":1982},
   {"author":"Victor Hugo",
   "title":"Les miserables",
   "year":1862}
}
```

