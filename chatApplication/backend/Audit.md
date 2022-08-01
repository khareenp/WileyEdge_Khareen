NewLobby() is the first function to be called
This function returns a pointer to chat room or lobby which has in 3 channels and a map

Next the start function which will broadcast all messages to connected clients
including if a user connects or disconnects.

Read() method which will constantly listen in for new messages
coming through on this Client’s websocket connection.

If there are any messages, it will pass these messages to the Pool’s Broadcast channel
which subsequently broadcasts the received message to every client within the pool.
