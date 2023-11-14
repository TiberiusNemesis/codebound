# What is HTTP?

HTTP (aka Hypertext Transfer Protocol) is an application-layer protocol for communicating over a network. It became the main communication protocol for the internet ever since its introduction in the early 90s. While it was originally supposed to carry only text (hence the name hypertext), it was quickly upgraded and now other forms of media, such as images or videos, were able to be transmitted through HTTP. So technically, the correct name would be Hypermedia Transfer Protocol. HMTP. 

Doesn't quite roll off the tongue, does it?

## HTTP 1.x vs HTTP 2

HTTP 1.0 had just a GET method. It was pretty barebones and just retrieved HTML.

HTTP 1.1 added several methods, including POST, PUT, HEAD and DELETE. It added status codes and headers, as well as 

HTTP 2.0 introduces several new concepts such as streams and push promises. 

Streams are a way of distributing data in a way that is more efficient than previous ones -- they are bidirectional channels that can carry any number of messages. Within these streams, the messages are formed of the smallest data unit called _frames._ A logical HTTP message, such as a request and a response, is made up of multiple frames. 

Push promises are basically a server's way of efficiently telling the client that they should not make additional requests for some content because they are _promised_ to be delivered that content.

These are all designed to improve speed but require the server side to actually implement these features in a decent way to be able to make use of them. 

## HTTPS

HTTPS (aka HTTP Secure) is a security protocol built on top of HTTP which allows the client to:
- Make sure their data is private when being sent
- Make sure that the data is not being tampered with when communicating with a server
- Be certain that the destination they are sending their message to is the real server they are looking to communicate with

It does this through encryption, most often through TLS (often called SSL/TLS). This involves a handshake which has several steps.
1. The client sends over a hello and *cipher suite,* containing the encryption algorithms it can work with.
2. The server sends over a hello with its choice of the best SSL/TLS version, one or more encryption algorithms among the ones the client sent, and their signing certificate + public key.
3. The client responds with a pre-master (or client) key which is encrypted with the server's public key and sends it back.
4. They both generate a 'shared secret' that can be used as a symmetric key, and a test is sent from both sides.
5. Finally, after a successful test, connection is now secured.