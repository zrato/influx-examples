
As noted in these two posts about telegraf...

[Getting started with Telegraf](https://docs.influxdata.com/telegraf/v1.10/introduction/getting-started/)

[How to Write Points from CSV to InfluxDB](https://www.influxdata.com/blog/how-to-write-points-from-csv-to-influxdb/)

This is the way to start out with a blank slate from which you can edit and continue to create your telegraf config file of your choosing.

```
telegraf -sample-config -input-filter file -output-filter influxdb > file.conf
```

```
telegraf --config ./ubnt.conf
```
