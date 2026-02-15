# WEIGHTED SLOPE ONE

Go implementation of the **weighted slope one** rating-based collaborative filtering scheme.

## Installation

```bash
go get github.com/mhdcodes/slopeone-go
```

## Usage

this library is designed to be very simple and straightforward to use. All you have to do is to load rating data,
then predict future ratings based on the training set provided.

### Creating an instance

```go
import slopeone "github.com/mhdcodes/slopeone-go"

algorithm := slopeone.New()
```

### Adding Rating values

Adding Rating values can be easily done by providing a slice of user ratings via the `Update()` method:

```go
data := []map[string]float64{
    {"squid": 1, "cuttlefish": 0.5, "octopus": 0.2},
    {"squid": 1, "octopus": 0.5, "nautilus": 0.2},
    {"squid": 0.2, "octopus": 1, "cuttlefish": 0.4, "nautilus": 0.4},
    {"cuttlefish": 0.9, "octopus": 0.4, "nautilus": 0.5},
}

algorithm.Update(data)
```

### Predicting ratings

To predict ratings for a new user, call the `Predict` method:

```go
results := algorithm.Predict(map[string]float64{
    "squid": 0.4,
})
```

This produces the following results:

```go
map[string]float64{
    "cuttlefish": 0.25,
    "octopus":    0.23333333333333,
    "nautilus":   0.1,
}
```

## Running the tests

```bash
go test -v
```

## Built With

- [Go](https://golang.org) - The programming language used

## Changelog

Please see the [changelog](./changelog.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING.md](./CONTRIBUTING.md) for details and a todo list.

## Security

If you discover any security related issues, please email author instead of using the issue tracker.

## Credits

- [Daniel Lemire](https://github.com/lemire)
- [SlopeOne Predictors for Online Rating-Based Collaborative Filtering](https://www.researchgate.net/publication/1960789_Slope_One_Predictors_for_Online_Rating-Based_Collaborative_Filtering)

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see
the [tags on this repository](https://github.com/mhdcodes/slopeone-go/tags).

## License

Please see the [Licence](./LICENSE) for more information.
