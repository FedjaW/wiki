# Networks

## Break the network
Now that you've seen how you can ping Google successfully, let's quarantine a container and make sure that we can't reach Google.

Run the "getting started" container
Start the getting started container, but do it in `--network none` mode. This removes the network interface from the container.

```shell
docker run -d --network none docker/getting-started
```

Next, run the ping command with a timeout of 2 seconds inside the container:

```shell
docker exec CONTAINER_ID ping google.com -W 2
```

## Load Balancers
Let's try something a bit more complex: let's configure a load balancer!

### What is a load balancer?
A load balancer behaves as advertised: it balances a load of network traffic. Think of a huge website like Google.com. There's no way that a single server (literally a single computer) could handle all of the Google searches for the entire world. Google uses load balancers to route requests to different servers.

### How does a load balancer work?
A central server, called the "load balancer", receives traffic from users (aka clients), then routes those raw requests to different back-end application servers. In the case of Google, this splits the world's traffic across potentially many different thousands of computers.

### Application Servers
First, we need to start some application servers so that we have something to load balance! We'll be using Caddy, an awesome open-source load balancer/web server. Nginx and Apache are other popular alternatives that do similar things, but Caddy is a modern version written in Go, so I think it will be cool to play with.

### What will our application servers do?
Each application server will serve a slightly different HTML webpage. The reason they're different is just so that we can see load balancing in action!

1. Pull down the caddy image

```shell
docker pull caddy
```

2.  Create an index1.html file in your working directory

```html
<html>

<body>
    <h1>Hello from server 1</h1>
</body>

</html>
```

3. Create an index2.html file in your working directory

```html
<html>

<body>
    <h1>Hello from server 2</h1>
</body>

</html>
```

4. Run caddy containers to serve the HTML
Run a container for index1.html on port 8001:

```shell
docker run -p 8001:80 -v $PWD/index1.html:/usr/share/caddy/index.html caddy
```

Run a container for index2.html on port 8002:

```shell
docker run -p 8002:80 -v $PWD/index2.html:/usr/share/caddy/index.html caddy
```

You can run them in separate terminal sessions, or you can run them in detached mode with -d, whichever you prefer.

Navigate to localhost:8001 in a browser. You should see "Hello from server 1". Next, navigate to localhost:8002 and hopefully, you'll see "Hello from server 2"!

## Custom Network

Docker allows us to create custom bridge networks so that our containers can communicate with each other if we want them to, but remain otherwise isolated. We're going to build a system where the application servers are hidden within a custom network, and only our load balancer is exposed to the host.

Let's create a custom bridge network called "caddytest".

```shell
docker network create caddytest
```

Restart your application servers on the network
Stop and restart your caddy application servers, but this time, make sure you attach them to the caddytest network and give them names:

```shell
docker run --network caddytest --name caddy1 -p 8001:80 -v $PWD/index1.html:/usr/share/caddy/index.html caddy
docker run --network caddytest --name caddy2 -p 8002:80 -v $PWD/index2.html:/usr/share/caddy/index.html caddy
```

Contacting the caddy servers through the bridge
To make sure it's working, let's get a shell session inside a "getting started" container on the custom network:

```shell
docker run -it --network caddytest docker/getting-started /bin/sh
```

By giving our containers some names, caddy1 and caddy2, and providing a bridge network, Docker has set up name resolution for us! The container names resolve to the individual containers from all other containers on the network. Within your docker/getting-started container shell, curl the first container:

```shell
curl caddy1
```

```shell
curl caddy2
```

Once you get the HTML responses that you expect, exit out of your shell session within the "getting started" container and then run and submit the tests.

Note that if you need to restart your caddy application servers after naming them, you can use: docker start caddy1 and docker start caddy2.

## Configuring the load balancer
We've confirmed that we have 2 application servers (Caddy) working properly on a custom bridge network. Let's create a load balancer that balances network requests between the two! We'll use a round-robin balancing strategy, so each request should route back and forth between the servers.

### Caddyfile for the loadbalancer
Caddy works great as a file server, which is what our little HTML servers are, but it also works great as a load balancer! To use Caddy as a load balancer we need to create a custom Caddyfile to tell Caddy how we want it to balance traffic.

Create a new file in your local directory called Caddyfile:

```
localhost:80

reverse_proxy caddy1:80 caddy2:80 {
	lb_policy       round_robin
}
```

This tells Caddy to run on localhost:80, and to round robin any incoming traffic to caddy1:80 and caddy2:80. Remember, this only works because we're going to run the loadbalancer within the same network, so caddy1 and caddy2 will automatically resolve to our application server's containers.

### Run the loadbalancer
Instead of providing an index.html to this Caddy server, we're going to give it our custom Caddyfile.

```shell
docker run --network caddytest -p 8080:80 -v $PWD/Caddyfile:/etc/caddy/Caddyfile caddy
```

Now you can hit the load balancer on http://localhost:8080/! You should either get a response from server 1 or server 2, and if you hard refresh the page, it will swap back and forth.

If it's not swapping properly, try using curl instead. Your browser might be caching the HTML.

```shell
curl localhost:8080
```

Each time you run curl, you should get a response from a different server!