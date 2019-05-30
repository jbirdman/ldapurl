package ldapurl

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

const (
	DefaultLdapPort  = 389
	DefaultLdapsPort = 636
)

type LdapURL struct {
	Scheme     string
	Host       string
	Port       int
	DN         string
	Attributes []string
	Scope      string
	Filter     string
	Extensions []string
}

func SplitHostPort(hostport string, defaultport int) (host string, port int) {
	host, portstring, err := net.SplitHostPort(hostport)
	if err != nil {
		port = defaultport
		host = hostport
		return
	}

	// Need to convert string port to int
	port, err = strconv.Atoi(portstring)
	if err != nil {
		port = defaultport
	}

	return
}

func Parse(rawurl string) (ldapurl *LdapURL, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return
	}

	host, port := SplitHostPort(u.Host, 0)

	// Start building the object
	ldapurl = &LdapURL{Scheme: u.Scheme, Host: host}

	// Check for supported schemes and set port defaults and TLS status appropriately
	switch u.Scheme {
	case "ldap":
		if ldapurl.Port = port; port == 0 {
			ldapurl.Port = DefaultLdapPort
		}
		break
	case "ldaps":
		if ldapurl.Port = port; port == 0 {
			ldapurl.Port = DefaultLdapsPort
		}
		break
	default:
		err = errors.New(fmt.Sprintf("Unsupported LDAP URL scheme: %s", u.Scheme))
		return
	}

	// DN part of the URL
	ldapurl.DN = strings.TrimPrefix(u.Path, "/")

	parts := strings.Split(u.RawQuery, "?")

	for i, v := range parts {
		switch i {
		case 0:
			ldapurl.Attributes = strings.Split(v,",")
		case 1:
			ldapurl.Scope = v
		case 2:
			ldapurl.Filter = v
		case 3:
			ldapurl.Extensions = strings.Split(v,",")
		}
	}

	return
}

func (ldapurl LdapURL) BuildHostnamePortString() (hostname string) {
	hostname = fmt.Sprintf("%s:%d", ldapurl.Host, ldapurl.Port)
	return
}

func (ldapurl LdapURL) IsTLS() bool {
	return ldapurl.Scheme == "ldaps"
}
