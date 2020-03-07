
Take a look at storage/engine_test.go and write some
code that writes data to the wal and tsm1.

Upon further exploration see
[engineX1_test.go](https://github.com/stormasm/influx-examples/blob/master/code/storage/engineX1_test.go)

cmd/influxd/launcher/launcher.go which is an expanded
version of firing up a

##### NewEngine

So this is a good starting point upon return home from Ashland...

#### Understand and pick one service to dive into and all of the associated code

```
rg BucketService
```
