# Steps to run ChatApp

## Server

1. Load terminal and cd to backend folder. Or right click open in integrated terminal.
2. Run command: "go run main_test.go". This will test backend.
3. Run command: "go run main.go". This will launch the server.

## Client

1. Load terminal and cd to frontend folder. Or right click open in integrated terminal.
2. Run command: "npm i". This will install all dependencies.
3. Run command: "npm run dev". This will launch the frontend and provide a local host url to connect to. http://127.0.0.1:5173/ | You can also open a browser and type copy paste this url.
4. After a client has successfully connected the backend terminal will display a message similar to this:

Size of Connection Lobby: 1
&{ 0x1400019a000 0x1400006e140}
This is an example of a client connected successfully which shows the loby size and address of the client and current lobby.

If a second client connects the lobby size will increase to 2 (eg):
Size of Connection Lobby: 2
&{ 0x140000ce000 0x1400006e140}
&{ 0x1400019a000 0x1400006e140}

The new client address is added, and the lobby address is displayed for both clients.

## Multiple Clients

### Mac OS

1. Launch client and connect.
2. After connecting to first client, navigate to system preferences > Apple ID > toggle on Private Relay.
3. Launch safari browser private window and connect to a new client using localhost url.
4. This client will have a separate IP address. You can simulate 2 separate clients commnicating with each other.

### Windows

1. Launch client and connect.
2. Launch second client and connect.
3. Communicate between clients.

_NB_ use vpn to create second client with different ip address.

## Chatrooms

1. Clients are first loaded into a general lobby where they can begin communicating with each other.
2. To create a new room, input the room name and press "enter" or click "Add".
3. all room names will display for each client.
4. The user that creates the room will automatically enter it.
5. Other clients can join the room by clicking it.

_NB_ Currently there is no method to return to main lobby. Refresh browser to return to main lobby.
