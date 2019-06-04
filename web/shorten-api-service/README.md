### Shorten-URL-Service

基于Beego的短域名服务API，提供两个接口，使用如下：
```
# shortening url example
http://localhost:8080/v1/shorten/?longurl=http://google.com

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}

# expanding url example
http://localhost:8080/v1/expand/?shorturl=5laZG

{
  "UrlShort": "5laZG",
  "UrlLong": "http://google.com"
}
```
