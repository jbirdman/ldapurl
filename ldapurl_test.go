package ldapurl

import "testing"

func TestParseValidURL(t *testing.T) {
	ldapurl, err := Parse("ldaps://ldap.jcu.edu.au/dc=jcu,dc=edu,dc=au")
	if err != nil {
		t.Errorf("Error doing basic URL parse")
	}

	if ldapurl.Scheme != "ldaps" {
		t.Errorf("Scheme not parsed correctly")
	}

	if ldapurl.Host != "ldap.jcu.edu.au" {
		t.Errorf("Host not parsed correctly")
	}

	if ldapurl.Port != DefaultLdapsPort {
		t.Errorf("Port does not have correct default value")
	}

	if !ldapurl.IsTls {
		t.Errorf("Not marked as being a TLS URL")
	}
}

func TestLdapURL_BuildHostnamePortString(t *testing.T) {
	ldapurl, err := Parse("ldaps://ldap.jcu.edu.au/dc=jcu,dc=edu,dc=au")
	if err != nil {
		t.Errorf("Error doing basic URL parse")
	}

	if ldapurl.BuildHostnamePortString() != "ldap.jcu.edu.au:636" {
		t.Errorf("Hostname port string not correct: %s", ldapurl.BuildHostnamePortString())
	}
}