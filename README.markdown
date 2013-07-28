# groupcachedb
This project was created to play with an experiment with go and the groupcache project.

# Demo

### Build everything

1. ```git clone git@github.com:capotej/groupcachedb```
2. sh build.sh

### Start DB server

1. ```cd dbserver && ./dbserver```

This starts a delibrately slow k/v datastore on :8080

### Start Multiple Frontends

1. ```cd frontend```
2. ```./frontend -port 8001```
3. ```./frontend -port 8002```
4. ```./frontend -port 8003```

### Use the CLI to set/get values

1. ```cd cli```
2. ```./cli -set -key foo -value bar```
3. ```./cli -get -key foo``` should see bar in 300 ms
4. ```./cli -cget -key foo``` should see bar in 300ms (cache is loaded)
5. ```./cli -cget -key foo``` should see bar instantly
