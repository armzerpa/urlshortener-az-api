# urlshortener-az-api
Go practice project 

Creator: Armando Zerpa<armando.zerpa@mercadolibre.com.co>

Descripcion: Api para acortar URLs

POST: /v1/shortener

{
    "url": "<some_long_url>"
}

Respuesta:
{
    "id": "B80q91Q",
    "longUrl": "https://github.com/gin-gonic/gin#grouping-routes",
    "shortUrl": "http://localhost:8080/u/B80q91Q"
}


GET: /u/<ID>

Respueta: Redirect a la URL original


