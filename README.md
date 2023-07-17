# kik_go_api

kik_go_api is a Go package for interacting with the Kik messaging API. It allows you to build bots and clients that can send/receive messages on Kik.

The package was inspired by the [kik-node-api](https://github.com/YassienW/kik-node-api) Node.js library.

## Usage

To use kik_go_api, first create a Client instance:

```go
client := kik_go_api.Client
```

Then configure the client with your username, password, and optionally the Kik version:

```go
client.Settings("username", "password") 
// Kik version defaults to 15.25.0.22493

client.Settings("username", "password", "15.15.0.12345", "pNtboj79GGFYk9w2RbZZTxLpZUY=")  
// Custom Kik version and SHA1 hash
```

> NOTE: The Kik version defaults to 15.25.0.22493 if not provided. Only the major and minor version numbers are needed.

Next, connect the client:

```go 
messages := make(chan string)
go client.Connect(messages)
```

This will start receiving messages on the `messages` channel. 

To send messages, use the `SendMsg` method:

```go
client.SendMsg("Hello world!", "someone@talk.kik.com") // Send to a user
client.SendMsg("Hello group!", "group@groups.kik.com") // Send to a group
``` 

There is also a `SendRaw` method for sending raw XML stanzas:

```go
client.SendRaw(`
    <message to="someone@talk.kik.com">
      <body>Hello!</body> 
    </message>
`)
```

### Receiving Messages

To parse incoming messages, you can pass the `msg` string from the `messages` channel to your own parsing function. 

For example, here is a basic parser from `bot.go.example`:

```go
for msg := range messages {

  if strings.Contains(msg, "type=\"chat\"") {
      // Extract JID 
      jid := extractJid(msg)  

      // Build message stanza
      stanza := buildMessageStanza("Hello!", jid) 

      // Send message
      client.SendRaw(stanza)
  }

}

func extractJid(msg string) string {
  // Use regexp to extract JID 
  return jid 
}

func buildMessageStanza(body string, to string) string {
  // Build XML stanza
  return stanza
}
```

This extracts the incoming message JID and sends a response.

You can replicate the parsing logic from other Kik API wrappers like [kik-bot-api-unofficial](https://github.com/tomer8007/kik-bot-api-unofficial) (Python) or [kik-node-api](https://github.com/YassienW/kik-node-api) (Node.js).

## Contributing

Contributions are welcome! kik_go_api is missing some features like group chat and read receipts. See the [issue tracker](https://github.com/nanofuxion/kik_go_api/issues) for details.

## License

MIT License - see [LICENSE](LICENSE) for details.
