# Websocket example

The OpenAPI specification doesn't support WebSockets yet so this is an example
that shows how to use the websocket to stream live data from Span.

## API Token

If you use the WebSocket with a query parameter to authenticate the API token
must be set to read-only since the token could potentially end up in various
request logs on proxies on the Internet. If it does and somebody tries to access
the Span service there's nobody with write access to your data.
