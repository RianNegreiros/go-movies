# Go Movies

A movie catalogue web application. Browse movies, manage a favorites list and a watchlist, with user authentication.

## Stack

- **Backend**: Go, standard `net/http`
- **Frontend**: Vanilla JS (Web Components, client-side routing)
- **Database**: PostgreSQL
- **Auth**: JWT

## Setup

**1. Configure environment**

Copy `.env.local` and set your values:

```
DATABASE_URL=postgres://user:password@localhost:5432/gomovies?sslmode=disable
JWT_SECRET=<random secret>
```

**2. Start the database and seed data**

```sh
make setup
```

**3. Run**

```sh
make run
```

Serves on `http://localhost:8080`.

## API

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/movies/top` | Top movies |
| GET | `/api/movies/random` | Random movies |
| GET | `/api/movies/search` | Search (`?q=`, `?order=`, `?genre=`) |
| GET | `/api/movies/{id}` | Movie by ID |
| GET | `/api/genres` | All genres |
| POST | `/api/account/register` | Register |
| POST | `/api/account/authenticate` | Login, returns JWT |
| GET | `/api/account/favorites` | User favorites |
| GET | `/api/account/watchlist` | User watchlist |
| POST | `/api/account/save-to-collection` | Add to favorites or watchlist |
