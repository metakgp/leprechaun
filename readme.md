# Leprechaun

[![Go Report Card](https://goreportcard.com/badge/github.com/icyflame/leprechaun)](https://goreportcard.com/report/github.com/icyflame/leprechaun)
> Connect an email ID to your IIT KGP roll number!

## TOC
- [Public API](#public-api)
- [Development](#development)
- [Origin](#origin)
- [Etymology](#etymology)
- [TODO](#todo)
- [License](#license)

## Public API

**Note:** The public API can be accessed only by a few applications right now.
If you would like to use the API, contact one of the
[maintainers](https://wiki.metakgp.org/w/Metakgp:Governance#Current_maintainers)
on [Metakgp Slack](https://metakgp.slack.com).

**Authentication:** In your request, add an `Authorization` header. The
value of the header should be `Bearer <TOKEN>`. You can get this token from one
of the maintainers.

**Endpoint:** `GET /get/{input_type}/{input}`

Valid values of `input_type`:

1. `roll`
1. `email`

Examples:

1. `GET /get/roll/12CS40067`
1. `GET /get/email/bob@example.com`

**Response:** 

- 401 if request is unauthorized
- 404 if no record with the given input_type and input could be find
- 200 if a record was found. The response will contain a JSON object with the following keys set:

	1. `roll` - Roll number of the user
	1. `email` - Email of the user
	1. `authenticated` - Timestamp that the authentication was completed in [RFC3339][1] format

**Curl Verbose:**

```sh
$ curl -vvv -H "Authorization: Bearer abcd" https://leprechaun.metakgp.org/get/roll/12CS40067
*   Trying 172.16.2.30...
* Connected to 172.16.2.30 (172.16.2.30) port 8080 (#0)
* Establish HTTP proxy tunnel to leprechaun.metakgp.org:443
> CONNECT leprechaun.metakgp.org:443 HTTP/1.1
> Host: leprechaun.metakgp.org:443
> User-Agent: curl/7.47.0
> Proxy-Connection: Keep-Alive
>
< HTTP/1.1 200 Connection Established
< Proxy-Agent: IWSS
< Date: Wed, 14 Feb 2018 09:51:48 GMT
<
* Proxy replied OK to CONNECT request
* found 148 certificates in /etc/ssl/certs/ca-certificates.crt
* found 597 certificates in /etc/ssl/certs
* ALPN, offering http/1.1
* SSL connection using TLS1.2 / ECDHE_ECDSA_AES_128_GCM_SHA256
* ...
* ALPN, server accepted to use http/1.1
> GET /get/roll/12CS40067 HTTP/1.1
> Host: leprechaun.metakgp.org
> User-Agent: curl/7.47.0
> Accept: */*
> Authorization: Bearer abcd
>
< HTTP/1.1 200 OK
< Date: Wed, 14 Feb 2018 09:51:43 GMT
< Content-Type: application/json; charset=utf-8
< Content-Length: 59
< Connection: keep-alive
< Via: 1.1 vegur
< Server: cloudflare
<
* Connection #0 to host 172.16.2.30 left intact
{"authenticated":"2018-01-04T13:12:50Z","email":"bob@example.com","roll":"12CS40067"}
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

To push this to Heroku, run the build script. Note that this is run inside
Heroku under a GoLang buildpack. But before `go build`, we need to run `gulp`
and create all the HTML and Template assets. The build script takes care of
this. (It switches to a different branch, deletes `dist/` from `gitignore`, runs
`gulp`, adds everything, commits, pushes to Heroku and then, switches back to
`master`)

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
   user is given a unique verification key (say `VERIFIER_TOKEN`) and a unique
   link to visit.
2. The user is shown a verification key which they must add to one of their
   secret question text in ERP.

   **Eg:**  
   Old Question: _What was the name of your first pet?_  
   New Question: _What was the name of your first pet? (VERIFIER_TOKEN)_ 

   The question can be changed in the ERP using the Forgot password section on the
   IIT KGP ERP.

   Once they have changed the text of the question, they can visit the link
   provided in step 1.
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

<img src="public/leprechaun.png" height="600" />

Leprechaun

_n._ a mischievous elf of Irish folklore usually believed to reveal the
hiding place of treasure if caught

## TODO

Search for the strings `TODO` and `ENHANCE` in the codebase.

## License

Code licensed under MIT.

Metakgp 2018

[1]: https://tools.ietf.org/html/rfc3339
