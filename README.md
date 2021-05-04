# URL Shortener using Golang and Redis
![banner](https://socialify.git.ci/skamranahmed/golang-url-shortener-with-redis/image?description=1&font=Inter&language=1&owner=1&pattern=Floating%20Cogs&theme=Light)


## Architecture

#### Generate a Short URL from a Long URL
![generate-short-url-from-long-url]

#### Redirect a Short URL to the Original Long URL
![redirect-short-url-to-long-url]

## API Reference

#### Generate a Short URL from a Long URL

```http
  POST /generate
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `long_url` | `string` | **Required**. Long URL that you want to convert to a Short URL |

#### Redirect Short URL to Original URL

```http
  GET short-url
```

#### Get Short URL Info

```http
  GET short-url/info
```

[generate-short-url-from-long-url]: architecture/url-shortener-1.png
[redirect-short-url-to-long-url]: architecture/url-shortener-2.png
