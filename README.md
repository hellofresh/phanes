<p align="center">
  <a href="https://hellofresh.com">
    <img width="120" src="https://www.hellofresh.de/images/hellofresh/press/HelloFresh_Logo.png">
  </a>
</p>

# Phanes

[![Go Report Card](https://goreportcard.com/badge/github.com/hellofresh/phanes)](https://goreportcard.com/report/github.com/hellofresh/phanes)

> A CLI Tool to create oauth clients

This is a simple, CLI tool that helps you to create oauth clients for many different identity providers.
Through this tool you can create and delete clients, you can also hook this up in any web application that you wish.
At the moment we support the `hellofresh`, `google` and `facebook` identity providers.

## Why Phanes?

> Phanes (Ancient Greek: Φάνης, genitive Φάνητος), or Protogonos (Greek: Πρωτογόνος, "First-born"), 
was the mystic primeval deity of procreation and the generation of new life, who was introduced into Greek mythology by 
the Orphic tradition. [Wikipedia](https://en.wikipedia.org/wiki/Phanes_(mythology))

Since this was the god of creation, why not naming a creation tool after it?

## What is a Client *Application*?

The client is the application that wants to access the user's account. Before it may do so, it must be authorized by the user, 
and the authorization must be validated by the API.

## Installation

You can get the binary and play with it in your own enviroment (or even deploy it whereever you like it).
Just go the [releases](https://github.com/hellofresh/phanes/releases) and download the latest one for your platform.

Just place the binary in your $PATH and you are good to go.

## Getting Started

After you have *phanes* up and running we can create our first client. Let's see how we can create a new client:

```sh
phanes -u "http://localhost:8000/clients/" -p "hellofresh" create -n "Test 1"
```

This command should return something like:

```sh
Credentials for Example created!
Client ID: 69278640-eff2-4bc3-b6ef-a12c86810324
Client Secret: 03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4
```

If you want to delete a client you simply need to call

```sh
phanes delete -p "hellofresh" -id 69278640-eff2-4bc3-b6ef-a12c86810324
```

## Contributing

To start contributing, please check [CONTRIBUTING](CONTRIBUTING.md).

## Documentation

* Phanes Docs: https://godoc.org/github.com/hellofresh/phanes
* Go lang: https://golang.org/
