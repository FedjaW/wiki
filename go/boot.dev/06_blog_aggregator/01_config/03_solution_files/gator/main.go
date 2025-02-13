package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/fedjaw/gator/internal/config"
)

func main() {
    cfg, _ := config.Read()

    s := state{
        cfg: &cfg,
    }

    cmds := commands{
        cmds: make(map[string]func(*state, command) error),
    }

    cmds.register("login", handlerLogin)

    args := os.Args
    if len(args) < 2 {
        log.Fatalf("error reading arguments")
    }
    commandName := args[1] // args[0] is the program name
    commandArgs := args[2:]

    cmd := command{
        name: commandName,
        args: commandArgs,
    }

    err := cmds.run(&s ,cmd)
    if err != nil {
        log.Fatalf("%v", err)
    }

    // cfg, _ = config.Read()
    // fmt.Printf("%+v\n", cfg)
}

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return errors.New("no argument given for login")
    }

    username := cmd.args[0]
    err := s.cfg.SetUser(username)
    if err != nil {
        return err
    }
    
    fmt.Printf("user has been set to: %s\n", username)

    return nil
}

type state struct {
    cfg *config.Config
}

type command struct {
    name string
    args []string
}

type commands struct {
    cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
    c.cmds[name] = f
}

func (c *commands) run(state *state, cmd command) error {
    err := c.cmds[cmd.name](state, cmd)

    return err
}
