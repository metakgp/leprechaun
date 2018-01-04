# Leprechaun

> Connect an email ID to your IIT KGP roll number!

## TOC
- [Public API](#public-api)
    - [GET Endpoint](#get-endpoint)
    - [POST Endpoint](#post-endpoint)
- [Development](#development)
- [Origin](#origin)
- [Etymology](#etymology)
- [License](#license)

## Public API

**Note:** The public API can be accessed only by a few applications right now.
If you would like to use the API, contact one of the
[maintainers](https://wiki.metakgp.org/w/Metakgp:Governance#Current_maintainers)
on [Metakgp Slack](https://metakgp.slack.com).

### GET endpoint

**Endpoint:** GET /get/{rollNumber}

**Response:** 

`application/json` with the key `email` if this Roll number is authenticated

401 if you don't send the correct bearer token

404 if the roll number is not authenticated

**Curl Verbose:**

```sh
$ curl -vvv -H "Authorization: Bearer abcdefghijklmnopqrstuvwxyz" http://localhost:8080/get/12AB3456789
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /get/12AB34567 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization: Bearer abcdefghijklmnopqrstuvwxyz
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 04 Jan 2018 10:08:58 GMT
< Content-Length: 41
< 
* Connection #0 to host localhost left intact
{"email": "foobar@example.com"}
```

### POST endpoint

**Endpoint:** POST /get with the header `Content-Type: application/x-www-form-urlencoded`

**Response:** 

`application/json` with the key `email` if this Roll number is authenticated

401 if you don't send the correct bearer token

404 if the roll number is not authenticated

**Curl Verbose:**

```sh
$ curl -d 'roll=12AB3456789' -vvv -H "Authorization: Bearer abcdefghijklmnopqrstuvwxyz" http://localhost:8080/get
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /get HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization: Bearer abcdefghijklmnopqrstuvwxyz
> Content-Length: 14
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 14 out of 14 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 04 Jan 2018 10:11:26 GMT
< Content-Length: 41
< 
* Connection #0 to host localhost left intact
{"email": "foobar@example.com"}
```

## Development

To run this locally:

```sh
$ go get github.com/icyflame/leprechaun
$ cd $GOPATH/github.com/icyflame/leprechaun
$ npm install -g gulp-cli
$ npm install
$ gulp
$ go build
$ cp .env.template .env
# populate .env with the appropriate values
$ PORT=8080 ./leprechaun
```

You need to have MongoDB, Node.js, npm, Golang and gulp installed.

To populate the `.env` file, you need an active Sendgrid account.

## Origin

[Vikrant Varma](https://github.com/amrav) posted the [Idea
document](https://paper.dropbox.com/doc/Leprechaun-BK0eQTGGvMLbVoor4L0dJ) for
Leprechaun to the Metakgp Slack on 3rd January, 2018.

Leprechaun was envisioned as an authentication service which will associate an
email address to an IIT Kharagpur Roll Number, using proof-of-control over the
ERP account. Since only students of IIT Kharagpur have active ERP IIT KGP
accounts, this service can only be used by students of IIT Kharagpur.

This service ensures that students don't have to share their ERP credentials
with any application where they would like to authenticate themselves as
students of IIT Kharagpur.

The authentication flow goes like this:

1. User visits Leprechaun and asserts their roll number and email address. The
   user is given a unique verification key and a unique link to visit.
2. The user is shown a verification key which they must add to one of their
   secret question text. When they have completed this addition, they can visit
   the link provided to them in step 1.
3. Leprechaun will make requests to ERP IIT KGP to check whether one of the
   secret questions tied to the user's roll number has the verification key in
   it's text.
4. If step 3 is not successful, it tells the user to check their association.
5. If step 3 is successful, **Authentication Step 1** is now complete. The service
   has verified that the user has control over the roll number's ERP account.
   Now, a verification email is sent to the email address entered in step 1.
6. The user visits their inbox and clicks on the link in this verification
   email.
7. This leads them back to the service which will verify the token embedded in
   this verification link. If verified succesfully, **Authentication Step 2** is
   complete! _The roll number and email supplied by the user in step 1 are now
   associated in Leprechaun's DB_

Services (like [Metakgp Dashboard](https://github.com/metakgp/dashboard-beta))
can use Leprechaun using the [Public API](#public-api) to find the email address
that a supplied roll number is associated with. This email address can be used
as a proxy for the user's roll number.

## Etymology

![img](public/leprechaun.png)

Leprechaun

_n._ a mischievous elf of Irish folklore usually believed to reveal the
hiding place of treasure if caught

## License

Code licensed under MIT.

Metakgp 2017
