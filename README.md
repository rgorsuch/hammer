# Run

Run the hammer like this.

```
docker build -t hammer-app .
docker run -it --rm -e RENO_URL --name hammer hammer-app
OR
docker build -t hammer-app . && docker run -e RENO_URL -it --rm --name hammer hammer-app
```
