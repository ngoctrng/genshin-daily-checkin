## How to run

#### Running on Docker Swarm

Example:

```shell
docker service create --name genshin-daily-login truong996/genshin-daily-login
```

#### Running Docker only

```shell
docker run --env COOKIE="your-cookie" -d truong996/genshin-daily-login
```
