package ldapurl

import (
	"net/url"
	"errors"
	"fmt"
)

const (
	DefaultLdapPort = 389
	DefaultLdapsPort = 636
)

type LdapURL struct {
	Scheme string
	Host   string
	Port   int
	IsTls  bool
}

func Parse(rawurl string) (ldapurl *LdapURL, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return
	}

	// Start building the object
	ldapurl = &LdapURL{Scheme:u.Scheme, Host:u.Host, }

	// Check for supported schemes and set port defaults and TLS status appropriately
	switch u.Scheme {
	case "ldap":
		ldapurl.IsTls = false
		ldapurl.Port = DefaultLdapPort
		break
	case "ldaps":
		ldapurl.IsTls = true
		ldapurl.Port = DefaultLdapsPort
		break
	default:
		err = errors.New(fmt.Sprintf("Unsupported LDAP URL scheme: %s", u.Scheme))
		return
	}

	return
}

func (ldapurl LdapURL) BuildHostnamePortString() (hostname string) {
	hostname = fmt.Sprintf("%s:%d", ldapurl.Host, ldapurl.Port)
	return
}