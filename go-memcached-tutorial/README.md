# go-memcached-tutorial

Small Go service that demonstrates a read-through cache with Memcached in front of PostgreSQL.

The API exposes `GET /names/{id}`. For each request it:

1. Tries to fetch the record from Memcached.
2. Falls back to PostgreSQL on a cache miss.
3. Stores the database result back in Memcached.
4. Returns the record as JSON.

The code uses:

- Go `net/http` with `gorilla/mux`
- PostgreSQL via `pgx/v5`
- Memcached via `gomemcache`
- Structured logs via `log/slog`

## Project Layout

- `main.go`: HTTP server and request flow
- `postgres.go`: PostgreSQL connection pool and lookup query
- `memcached.go`: Memcached client, gob encoding, and cache TTL handling
- `schema.sql`: `names` table definition
- `compose.yml`: PostgreSQL and Memcached containers
- `justfile`: local development commands

## Requirements

- Go `1.26.1`
- `podman` with `podman compose`
- `just` for the provided task runner commands

You can adapt the container commands to Docker Compose if preferred, but the checked-in workflow uses Podman.

## Configuration

The API expects these environment variables:

- `DATABASE_URL`: PostgreSQL connection string
- `MEMCACHED`: Memcached address in `host:port` format

Values used by the `justfile`:

```env
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
MEMCACHED=localhost:11211
```

## Run Locally

Start PostgreSQL and Memcached:

```bash
just up
```

Wait for PostgreSQL to become ready:

```bash
just wait-db
```

Create the schema:

```bash
just db-init
```

Add a sample row so the endpoint has something to return:

```bash
podman compose exec -T postgres psql -U user -d dbname -c \
  "INSERT INTO names (nconst, primary_name, birth_year, death_year) VALUES ('nm0000102', 'Kevin Bacon', '1958', '') ON CONFLICT DO NOTHING;"
```

Start the API:

```bash
just api
```

Query the service:

```bash
just curl-name nm0000102
```

Or call it directly:

```bash
curl -s http://localhost:8080/names/nm0000102
```

Example response:

```json
{
  "nconst": "nm0000102",
  "name": "Kevin Bacon",
  "birthYear": "1958",
  "deathYear": ""
}
```

## Notes

- Cache entries are stored with gob encoding.
- Cached records expire about 25 seconds after being written.
- `schema.sql` creates the table but does not load IMDb-style source data.
- The current handler returns an internal error when PostgreSQL does not find a row; it does not map missing records to `404`.
