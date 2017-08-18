# Simple consumer
Connects to RPC server on default port and prints response.

# http consumer using swagger proxy

`curl -X POST -k http://localhost:7070/v1/example/hello -H "Content-Type: text/plain" -d '{"name": "foo"}'`