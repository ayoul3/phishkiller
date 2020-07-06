# PhishKiller

PhishKiller is a tool to spam phishing websites with fake credentials drowning real credentials in a sea of uselessness.

Each request is sent using a User-Agent with slightly different version numbers.

Ideally, you would want to run it every half hour on a lambda function for 10 minutes to have that automatic IP rotation making it even harder to weed out your fake results.

The tool supports common fakers, courtesy of [syreclabs.com/go/faker](https://godoc.org/syreclabs.com/go/faker) Go package. Feel free to add more depending on your use case.


## Usage
* [Install Go](https://golang.org/dl/) in your environment.
* Clone the package  
```git clone https://github.com/ayoul3/phishkiller```
* Alter yml config file (see below)
* Run   
```go run main.go```

## Config
```
# 0: warning, 1: info, 2: debug
LogLevel: 0

# Optional if you wish to tunnel your traffic through an HTTP proxy
Proxy: http://127.0.0.1:8080

# Number of workers in parallel to send play your requests.
Workers: 1

# These requests will play in order, no matter the number of workers
Requests:

  # Full path of the target
  - Path: 'https://www.test.com/test'
    
    # Only supports get and post
    Method: 'get'
    
    # Optional: list of headers to add to the default ones (see below). All lowercase
    Headers:
      referer: "www.test.com"
      x-custom-header: "value"
          
    # Optional: List of parameters to include
    Params:
      # Type holds the type of faker data to generate. This will send: param1=<random_ip>
      - Name: param1
        Type: ipv4
        
      # This will send param2=fixed
      - Name: param2
        Value: fixed
```

By default POST requests send form-encoded parameters:
```
param1=1211122112342201&email=alexis.bunny%40gmail.com
```
JSON payload is automatically detected if you addd the following header: `content-type: application/json`
```
Headers:
  content-type: application/json
Params:
  - Name: param1
    Type: email
    
```
The body sent will be:
```
{"param1": "crudge_magdak@yahoo.com"}
```
## Default headers
```
User-Agent: Mozilla/5.0 (Windows NT 8.0; Win64; x64; rv:<random>) Gecko/20100<random> Firefox/<random>
Accept: text/html,application/xhtml+xml,application/xml;q=0.6,image/webp,*/*;q=0.5
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.5
Connection: close
Dnt: 1
Upgrade-Insecure-Requests: 1
```

## To do
* add support for PUT requests
* improve the CC faker
* support more fakers
* reuse fake value across multiple requests