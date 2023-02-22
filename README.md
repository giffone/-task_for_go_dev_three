## HTTP server for proxying HTTP-requests to 3rd-party services

### Run

```bash
go run .
```

### endpoint /proxy (HTTP POST)
request
```json
curl --request POST \
  --url http://localhost:8000/proxy \
  --header 'Content-Type: application/json' \
  --data '{
    "method": "GET",
    "url": "http://google.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ="
    }
}
'
```
response
```json
{
	"id": "6b44ad3c-83c0-4831-8732-66f4bb9b3990",
	"status": "200 OK",
	"headers": {
		"Cache-Control": [
			"private, max-age=0"
		],
		"Content-Type": [
			"text/html; charset=ISO-8859-1"
		],
		"Date": [
			"Wed, 22 Feb 2023 21:37:42 GMT"
		],
		"Expires": [
			"-1"
		],
		"P3p": [
			"CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\""
		],
		"Server": [
			"gws"
		],
		"Set-Cookie": [
			"1P_JAR=2023-02-22-21; expires=Fri, 24-Mar-2023 21:37:42 GMT; path=/; domain=.google.com; Secure",
			"AEC=ARSKqsJl4MBL4AovwtcIPupbzrLnnN9Htz7kEq8PHY7_r9s5nfKj3g8ljg; expires=Mon, 21-Aug-2023 21:37:42 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax",
			"NID=511=VxeYeB0NTtKMFyxdZPzQQcv2jRrRPd1IZrYo5qwjV9n4Eqv94qjsYvJijuDVXnJR2NQA0vtBlP62cVNhskLWS-8hr73vwgii_d_3f3bA3fyNeIkjaYAGu4sQ9ryVT2_3oDHIo1XpYS9ERGHeN3uZpTxQZ_9M_tnKwdMjVH6fzKk; expires=Thu, 24-Aug-2023 21:37:42 GMT; path=/; domain=.google.com; HttpOnly"
		],
		"X-Frame-Options": [
			"SAMEORIGIN"
		],
		"X-Xss-Protection": [
			"0"
		]
	},
	"length": 0
}
```