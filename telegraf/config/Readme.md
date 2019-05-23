
This is the way to start out with a blank slate from which you can edit and continue to create your telegraf config file of your choosing.

```
telegraf -sample-config -input-filter file -output-filter influxdb > file.conf
```

Note in this directory the file **orig.conf** is the output from the above command.

Here is the doc on
[CSV input data format](https://docs.influxdata.com/telegraf/v1.10/data_formats/input/csv/)
with a simple example that works.

### Equity Data Example

This writes data to inflxdb with equity data...

```
telegraf --config ./ubnt.conf
```
