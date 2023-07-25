# JWT Authentication With Refresh Token
This example shows how to authenticate with JWT and refresh token.

## Endpoints
- `POST /login` - Login with username and password. Returns JWT and refresh token.
- `POST /refresh` - Refresh JWT with refresh token. Returns new Access token.
- `GET /secret-page` - Protected endpoint. Returns user info.