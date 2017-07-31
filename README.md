# Device Hub
Search mobile device model info.


## Build
```shell
$ glide update && glide install && go build
```

## How To Run?

Docker is recommended:

```shell
$ docker run --rm -ti -p 3000:3000 jjeffcaii/devicehub:latest
```

## Restful Docs

Example: search for `iPhone7,2`

```shell
 $ curl -X GET \
    -H "Content-Type: application/json" \
    "http://127.0.0.1:3000/iphone7,2"
```

It will return JSON string:

```json
"iPhone 6"
```

## FAQ

> Q: Where can I get latest device models?

> A: Just execute script below.

```shell
$ curl http://storage.googleapis.com/play_public/supported_devices.csv | iconv -f UTF-16 -t UTF-8 > db/android.txt
$ curl -o db/ios.txt http://7xpeg1.com1.z0.glb.clouddn.com/ios.devicemodels.properties
```
