# EVENT LOG DUMP

Create and retrieve a stream of events. Practising GO Rest API development.
Stores entries in local sqlite DB file `event_log_dump`.

## How to run
```bash
# Run the server
go run .

# Endpoints

# Get all events
curl http://localhost:8000/events

# Create an event
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"event_type":"UserCreated","created_at":"2022-02-02","metadata":"{hello: world}"}' \
  http://localhost:8000/events
```
