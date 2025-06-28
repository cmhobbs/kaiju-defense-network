# Kaiju Defense Network!

This is a simple project I'm noodling with to refamiliarize myself with Go.

## Usage

To generate a random kaiju sighting and print it to `stdout`, run the following command from the `kaiju-sightings-generator` directory:

```
go run main.go
```

To log a random kaiju sighting to `stdout`, run the following command from the `kaiju-alert-system` directory:

```
go run main.go
```

To generate multiple sightings, use the `generateMultiple()` function with a count parameter in your code:

```
generateMultiple(5)
```

To run tests, use `go test -v` in each desired directory.

## Kaiju Sightings Generator Tasks
- [x] Generate a random kaiju sighting with name/location/threat level
- [x] Output the sighting to the terminal
- [x] Add size and element traits to each kaiju randomly
- [x] Add a timestamp to each sighting ("2025-06-20 22:14")
- [x] Generate multiple sightings in one run

## Kaiju Alert System
- [x] Output an alert to stdout when a high-threat kaiju is sighted above a defined threshold
  - [x] write tests for `shouldAlert()` and `formatAlert()`
- [x] Output an action based on the threat level (e.g., evacuate, deploy defenses)
- [x] Add timestamp to each alert from sighting
- [ ] Write alerts to a log file like `alerts.log` including full details
- [ ] Process multiple sightings in one batch

## Further Review:  Connecting Sightings and Alerts

- [ ] Modify the modules to call each other by function
- [ ] Refactor `Generate()` to be less complicated/cute.  Pointers might not be necessary.
- [ ] On a separate branch, modify the alert system to read sightings from a JSON file from the generator
- [ ] On a seprate branch, use a Go channel to send sightings from the generator to the alert system (alert system running in a separate goroutine and listening on the channel)
- [ ] On a separate branch, stuff each service into podman containers
  - [ ] Make them communicate with a shared volume and the JSON file
  - [ ] Make them communicate using Valkey (generator publishes, alert system subscribes)
  - [ ] Make them communicate with a REST API
  - [ ] Get weird and throw them into minikube
  - [ ] Add more output to stdout/stderr and shovel logs somewhere neat
  - [ ] Add `HEALTHCHECK` instructions to the Containerfiles
