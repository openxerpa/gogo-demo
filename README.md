# Gogo demo

A short description of the service

## Local development

```bash
docker compose up
```

## Deployment

Push to `main`, `dev`, or `test`; CI/CD builds and deploys automatically.

### Required GitHub secrets

| Secret | Description |
|--------|-------------|
| `TAILSCALE_OAUTH_CLIENT_ID` | Tailscale OAuth client ID for deploy network access |
| `TAILSCALE_OAUTH_SECRET` | Tailscale OAuth client secret (OAuth client must have `auth_keys` scope and `tag:ci`) |
| `PROD_SERVER_IP` | Deploy host for main branch |
| `DEV_SERVER_IP` | Deploy host for dev branch |
| `TEST_SERVER_IP` | Deploy host for test branch |
| `PROD_SSH_KEY` | SSH private key for prod |
| `DEV_SSH_KEY` | SSH private key for dev |
| `TEST_SSH_KEY` | SSH private key for test |
| `DEPLOY_REGISTRY_TOKEN` | GHCR token for pulling images (optional; defaults to `GITHUB_TOKEN` for same-repo) |

## Environment variables

| Variable | Description |
|----------|-------------|
| `PORT` | HTTP server port (default: 3000) |


