package controllers

import (
	"crypto/tls"
	"fmt"
	"time"
)

type CertificateData struct {
	Host          string `json:"host"`
	ExpiryDate    string `json:"expiry_date"`
	RemainingDays int    `json:"remaining_days"`
}

// getConnectionState returns basic TLS details about the connection
func getConnectionState(host string) (tls.ConnectionState, error) {
	// Request to host
	conn, err := tls.Dial("tcp", fmt.Sprint(host+":443"), nil)
	if err != nil {
		return tls.ConnectionState{}, err
	}
	defer conn.Close()

	// VerifyHostname checks that the peer certificate chain is valid
	err = conn.VerifyHostname(host)
	if err != nil {
		return tls.ConnectionState{}, err
	}

	return conn.ConnectionState(), nil
}

//SSLChecker function that check SSL certificate
func SSLChecker(host string) (CertificateData, error) {

	connectionState, err := getConnectionState(host)
	if err != nil {
		return CertificateData{}, err
	}

	// Get expiration date to host
	expiryDate := connectionState.PeerCertificates[0].NotAfter
	// IsZero reports whether expiry represents the zero time instant,
	// January 1, year 1, 00:00:00 UTC.
	if expiryDate.IsZero() {
		return CertificateData{}, fmt.Errorf("expiry date is a zero time")
	}

	// Get system current date
	now := time.Now()
	// Difference in days between the expiration date of the certificate and
	// the date of the system at the time of verification
	remainingDays := int(expiryDate.Sub(now).Hours() / 24)

	return CertificateData{
		Host:          host,
		ExpiryDate:    expiryDate.Format("02/Jan/2006"),
		RemainingDays: remainingDays}, nil
}
