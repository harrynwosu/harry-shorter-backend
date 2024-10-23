# harry-shortener

This application is my own personal implementation of a URL shortener, converting long and verbose website links to shorter and manageable ones. Think _Bit.ly_ or _TinyUrl_.

This is my first stab at a project with Go while learning, so baby steps 👣 !

#### To-do / Coming soon:

-   [] Dockerize and deploy the service in the cloud
-   [] Deploy (possibly with _Vercel_\*\*) user-friendly UI implementation with React available on the web for consuming this service
-   [] Add a SQL DB (Postgres or MySQL) as backup to the Redis main store to be used for storing cold values/least frequently fetched (LFU) data.
-   [] Implement QR code generator and add to the frontend as well.

## Main Features

-   SHA-256 hashing and Base54 encoding used for generating the short links from the longer URLs.
-   Redis used for fast data retrieval and storing mappings between short links and their orinigal URLs.
-   Unit tests implemented for major functional components.

## Usage

1. Start the web server my running the `main.go` file with:

```
go run main.go
```

2. Start the Redis server in another terminal window with:

```
redis-server
```

## Testing

You can use any REST API tester of your choice to test following the example convention below. My personal favorite is Postman.

#### Request

```
POST http://localhost:9808/generate-short-url
```

#### Body

```
{
    "original_url": "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}
```

#### Expected response

```
{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/Xbm18Q"
}
```

Visiting the output short URL generated `http://localhost:9808/Xbm18Q`in your browser appropriately redirects you to the site of the original link, `https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator`
