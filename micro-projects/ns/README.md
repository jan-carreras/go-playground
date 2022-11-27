Networking Stack
================

This µ-project is about implementing an HTTP call to [example.org](https://example.org/) inspired by
the [The HTTP crash course nobody asked for](https://fasterthanli.me/articles/the-http-crash-course-nobody-asked-for) article.

The idea is to implement an HTTP/1.1 client and going down the network stack, using fewer and fewer libraries.

The purpose of the exercise is to learn new low-level stuff.

# Iterations

## Iteration 1

Build an HTTP Client using a common HTTP library. Playing the "GitHub popularity contest" rule of thumb,
I've seen [go-resty/resty](https://github.com/go-resty/resty) with 7k stars. Let's start with this one.

### Learnings

#### Bazel

[Bazel](https://bazel.build/):  `From startup to enterprise, choose the Bazel open source project to build and test your multi-language, multi-platform projects.`


#### Check your assumptions

The [library's implementation of `String()` is](https://github.com/go-resty/resty/blob/313f4190d9b8f7605ff298b769de613621c50220/response.go#L42):

```go
// String method returns the body of the server response as String.
func (r *Response) String() string {
if r.body == nil {
return ""
}
return strings.TrimSpace(string(r.body))
}
```

And I though: `wow! Are we not creating a new string and using the Bytes.TrimSpace be more efficient?`.

The answer is a clear: **NO**:

```
goos: darwin
goarch: amd64
pkg: github.com/go-resty/resty/v2/example
cpu: VirtualApple @ 2.50GHz
BenchmarkResponse_TrimmingByes
BenchmarkResponse_TrimmingByes-8        149926316             7.966 ns/op  # Bytes is actually slower
BenchmarkResponse_TrimmingString
BenchmarkResponse_TrimmingString-8    153859443             7.807 ns/op  # And string uses slicing thus not creating a new string
PASS
```

So: **my assumptions were wrong**, but it took 5 minutes to check them cloning the repo, doing the changes and
running a benchmark, which is awesome.

### Output

```
Response Info:
Error      :
<nil>
    Status Code: 200
    Status : 200 OK
    Proto : HTTP/2.0
    Time : 539.360792ms
    Received At: 2022-11-26 22:02:30.439694 +0100 CET m=+0.540810335
    Body :
    <!doctype html>
    <html>
    <head>
        <title>Example Domain</title>

        <meta charset="utf-8"/>
        <meta http-equiv="Content-type" content="text/html;

Request Trace Info:
  DNSLookup     : 3.050792ms
  ConnTime      `:` 426.038041ms
  TCPConnTime   : 112.612708ms
  TLSHandshake  : 308.855291ms
  ServerTime    : 112.521375ms
  ResponseTime  : 966.584µs
  TotalTime     : 539.360792ms
  IsConnReused  : false
  IsConnWasIdle : false
  ConnIdleTime  : 0s
  RequestAttempt: 1
  RemoteAddr    : 93.184.216.34:443
```

## Iteration 2

Standard [net/http](https://pkg.go.dev/net/http) client, using the `Get(...)` method.

### Output

```
HTTP Call time: 467.371291ms

HEADERS
        Content-Type: [text/html; charset=UTF-8]
        Date: [Sat, 26 Nov 2022 21:36:14 GMT]
        Accept-Ranges: [bytes]
        Cache-Control: [max-age=604800]
        Last-Modified: [Thu, 17 Oct 2019 07:18:26 GMT]
        Server: [ECS (dcb/7F3A)]
        Vary: [Accept-Encoding]
        X-Cache: [HIT]
        Age: [497136]
        Etag: [" 3147526947
        "]
        Expires: [Sat, 03 Dec 2022 21:36:14 GMT]
        BODY
        <!doctype html>
        <html>
        <head>
            <title>Example Domain</title>

            <meta charset="utf-8"/>
            <meta http-equiv="Content-type" content="text/html;
```


## Iteration 3

Standard [net/http](https://pkg.go.dev/net/http) client, using `http.Client{}`.

### Learnings

* Using TLS directly is easier than I thought
* Implementing TLS seems a good learning experience, but might involve reading a lot of standards

### Output

```
Establishing the TLS connection...
Established.
Marking the HTTP request...
Done.
Reading HTTP response...
Done.
HTTP Header response:
         HTTP/1.1 200 OK
HTTP Headers response:
         Age: 477842
         Cache-Control: max-age=604800
         Content-Type: text/html; charset=UTF-8
         Date: Sun, 27 Nov 2022 20:45:25 GMT
         Etag: " 3147526947+ident
            "
            Expires: Sun, 04 Dec 2022 20:45:25 GMT
            Last-Modified: Thu, 17 Oct 2019 07:18:26 GMT
            Server: ECS (dcb/7F3C)
            Vary: Accept-Encoding
            X-Cache: HIT
            Content-Length: 1256
            Connection: close
            HTTP Header response:
            <!doctype html>
            <html>
            <head>
                <title>Example Domain</title>

                <meta charset="utf-8"/>
                <meta http-equiv="Content-type" content="text/html;
```


## Iteration 4

Using a parsing library like [participle](https://github.com/alecthomas/participle) or maybe the much simpler [test/scanner](https://pkg.go.dev/text/scanner) to parse the HTTP Header, HTTP Headers and then the body (using the Content-Length).

**TODO**.






