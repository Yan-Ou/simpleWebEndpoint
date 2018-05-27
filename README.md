## Just checkout this repository?
* Step 1: Ensure you have Golang 1.9.4 installed.
* Step 2: Ensure GOPATH is correctly configured.
* Step 3: Ensure port 8000 is not in use on your localhost.

## Build
make build

## Run the 
./webapp 

## Unit test
go test -v -race ./...

## Access the app endpoints
1. Access the root endpoint: curl -i localhost:8000
2. Access the health endpoint: curl -i localhost:8000/health
  1) The threshold vaule for diskspace check is 1%, hence it will always throw errors.
  2) To modify the threshold vaule for diskspace check, modify the second argument in health.go:
      healthcheck.WithObserver(
			"diskspace", checkers.DiskSpace("/var/log", 1),
		),
  3) The hearbeat check is on endpoint "/", modify the second argument in health.go:
      healthcheck.WithChecker("heartbeat", checkers.Heartbeat("/"))
3. Access the metadata endpoint: curl -i localhost:8000/metadata
