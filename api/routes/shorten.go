package routes

import "time"

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Exipiry     time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"custom_short"`
	Exipiry         time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"x_rate_remaining"`
	XRateLimitReset time.Duration `json:"x_rate_limit_reset"`
}
