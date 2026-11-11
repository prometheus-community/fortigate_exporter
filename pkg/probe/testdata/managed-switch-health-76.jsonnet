# api/v2/monitor/switch-controller/managed-switch/health-status?vdom=*
# Synthetic: one switch with a non-good CPU rating and a sensor in an unsupported unit.

[
  {
    "http_method": "GET",
    "results": {
      "FOO-SW-76": {
        "performance-status": {
          "cpu": {
            "user": { "value": 70, "unit": "%%" },
            "system": { "value": 15, "unit": "%%" },
            "nice": { "value": 0, "unit": "%%" },
            "idle": { "value": 15, "unit": "%%" }
          }
        },
        "temperature": [
          { "module": "sensor1(CPU Board Temp)", "status": { "value": 55.5, "unit": "celsius" } },
          { "module": "sensor9(Unknown)", "status": { "value": 1, "unit": "other" } }
        ],
        "summary": {
          "cpu": { "value": 85, "rating": "poor" },
          "memory": { "value": 40, "rating": "good" },
          "uptime": { "value": 3600, "rating": "good" },
          "temperature": { "value": 55.5, "rating": "good" },
          "overall": "poor"
        }
      }
    },
    "vdom": "root",
    "path": "switch-controller",
    "name": "managed-switch",
    "action": "health-status",
    "status": "success",
    "serial": "FGVM000000000000",
    "version": "v7.6.0",
    "build": 3401
  }
]
