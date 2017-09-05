package comm;

type Cmd_type int;
const (
	Cmd_type_tcp Cmd_type = iota
	Cmd_type_udp
	Cmd_type_http
)

type Cmd_http_req struct{
	s_cmd_type Cmd_type
	s_cmd_req

}

type Cmd_http_request struct{
	s_cmd_type Cmd_type
	s_cmd_req

}


type Cmd_tcp struct{
	s_cmd_type Cmd_type

} 

type Cmd_udp struct{
	s_cmd_type Cmd_type

} 

