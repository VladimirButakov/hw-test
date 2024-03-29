package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *Client) Close() (err error) {
	err = c.conn.Close()
	return
}

func (c *Client) Send() error {
	if _, err := io.Copy(c.conn, c.in); err != nil {
		return fmt.Errorf("error occurred while sending: %w", err)
	}

	if _, err := fmt.Fprintln(os.Stderr, "...EOF"); err != nil {
		return err
	}

	return nil
}

func (c *Client) Receive() error {
	if _, err := io.Copy(c.out, c.conn); err != nil {
		return fmt.Errorf("error occurred while receiving: %w", err)
	}
	if _, err := fmt.Fprintln(os.Stderr, "...Connection was closed by peer"); err != nil {
		return err
	}

	return nil
}

func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	c.conn = conn

	return nil
}
