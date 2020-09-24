package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "time"
    "sort"

    "github.com/urfave/cli"
)


func main() {
    app := cli.NewApp()     //creating new app
    app.Name = "Lookup CLI"  //providing name to the app
    app.Usage = "Use to query IPs, CNAMEs, MX records!"
    app.Version = "1.01.001"  // current version
    app.Compiled = time.Now()

    // author details
    author := cli.Author{   
        Name:  "Deepak Singh Rathore",
        Email: "dsrdeepak8@gmail.com",
      }
    app.Authors = []*cli.Author{
        &author,
      }
  
    flag1 := cli.StringFlag{
        Name:"host, ",  
    }
    flagAddr := cli.StringFlag{
        Name:"addr, ",
    }
    myFlags := []cli.Flag{
       &flag1,
       &flagAddr,
    }


    app.Commands = []*cli.Command{
        {
            Name:  "addr",
            Usage: "Performs a reverse lookup for the given address, returning a list of names mapping to that address.",
            Flags: myFlags,
            // the action, or code that will be executed when
            // we execute our `addr` command
            Action: func(c *cli.Context) error {
                addr, err := net.LookupAddr(c.String("addr"))
                if err != nil {
                    fmt.Println(err)
                }
                for i := 0; i < len(addr); i++ {
                    fmt.Println(addr[i])
                }
                return nil
            },
        },
        {
            Name:  "host",
            Usage: "Looks up the given host using the local resolver. It returns a slice of that host's addresses.",
            Flags: myFlags,
            // the action, or code that will be executed when
            // we execute our `host` command
            Action: func(c *cli.Context) error {
                addr, err := net.LookupAddr(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                }
                for i := 0; i < len(addr); i++ {
                    fmt.Println(addr[i])
                }
                return nil
            },
        },
        {
            Name:  "ns",
            Usage: "Looks Up the NameServers for a Particular Host",
            Flags: myFlags,
            // the action, or code that will be executed when
            // we execute our `ns` command
            Action: func(c *cli.Context) error {
                // a simple lookup function
                ns, err := net.LookupNS(c.String("host"))
                if err != nil {
                    return err
                }
                // we log the results to our console
                // using a trusty fmt.Println statement
                for i := 0; i < len(ns); i++ {
                    fmt.Println(ns[i].Host)
                }
                return nil
            },
        },
        {
            Name:  "ip",
            Usage: "Looks up the IP addresses for a particular host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                ip, err := net.LookupIP(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                }
                for i := 0; i < len(ip); i++ {
                    fmt.Println(ip[i])
                }
                return nil
            },
        },
        {
            Name:  "cname",
            Usage: "Looks up the CNAME for a particular host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                cname, err := net.LookupCNAME(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println(cname)
                return nil
            },
        },
        {
            Name:  "mx",
            Usage: "Looks up the MX records for a particular host",
            Flags: myFlags,
            Action: func(c *cli.Context) error {
                mx, err := net.LookupMX(c.String("host"))
                if err != nil {
                    fmt.Println(err)
                }
                for i := 0; i < len(mx); i++ {
                    fmt.Println(mx[i].Host, mx[i].Pref)
                }
                return nil
            },
        },
    }
    sort.Sort(cli.CommandsByName(app.Commands))
    
    // start our application
    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}