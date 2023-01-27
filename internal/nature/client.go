package nature

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SensorValue struct {
	Value     float64   `json:"val"`
	CreatedAt time.Time `json:"created_at"`
}

type Device struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	TemperatureOffset float64   `json:"temperature_offset"`
	HumidityOffset    float64   `json:"humidity_offset"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	FirmwareVersion   string    `json:"firmware_version"`
	MacAddress        string    `json:"mac_address"`
	SerialNumber      string    `json:"serial_number"`
	NewestEvents      struct {
		Temperature  SensorValue `json:"te"`
		Humidity     SensorValue `json:"hu"`
		Illumination SensorValue `json:"il"`
		Movement     SensorValue `json:"mo"`
	} `json:"newest_events"`
}

type clientOptions struct {
	accessToken string
}

type ClientOption interface {
	apply(opts *clientOptions)
}

type clientOptionImpl struct {
	f func(opts *clientOptions)
}

func (i *clientOptionImpl) apply(opts *clientOptions) {
	i.f(opts)
}

func AccessToken(s string) ClientOption {
	return &clientOptionImpl{
		f: func(opts *clientOptions) {
			opts.accessToken = s
		},
	}
}

type Client interface {
	GetDevices() ([]Device, error)
}

type clientImpl struct {
	httpClient  *http.Client
	accessToken string
}

func NewClient(opts ...ClientOption) Client {
	o := &clientOptions{}
	for _, opt := range opts {
		opt.apply(o)
	}

	return &clientImpl{
		httpClient:  &http.Client{},
		accessToken: o.accessToken,
	}
}

func (i *clientImpl) GetDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", "https://api.nature.global/1/devices", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+i.accessToken)

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to fetch devices: %s", resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	devices := []Device{}
	if err := dec.Decode(&devices); err != nil {
		return nil, err
	}

	return devices, nil
}
