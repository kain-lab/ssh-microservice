package manage

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"net"
	"ssh-microservice/app/schema"
	"ssh-microservice/app/types"
	"ssh-microservice/app/utils"
)

type ClientManager struct {
	options       map[string]*types.SshOption
	tunnels       map[string]*[]types.TunnelOption
	runtime       map[string]*ssh.Client
	localListener map[string]map[string]*net.Listener
	localConn     *utils.SyncMapConn
	remoteConn    *utils.SyncMapConn
	schema        *schema.Schema
}

func NewClientManager() *ClientManager {
	c := new(ClientManager)
	c.options = make(map[string]*types.SshOption)
	c.tunnels = make(map[string]*[]types.TunnelOption)
	c.runtime = make(map[string]*ssh.Client)
	c.localListener = make(map[string]map[string]*net.Listener)
	c.localConn = utils.NewSyncMapConn()
	c.remoteConn = utils.NewSyncMapConn()
	c.schema = schema.New()
	return c
}

func (c *ClientManager) empty(identity string) error {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		return errors.New("this identity does not exists")
	}
	return nil
}

func (c *ClientManager) GetIdentityCollection() []string {
	var keys []string
	for key := range c.options {
		keys = append(keys, key)
	}
	return keys
}

// Get ssh client information
func (c *ClientManager) GetSshOption(identity string) (option *types.SshOption, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	option = c.options[identity]
	return
}

func (c *ClientManager) GetRuntime(identity string) (client *ssh.Client, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	client = c.runtime[identity]
	return
}

func (c *ClientManager) GetTunnelOption(identity string) (option []types.TunnelOption, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		option = *c.tunnels[identity]
	}
	return
}
