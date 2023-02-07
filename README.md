# Nature Remo Exporter

[![Go](https://github.com/kou64yama/nature-remo-exporter/actions/workflows/go.yml/badge.svg)](https://github.com/kou64yama/nature-remo-exporter/actions/workflows/go.yml)
[![Docker](https://github.com/kou64yama/nature-remo-exporter/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/kou64yama/nature-remo-exporter/actions/workflows/docker-publish.yml)

Prometheus exporter for [Nature Remo](https://nature.global/nature-remo/).

## Getting started

Create an access token at https://home.nature.global.

Run the following command to start nature-remo-exporter.

```bash
echo -n 'Nature Access Token: ' >&2; read -s NATURE_ACCESS_TOKEN; echo >&2
docker run --name=nature-remo-exporter -d --rm -p 8080:8080 \
  -e NATURE_ACCESS_TOKEN="$NATURE_ACCESS_TOKEN" \
  ghcr.io/kou64yama/nature-remo-exporter:main -h 0.0.0.0
```

Open http://localhost:8080/metrics.

## Metrics

### `natureremo_temperature{name,firmware_version,mac_address,serial_number}` (gauge)

This is the value of the temperature sensor, and the unit is "℃".

### `natureremo_humidity{name,firmware_version,mac_address,serial_number}` (gauge)

This is the value of the humidity sendor and ranges from 0 to 100.

### `natureremo_illumination{name,firmware_version,mac_address,serial_number}` (gauge)

This is the value of the illuminance sensor and ranges from 0 to 200.

### `natureremo_movement{name,firmware_version,mac_address,serial_number}` (gauge)

This is the value of the motion sensor. **This value is always 1.**

https://swagger.nature.global/#/default/get_1_devices

> The val of "mo" is always 1 and when movement event is captured created_at is updated.

## References

- [Writing exporters | Prometheus](https://prometheus.io/docs/instrumenting/writing_exporters/)
- [Nature Inc. | Nature Developer Page](https://developer.nature.global)
