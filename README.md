```
        ██████╗  █████╗ ██████╗ ██████╗ ██╗     ███████╗███████╗███████╗██████╗ ██╗   ██╗
        ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██║     ██╔════╝██╔════╝██╔════╝██╔══██╗██║   ██║
        ██████╔╝███████║██████╔╝██████╔╝██║     █████╗  ███████╗█████╗  ██████╔╝██║   ██║
        ██╔══██╗██╔══██║██╔══██╗██╔══██╗██║     ██╔══╝  ╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝
        ██████╔╝██║  ██║██████╔╝██████╔╝███████╗███████╗███████║███████╗██║  ██║ ╚████╔╝
        ╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝
```

Babbleserv is a Matrix homeserver built on top of FoundationDB primitives.

# Setup

1. Install the FoundationDB client and server packages from [here](https://github.com/apple/foundationdb/releases/latest)

2. Build Babbleserv (and my key gen thingy):
```
go build -o babbleserv cmd/babbleserv/main.go
go build -o genkey cmd/genkey/main.go
```

3. Generate a federation key (ed25519):
```
./genkey myverysecurechmod777.key
```

4. Generate TLS keys:
```
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem
```

5. Copy `config.example.yaml` to `config.yaml` and adjust as needed.

6. Run it! `./babbleserv -routes -prettyLogs`

# Usage

Authentication is done by setting your username as the token:
```
curl -H "Authorization: Bearer user123" -k https://localhost:8888/_matrix/client/v3/createRoom --json '{}'
{"room_id":"!gbeG4kGrgFAm4vld6hg4sg:localhost"}
```
