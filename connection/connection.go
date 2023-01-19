package connection

type Connection struct {
	group    string
	name     string
	role     string
	ip       string
	port     int
	password string
}

func New(group string, name string, role string, ip string, port int, password string) *Connection {
	return &Connection{group: group, name: name, role: role, ip: ip, port: port, password: password}
}

func (c *Connection) Group() string {
	return c.group
}

func (c *Connection) SetGroup(group string) {
	c.group = group
}

func (c *Connection) Name() string {
	return c.name
}

func (c *Connection) SetName(name string) {
	c.name = name
}

func (c *Connection) Role() string {
	return c.role
}

func (c *Connection) SetRole(role string) {
	c.role = role
}

func (c *Connection) Ip() string {
	return c.ip
}

func (c *Connection) SetIp(ip string) {
	c.ip = ip
}

func (c *Connection) Port() int {
	return c.port
}

func (c *Connection) SetPort(port int) {
	c.port = port
}

func (c *Connection) Password() string {
	return c.password
}

func (c *Connection) SetPassword(password string) {
	c.password = password
}
