### NASA Collage

One of the most popular websites at NASA is the [Astronomy Picture of the Day](https://apod.nasa.gov/apod/astropix.html). And it uses [NASA Open API](https://api.nasa.gov/api.html#apod) which returns image of the day by specifying date in the format YYYY-MM-DD.

```bash
curl https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&date=2018-12-13
```

Using this API and Go build the smallest possible rectangular collage. Or return an error if there is no possible combination between all dates.

Rules:

- One photo can't be used more than once on the collage.
- It should be more than one photo on the collage.

Note: you may hit the rate limit (1000 requests per hour).

Submit PR with code and photo file `collage.png`.

### Prize

This is a Christmas challenge, so the first person who submits the solution will receive a nice sticker pack from me (I will ask you your address after receiving the correct PR)!