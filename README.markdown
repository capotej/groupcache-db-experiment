# groupcache-db-experiment
This project simulates a scenario wherein a few frontends running [groupcache](http://github.com/golang/groupcache) are fronting a slow database. See my [blog post](http://www.capotej.com/blog/2013/07/28/playing-with-groupcache/) about it for more details.

# Getting it running
The following commands will set up this topology:
![groupcache topology](https://raw.github.com/capotej/groupcache-db-experiment/master/topology.png)

### Build everything

1. ```git clone git@github.com:capotej/groupcache-db-experiment```
2. ```sh build.sh```

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
