# Kaiju Defense Network!

This is a simple project I'm noodling with to refamiliarize myself with Go.

## Usage

To log a few random kaiju sightings to `alert.log` and display them in stdout, run the following command from the `kaiju-alert-system` directory:

```
$ go run main.go
────────────────────────────────────────────────────────────
🚨 KAIJU SIGHTING ALERT AT 2025-06-27 22:16:43

A Gigantic Megadragon has been spotted in Beijing,
exhibiting Ambush behavior!

It is a Critical threat.

⚠️  SIGHTING DETAILS

	👾 Name: Megadragon
	📍 Location: Beijing
	⚡ Threat Level: Critical
	📏 Size: Gigantic
	🎭 Behavior: Ambush
	⏰ Timestamp: 2025-06-27 22:16:43
────────────────────────────────────────────────────────────
Threshold hit, logging alert...

────────────────────────────────────────────────────────────
🚨 KAIJU SIGHTING ALERT AT 2025-06-27 22:16:43

A Supreme Neomoth has been spotted in Rome,
exhibiting Patrol behavior!

It is a Medium threat.

⚠️  SIGHTING DETAILS

	👾 Name: Neomoth
	📍 Location: Rome
	⚡ Threat Level: Medium
	📏 Size: Supreme
	🎭 Behavior: Patrol
	⏰ Timestamp: 2025-06-27 22:16:43
────────────────────────────────────────────────────────────
Threshold too low, alert will not be logged...

...
```

You can then check `alert.log` for details:

```
$ cat alert.log
2025-06-27 22:16:43 - WARNING: A CRITICAL LEVEL KAIJU HAS BEEN SPOTTED!  EVACUATE IMMEDIATELY.
...
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
- [x] Write alerts to a log file like `alerts.log` including full details
- [x] Process multiple sightings in one batch

## Further Review:  Connecting Sightings and Alerts

- [x] Modify the modules to call each other by function
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
