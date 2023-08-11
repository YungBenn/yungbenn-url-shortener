# YungBenn URL Shortener

A Simple URL Shortener API

Redis installed locally required !

## API Reference

#### Create short URL

```http
POST /create-short-url
```

```curl
curl --location 'http://127.0.0.1:9808/create-short-url' \
--header 'Content-Type: application/json' \
--data '{
    "long_url": "https://example-long-url.co.id",
    "user_id": "<Random UUID>"
}'
```

#### Short URL Redirect

```http
GET /:shortUrl
```
