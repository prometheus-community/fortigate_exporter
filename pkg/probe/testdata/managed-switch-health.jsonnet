# api/v2/monitor/switch-controller/managed-switch/health?vdom=*

[
  {
    "http_method": "GET",
    "results": {
      "FS00000000000027": {
        "performance-status": {
          "cpu": {
            "user": {
              "value": 0,
              "unit": "%%"
            },
            "system": {
              "value": 0,
              "unit": "%%"
            },
            "nice": {
              "value": 0,
              "unit": "%%"
            },
            "idle": {
              "value": 100,
              "unit": "%%"
            }
          },
          "memory": {
            "used": {
              "value": 15,
              "unit": "%%"
            }
          },
          "uptime": {
            "days": {
              "value": 454,
              "unit": "days"
            },
            "hours": {
              "value": 17,
              "unit": "hours"
            },
            "minutes": {
              "value": 49,
              "unit": "minutes"
            }
          }
        },
        "temperature": [
          {
            "module": "sensor1(CPU  Board Temp)",
            "status": {
              "value": 39.0,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor2(MAIN Board Temp1)",
            "status": {
              "value": 60.625,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor3(MAIN Board Temp2)",
            "status": {
              "value": 48.937,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor4(MAIN Board Temp3)",
            "status": {
              "value": 36.062,
              "unit": "celsius"
            }
          }
        ],
        "summary": {
          "cpu": {
            "value": 0,
            "rating": "good"
          },
          "memory": {
            "value": 15,
            "rating": "good"
          },
          "uptime": {
            "value": 39289740,
            "rating": "good"
          },
          "temperature": {
            "value": 46.156000000000006,
            "rating": "good"
          },
          "overall": "good"
        }
      },
      "FS00000000000024": {
        "performance-status": {
          "cpu": {
            "user": {
              "value": 0,
              "unit": "%%"
            },
            "system": {
              "value": 0,
              "unit": "%%"
            },
            "nice": {
              "value": 0,
              "unit": "%%"
            },
            "idle": {
              "value": 100,
              "unit": "%%"
            }
          },
          "memory": {
            "used": {
              "value": 10,
              "unit": "%%"
            }
          },
          "uptime": {
            "days": {
              "value": 454,
              "unit": "days"
            },
            "hours": {
              "value": 17,
              "unit": "hours"
            },
            "minutes": {
              "value": 48,
              "unit": "minutes"
            }
          }
        },
        "temperature": [
          {
            "module": "sensor1(CPU  Board Temp)",
            "status": {
              "value": 41.937,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor2(MAIN Board Temp1)",
            "status": {
              "value": 63.875,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor3(MAIN Board Temp2)",
            "status": {
              "value": 51.312,
              "unit": "celsius"
            }
          },
          {
            "module": "sensor4(MAIN Board Temp3)",
            "status": {
              "value": 38.687,
              "unit": "celsius"
            }
          }
        ],
        "summary": {
          "cpu": {
            "value": 0,
            "rating": "good"
          },
          "memory": {
            "value": 10,
            "rating": "good"
          },
          "uptime": {
            "value": 39289680,
            "rating": "good"
          },
          "temperature": {
            "value": 48.952749999999995,
            "rating": "good"
          },
          "overall": "good"
        }
      }
    },
    "vdom": "root",
    "path": "switch-controller",
    "name": "managed-switch",
    "action": "health",
    "status": "success",
    "serial": "FS000000000000302",
    "version": "v6.4.0",
    "build": 1030
  }
]
