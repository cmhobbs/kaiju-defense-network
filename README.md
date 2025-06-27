# Kaiju Defense Network!

This is a simple project I'm noodling with to refamiliarize myself with Go.

## Kaiju Sightings Generator Tasks
- [x] Generate a random kaiju sighting with name/location/threat level
- [x] Output the sighting to the terminal
- [x] Add size and element traits to each kaiju randomly
- [ ] Add a timestamp to each sighting ("2025-06-20 22:14")
- [ ] Append sightings to a JSON file
- [ ] Generate multiple sightings in one run

## Kaiju Alert System
- [x] Output an alert to stdout when a high-threat kaiju is sighted above a defined threshold
  - [ ] write tests for `shouldAlert()` and `formatAlert()`
- [ ] Output an action based on the threat level (e.g., evacuate, deploy defenses)
- [ ] Write alerts to a log file like `alerts.log` including full details
- [ ] Process multiple sightings in one batch

## Further Review:  Connecting Sightings and Alerts

- [ ] Modify the modules to call each other by function
- [ ] Modify the alert system to read the sightings from the JSON file
- [ ] Use a Go channel to send sightings from the generator to the alert system (alert system running in a separate goroutine and listening on the channel)
- [ ] Stuff each service into podman containers
  - [ ] Make them communicate with a shared volume and the JSON file
  - [ ] Make them communicate using Valkey (generator publishes, alert system subscribes)
  - [ ] Make them communicate with a REST API
  - [ ] Get weird and throw them into minikube
  - [ ] Add more output to stdout/stderr and shovel logs somewhere neat
  - [ ] Add `HEALTHCHECK` instructions to the Containerfiles
