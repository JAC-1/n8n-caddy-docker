# n8n + Redis (Docker)

```bash
# 1. Save as docker-compose.yml
# 2. Optional: create .env with your settings
# 3. Run

docker compose up -d
```

Access: http://localhost:5678

### Files
- `docker-compose.yml` – main file
- `.env` (optional) – set N8N_HOST, N8N_USER, N8N_PASSWORD
- Volume `n8n_data` – stores all workflows & credentials forever

### Commands
```bash
docker compose pull        # update n8n image
docker compose up -d       # restart with new version
docker compose down        # stop
docker compose logs -f n8n # view logs
```

All workflows are safe. Just copy this folder to any server and run `docker compose up -d`.

# Future Vector DB

Add Qdrant (or Chroma) in one line when you want real AI memory.
Replaces redis.

```yaml
# Add this service to your existing docker-compose.yml
  qdrant:
    image: qdrant/qdrant:latest
    restart: unless-stopped
    ports:
      - "6333:6333"
    volumes:
      - qdrant_data:/qdrant/storage

volumes:
  qdrant_data:
```

# Clean Up

```bash
# Stop and remove ALL containers
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

# Remove ALL images
docker rmi -f $(docker images -q)

# Remove ALL volumes (including your n8n_data !)
docker volume rm $(docker volume ls -q)

# Remove ALL networks (except default ones if you want)
docker network prune -f

# One-liner to nuke literally everything in one go (fastest)
docker system prune -a --volumes -f
```
