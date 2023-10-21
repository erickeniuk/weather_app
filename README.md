To clone, `cd` to the parent directory of your choice, then run

```
git clone https://github.com/erickeniuk/weather_app.git
```

# weather_app
A weather app that returns various information about weather. Backend in Go. Simple frontend with HTMX and possibly create a Flutter app with it.

## Start the server

Run the gin server with:

```
cd ./weather_app # <-- Should see router.go here
go run .
```

To test, open a separate terminal and run `curl` command:

```
# Get weather in Chicago
curl "localhost:8080/weather?city=Chicago
```

You should see some HTML return which if copy-pasted into an .html file would display a web page with your weather data in table format.

Alternatively, run the server and see below:

# HTMX Frontend
This weather app's frontend is designed using HTMX. HTMX is an open-source project from [bigskysoftware/hmtx @ GitHub](https://github.com/bigskysoftware/htmx) .

HTMX offers some nice, fast front-end features without writing a bunch of JS. Nice.

Run your server using `go run .` and confirm it's successfully running. Then, go to your browser and enter `localhost:8080`, or whichever port you chose to run on, and you should see:

![Basic Front End](./assets/basic_frontend_v1.png)

From there, you can enter the city of your choose and a nice css smooth transition should swap out the search bar with your results.