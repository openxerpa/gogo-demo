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
| `TAILSCALE_AUTHKEY` | Tailscale auth key for deploy network access |
| `TAILSCALE_OAUTH_CLIENT_ID` | Tailscale OAuth client ID (if using OAuth) |
| `TAILSCALE_OAUTH_SECRET` | Tailscale OAuth secret (if using OAuth) |
| `PROD_SERVER_IP` | Deploy host for main branch |
| `DEV_SERVER_IP` | Deploy host for dev branch |
| `TEST_SERVER_IP` | Deploy host for test branch |
| `PROD_SSH_KEY` | SSH private key for prod |
| `DEV_SSH_KEY` | SSH private key for dev |
| `TEST_SSH_KEY` | SSH private key for test |
| `DEPLOY_REGISTRY_TOKEN` | GHCR token for pulling images (shared across envs) |

## Environment variables

| Variable | Description |
|----------|-------------|
| `PORT` | HTTP server port (default: 3000) |


