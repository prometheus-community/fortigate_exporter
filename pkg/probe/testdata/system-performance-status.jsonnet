# api/v2/monitor/system/performance/status
{
  "http_method":"GET",
  "results": {
    "cpu": {
      "cores": [
        {
          "user": 0,
          "system": 13,
          "nice": 0,
          "idle": 0,
          "iowait": 0
        },
        {
          "user": 1,
          "system": 14,
          "nice": 0,
          "idle": 0,
          "iowait": 0
        },
        {
          "user": 2,
          "system": 0,
          "nice": 0,
          "idle": 0,
          "iowait": 0
        }
      ],
      "user": 200,
      "system": 0,
      "nice": 0,
      "idle": 0,
      "iowait": 0
    },
    "mem": {
      "total": 0,
      "used": 0,
      "free": 0,
      "freeable": 0
    }
  },
  "vdom":"root",
  "path":"system",
  "name":"fortimanager",
  "action":"status",
  "status":"success",
  "serial":"FGT61FT000000000",
  "version":"v6.0.10",
  "build":365
}