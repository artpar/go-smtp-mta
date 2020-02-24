package mta

import (
	"crypto/tls"
	"errors"
	"io"
	"net"
	"strings"

	smtp "github.com/emersion/go-smtp"
)

type Sender struct {
	Hostname string
}

func (s *Sender) Send(from string, to []string, r io.Reader) error {
	// TODO: buffer r if sending to multiple recipients
	// TODO: group recipients with same domain

	for _, addr := range to {
		_, domain, err := splitAddress(addr)
		if err != nil {
			return err
		}

		mxs, err := net.LookupMX(domain)
		if err != nil {
			return err
		}
		if len(mxs) == 0 {
			mxs = []*net.MX{{Host: domain}}
		}

		for _, mx := range mxs {
		   err := smtp.SendMail(mx.Host + ":25", nil, from, addr, c.Data())
			if err != nil {
				break;
			} else {
				return err	
			}
		}
	}

	return nil
}

func splitAddress(addr string) (local, domain string, err error) {
	parts := strings.SplitN(addr, "@", 2)
	if len(parts) != 2 {
		return "", "", errors.New("mta: invalid mail address")
	}
	return parts[0], parts[1], nil
}
