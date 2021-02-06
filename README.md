# Loader: HTTP[s] endpoint concurrency limit load tester

To build:

```bash
make build
```

To run sampel tests:

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

-rate Integer value for the request rate (i.e. concurrent requests). 

-duration: The duration of the load test: Ns, Nm, Nh (N is the rate, s-m-h are time units for second, minute, and hour).

-logfile: The log file name for a simple logging (currently logs time stamp and status code as that's what suffices for me!).

```


