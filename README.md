# Nature Remo Exporter

Prometheus exporter for [Nature Remo](https://nature.global/nature-remo/).

## Getting started

Create an access token at https://home.nature.global and create `.env` file.

```bash
echo -n 'Token: ' >&2; read -s; echo >&2
echo NATURE_ACCESS_TOKEN="$REPLY" > .env
```

Start all services.

```bash
docker compose up -d
```

Open http://localhost:3030 in your browser to access Grafana (user/pass is `admin/admin`).

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
