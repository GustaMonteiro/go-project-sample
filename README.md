# Go Project Sample

The purpose of this repo is just to test Go workspaces and modules, and also know how to create a dockerfile for individual applications and a docker-compose.yml that runs every app in a repo.

## Running docker-compose

``` bash
docker compose up
```

## Building individual app

``` bash
docker build -t <consumer/producer> -f apps/<Consumer|Producer>/dockerfile .
```
