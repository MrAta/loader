# Loader: HTTP[s] endpoint concurrency limit load tester

To build:

```bash
make build
```

To run sample tests:

```bash
make run
```

To clean up:

```
make clean
```

The http endpoint has to be set as an environment variable named `HTTP_ADDR`.

The command line flags:

```bash

-rate: Integer value for the request rate (i.e. concurrent requests). 
    #Example: ./loader -rate=120

-duration: The duration of the load test: Ns, Nm, Nh (N is the rate, s-m-h are time units for second, minute, and hour).
    #Example: ./loader -rate=120 -duration=10m

-logfile: The log file name for a simple logging (currently logs time stamp and status code as that's what suffices for me!).
    #Example: ./loader -rate=120 -duration=10m -logfile=Rate120Duration10Minute.log
```


