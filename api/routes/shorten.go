package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

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

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body) ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse json"})
	}

	//impeliment rate limititng
	

	//check if the err is an actual URL

	if !govalidate.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid URL"})
	}

	//check for a domain error
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"domain"})
	}

	//enforce http, SSL

	body.URL = helpers.EnforceHTTP(body.URL)

	return nil
}