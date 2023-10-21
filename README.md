# weather_app
A weather app that returns various information about weather. Backend in Go. Planning to frontend with HTMX and possibly create a Flutter app with it.

Run the gin server with:

```
cd ./<parent_dir> # <-- Should see weather_app.go here
go run .
```

To test, open a separate window and run `curl` command:

```
# Get weather in Seattle
curl "localhost:8080/weather?city=Seattle
```

# HTMX Frontend
This weather app's frontend is designed using HTMX. HTMX is an open-source project from [bigskysoftware/hmtx @ GitHub](https://github.com/bigskysoftware/htmx) .
