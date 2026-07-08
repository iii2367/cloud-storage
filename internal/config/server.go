package config

type Server struct {
	Host	string
    Port	string
}

func MustLoadServer(prefix string) *Server {
	return &Server {
		Host: mustEnv(prefix + "_SERVER_HOST"),
		Port: mustEnv(prefix + "_SERVER_PORT"),
	}   	
}
