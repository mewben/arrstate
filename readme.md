# ArrState

## Installation

1. Run ./install

- Enter the mongodb connection string
- Enter the Business Name, Given Name, Family Name, Email, and set Password
- Enter the Activation Key (Based on deviceId and business.domain)
- Save to db
  businesses
  users
  people
- Prompt exit

2. Run ./run

- This runs both the backend and front
- Browse to: http://localhost:7000

## Development

Backend: Go
Frontend: React (create-react-app)

To connect to mongo shell in docker:
docker exec -it 8ead3156cf5c mongo "mongodb://root:passwd123@localhost:27017/realtydevdb?authSource=admin"

## Testing

- To run sequentially:
  `go test ./... -p=1 -v`

## Processes

### Signup Users.signup()

payload: - person: - name: - first - last - timezone (auto) - business: - name - domain - user: - email - password

response: - currentUser: - user - person - currentBusiness - userBusinesses []

### Signin Users.signin()

payload: - email - password
response: (same with signup response)
