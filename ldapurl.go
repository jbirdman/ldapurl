package ldapurl

type LdapURL struct {
	Scheme string
	Host   string
	Port   int
}

func (u *LdapURL) Parse(ldapUrl string) {

}
