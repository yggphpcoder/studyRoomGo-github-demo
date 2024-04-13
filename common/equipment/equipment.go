package equipment

var Servers = make(map[string]Equip)

// Server is transport server.
type Equip interface {
	GetKey() string
	Setup()
	Close()
	Automate(mac string, code string, status int) (ok bool, err error)
	AutomateMain(mac string, code []string, status int) (ok bool, err error)
}

func NewServer(srvs ...Equip) {

	for _, srv := range srvs {
		key := srv.GetKey()
		srv.Setup()
		Servers[key] = srv
	}
}

func Close() {
	for _, srv := range Servers {
		srv.Close()
	}
}
