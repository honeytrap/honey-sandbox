package gvisor

/*

// containerConn defines a custom connection type which proxies the data
// for the container.
type containerConn struct {
	net.Conn
	container *gvisorContainer
}

// Read reads the giving set of data from the container connection to the
// byte slice.
func (c containerConn) Read(b []byte) (n int, err error) {
	c.container.stillActive()
	return c.Conn.Read(b)
}

// Write writes the data into byte slice from the container.
func (c containerConn) Write(b []byte) (n int, err error) {
	c.container.stillActive()
	return c.Conn.Write(b)
}

// stillActive returns an error if the containerr is not still active
func (c *containerConn) stillActive() error {
	if c.isStopped() {
		return fmt.Errorf("container not running %s", c.name)
	}
	if c.isFrozen() {
		return c.unfreeze()
	}
	if !c.isRunning() {
		return fmt.Errorf("container in unknown state %s:%s", c.name, c.c.State())
	}
	c.idle = time.Now()
	return nil
}

*/
