# Pet adoption platform

## build

Clone the repo locally:

```sh
git clone git@github.com:zarasfara/pet-adoption-platform.git
```

Run command:

Linux/Mac:
```sh
cp .env.example .env
```

Windows:
```sh
copy .env.example .env
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