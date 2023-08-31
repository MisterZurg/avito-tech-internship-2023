# Transcode HTTP/JSON to gRPC.

The first one is google/api/annotations.proto and the second is google/protobuf/timestamp.proto. We will explain both of them in more detail since I consider this part as a pain point that took most of my time while working on this project.

google/api/annotations.proto is a required third party protobuf file for the protoc compiler. If you don’t have this dependency inside your project, you can’t define your API calls (GET, POST, PUT, PATCH, DELETE) inside your services and you won’t be able to Transcode HTTP/JSON to gRPC.

Interesting enough, there’s currently no standard way to import this dependency. Even if you add the above lines inside your proto file, the protoc compiler won’t import this dependency for you because you have to do this manually.