# SteamSpy

Simple Steam Client of https://steamspy.com/api.php
It's free and easy to get all apps from steam. The data is refreshed once a day, there is no reason to request the same information more than once every 24 hours.

Allowed poll rate:
- 1 request per second for most requests
- 1 request per 60 seconds for the *all* requests.

# Install

Currently used and tested on Go 1.24.

To get the package please enter:

```go
go get github.com/opv88/steamspy
```

In your go files simply use:

```go
import "github.com/opv88/steamspy/client"
```

# Usage

Simple Examples:

1. Returns all games with owners data sorted by owners. Returns 1,000 entries per page. *Page* is a page number for the list (starts at 0)

```go
// create client
c := client.NewClient()
// required page
page := int64(2)
apps, err := c.GetAll(page)
if err != nil {
	...
}
```

2. Returns details for the specific application. Requires *appid* parameter.
```go
// create client
c := client.NewClient()
// required appId
appId := int64(9450)
apps, err := c.GetAppDetails(appId)
if err != nil {
	...
}
```
3. Returns games in this particular genre. Requires *genre* parameter.
```go
// create client
c := client.NewClient()
// required genre
genre := "Early+Access"
apps, err := c.GetByGenre(genre)
if err != nil {
	...
}
```
4. Returns games with this particular tag. Requires *tag* parameter.
```go
// create client
c := client.NewClient()
// needable page
genre := "Early+Access"
apps, err := c.GetByGenre(genre)
if err != nil {
...
}
```
5. Returns Top 100 games by players in the last two weeks.
```go
// create client
c := client.NewClient()
apps, err := c.GetTop100In2Weeks()
if err != nil {
...
}
```
6. Returns Top 100 games by players since March 2009.
```go
// create client
c := client.NewClient()
apps, err := c.GetTop100Forever()
if err != nil {
...
}
```
7. Returns Top 100 games by owners.
```go
// create client
c := client.NewClient()
apps, err := c.GetTop100Owned()
if err != nil {
...
}
```
