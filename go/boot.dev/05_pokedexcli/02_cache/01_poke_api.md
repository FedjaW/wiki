# PokeAPI

Now we're going to use the PokeAPI to get some real data from the Pokemon world!

## Assignment

1. Add the `map` command. It displays the names of `20` location areas in the Pokemon world. Each subsequent call to `map` should display the next 20 locations, and so on. This will be how we explore the Pokemon world. Example usage:`

```txt
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

Here are some pointers for implementing this command:

- You'll need to use the PokeAPI location-area endpoint (<https://pokeapi.co/docs/v2#location-areas>) to get the location areas. Note that this is a different endpoint than the "location" endpoint. Calling the endpoint without an `id` will return a batch of location areas.
- Update all commands (e.g. `help`, `exit`, `map`) to now accept a pointer to a "config" struct as a parameter. This struct will contain the `Next` and `Previous` URLs that you'll need to paginate through location areas.
- Here's an example (https://pkg.go.dev/net/http#example-Get) of how to make a GET request in Go.
- Here's how to unmarshal (https://blog.boot.dev/golang/json-golang/#example-unmarshal-json-to-struct-decode) a slice of bytes into a Go struct.
- You can make `GET` requests in your browser or by using `curl`! It's convenient for testing and debugging.

2. Add the `mapb` (map back) command. It's similar to the `map` command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.

If you're on the first "page" of results, this command should just print "you're on the first page". Example usage:

```txt
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
Pokedex > map
mt-coronet-1f-route-216
mt-coronet-1f-route-211
mt-coronet-b1f
great-marsh-area-1
great-marsh-area-2
great-marsh-area-3
great-marsh-area-4
great-marsh-area-5
great-marsh-area-6
solaceon-ruins-2f
solaceon-ruins-1f
solaceon-ruins-b1f-a
solaceon-ruins-b1f-b
solaceon-ruins-b1f-c
solaceon-ruins-b2f-a
solaceon-ruins-b2f-b
solaceon-ruins-b2f-c
solaceon-ruins-b3f-a
solaceon-ruins-b3f-b
solaceon-ruins-b3f-c
Pokedex > mapb
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

### Tips

- JSON lint is a useful tool for debugging JSON, it makes it easier to read.
- JSON to Go a useful tool for converting JSON to Go structs. You can use it to generate the structs you'll need to parse the PokeAPI response. Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options. For nullable strings, use `*string`.
- I recommend creating an internal package (<https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface>) that manages your PokeAPI interactions. It's not required, but it's a good organizational and architectural pattern.
