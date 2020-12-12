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

## Other

Then I got distracted and decided to read about details of bits/bytes: 
[https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go-8a99012dcc8f](https://www.youtube.com/watch?v=lmW-N07KX88&feature=youtu.be)

Note I also didn't implement FTP or an SFTP server b/c I was feeling quite lazy and had no desire to read the many RFCs on the FTP protocol. 