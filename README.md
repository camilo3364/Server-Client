# Server/Client

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Server/Client is a communication program using the TCP protocol for reception and transmission to different clients that are subscribed to a specific channel. It was developed using golang.

## Features


- server.go has the characteristics of active listening to the client that have a file to a specific channel and transmission of data to said channel for clients connected to said channel.

- client.go is in charge of requesting the user through the terminal, if he wants to send a file to the server or receive it through a certain channel.
- ✨Magic ✨

## Libraries

- [bytes] - Convert files to bytes
- [fmt]   - print by console
- [math/rand] - Generate random numbers for TCP ports
- [net]  - Generate the connections
- [os] -  Open computer files
- [strconv] - Convert integers to string
- [strings] - Convert to string
- [time] - Get a different random number using the current date.

![Server waiting for files to be sent by a client in the state of sending files and 3 clients waiting for files 2 on channel 1 and 1 on channel 2.]

[![example.jpg](https://i.postimg.cc/q7rvv3mp/example.jpg)](https://postimg.cc/dhWYWLPf)



## License
**Free Software**
