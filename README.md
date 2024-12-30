# Shorturl

A [Roadmap.sh project](https://roadmap.sh/projects/url-shortening-service)

## Launching

It is packaged with docker, and uses a redis db to store all the urls. 
```sh
docker compose up -d
```

The app is available on `:8080`. 

## How to use it

To get a shortened URL, you can make a `POST` request to `/shorten` by passing the url in a JSON body

```
POST /shorten
{
  "url": "https://<longUrl>"
}
```

Returns something like

```json
{
  "id": "<hash>",
  "longUrl": "https://<longUrl>",
  "shortUrl": "https://<host>/<hash>",
}
```

You can now access the original long url by making a request to `https://<host>:<hash>`, which will redirect you to the original long url.