# fileuploadserver
Building a file upload server over layer 7 and then trying layer 4.

## Background

Randomly came upon the idea of "Okay, I've sadly not coded as much as doing testing or ops work recently. Might be fun to write some go code."

Found this quora [link](https://www.quora.com/What-are-some-beginner-network-programming-project-ideas) and was intrigued by the idea of building a file upload server. The premises mentioned in the article were this:

1. Client/Server using regular sockets, nothing fancy, language was C/C++
2. Server can support one socket at a time, no parallel connections
3. Client connects to server and sends a file, file will be under 32MB. (File can be normal file or .zip file).
4. Server doesn’t need the file name, it can save it in any name
5. One hour to complete, with Internet, also you can copy+paste code from any source

To expand this further:
- Server supports concurrent uploads
- Client sends over file name
- No limits on file size
- Test on LAN, then on WAN
- More advanced: support SSL

I decided that I would first try this at layer 7 and then try writing a socket file upload server as well.

# Layer 7 Notes

Server side - use net/http and assume data is an HTTP request of a multipart form format
Client side - use curl

Some links:
- [https://medium.com/@petehouston/upload-files-with-curl-93064dcccc76](https://medium.com/@petehouston/upload-files-with-curl-93064dcccc76)
- [https://medium.com/@owlwalks/dont-parse-everything-from-client-multipart-post-golang-9280d23cd4ad](https://medium.com/@owlwalks/dont-parse-everything-from-client-multipart-post-golang-9280d23cd4ad)
- [https://tutorialedge.net/golang/go-file-upload-tutorial/](https://tutorialedge.net/golang/go-file-upload-tutorial/)

## Layer 4 Notes

Seems similar but server-side basically reading bytes from every underlying TCP or UDP packet and adding to a bytes buffer - which is eventually written to a persistent file.
- [https://www.youtube.com/watch?v=lmW-N07KX88&feature=youtu.be](https://mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket/)

Question to ask:
- are the bytes held in memory before writing to disk?
- what are the memory implications of this with multiple users uploading files
- how do we handle .zip files and uploads?

I did a wee bit of research on how the go net/http library parses data:

1. Essentially it opens a tcp socket, grabs data into a buffer
2. Reads buffer line by line (request is newline separated) and parses various fields
3. End of header is indicated by an empty line (before the body)
4. Passes the unmarshalled struct and buffer to another function whch then reads the body

A remaining question tho - how are giant HTTP requests handled that are spread over multiple tcp packets? Here's an interesting read:
https://stackoverflow.com/questions/60124769/can-http-1-1-requests-be-split-up-into-multiple-tcp-messages

I _think_ that the net library in Go handles a lot of the underlying work and just pushes all of the data to a buffer.
- i.e. it's reading tcp packets (which are received in order) and pushes the body to a buffer
- the entire buffer is then made available when making calls from the net library

Doing this in C for example using just the std library would probably require the engineer's application to handle reassembly of all the data in one buffer. 

## Other

Then I got distracted and decided to read about details of bits/bytes:
[https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go-8a99012dcc8f](https://www.youtube.com/watch?v=lmW-N07KX88&feature=youtu.be)

Note I also didn't implement FTP or an SFTP server b/c I was feeling quite lazy and had no desire to read the many RFCs on the FTP protocol.
