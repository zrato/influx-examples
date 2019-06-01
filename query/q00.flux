from(bucket: "rick")
  |> range(start: -150m)
  |> filter(fn: (r) => r._measurement == "origxa")
