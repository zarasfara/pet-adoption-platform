# Pet adoption platform

## build

Clone the repo locally:

```sh
git clone git@github.com:zarasfara/pet-adoption-platform.git
```

Run command:
```sh
make init
```

Set credentials in .env

Run this command to build app:

```sh
docker compose up -d --build
```

Run migrations:

```sh
make migrate-up
```