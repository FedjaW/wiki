# Start command in container

When you start a container it runs only as long as the process in the container is running. So if a webserver or so in the container crashed the container stops running.

_Who defines what process is run within the container?_

In dockerfiles a command is defined at the bottom of the file e.g. `CMD ["nginx]`. This command defines the programm that is run when the container starts.

Example:

```CLI
docker run ubuntu
```

This will terminate directly after start, because the default command in the dockerfile of the ubuntu image is `CMD ["bash]`. By default docker does not attach a terminal onto a container when it is run and so the bash programm does not find a terminal and so it exits. Since the process that was started when the container was created finished the container exits as well.

_How to specify to a different command to start the container?_

## CMD

One option is to append a command to the docker run command. In that way it overrides the default command specified within the image. In our example:

```CLI
docker run ubuntu [COMMAND]
```

```CLI
docker run ubuntu sleep 5
```

Now the container will run the sleep process and exit after 5 seconds.

_How to make this change permanently to run the sleep command always when the container starts?_

You would create your own image based on the ubuntu image and specify a new command:

```DOCKERFILE
FROM Ubuntu

CMD sleep 5
```

There are different ways to specify the command:

```DOCKERFILE
CMD command param1 -> CMD sleep 5
```

```DOCKERFILE
CMD ["command", "param1"] -> CMD ["sleep", "5"]
```

`DO NOT` specify the command and the paramters together like this: CMD ["sleep 5"]

Now, we build the new image with the new command:

```CLI
docker build -t ubuntu-sleeper .
```

```CLI
docker run ubuntu-sleeper
```

Now, the sleep command is executed always, when the container is created.

_How to change the number of seconds it sleeps?_

We learned that we just can run: `docker run ubuntu-sleeper sleep 10`

But this does not look very good, because the ubuntu-sleeper already implies that is will sleep, it should be no need to write sleep again. It should look like: `docker run ubuntu-sleeper 10`

## ENTRYPOINT

Here is where the `ENTRYPOINT` comes into play.

```DOCKERFILE
FROM Ubuntu

ENTRYPOINT ["sleep"]
```

The `ENTRYPOINT` instruction is like the `CMD` instruction as you can specify the programm that will be run when the container starts. And what ever you specify in the command line, in this case `10`, will be appended to the `ENTRYPOINT`.

```CLI
docker run ubuntu-sleeper 10
```

Command at startup: `sleep 10`

So the difference between `ENTRYPOINT` and `CMD` is that for the `CMD` case the command line parameters will replace the default command set in the docker file entirely. Where as in case of the `ENTRYPOINT` the command line parameters will be appended to the default command set in the docker file.

_So what will happen when I run the docker run command without specifiyng a parameter (in case of `ENTRYPOINT`):_

```CLI
docker run ubuntu-sleeper
```

Then the command at startup will be only `sleep` and you will get the error that the operand is missing.

## ENTRYPOINT and CMD together

That is why you would use both the `ENTRYPOINT` and the `CMD`:

```DOCKERFILE
FROM Ubuntu

ENTRYPOINT ["sleep"]

CMD ["5"]
```

In this case the `CMD` instruction will be appended to the `ENTRYPOINT` instruction. So at startup the command is: `sleep 5` if you did not specify any parameters in the command line. If you did, that will override the `CMD` instruction.

```CLI
docker run ubuntu-sleeper 10
```

Command at startup -> `sleep 10`

In order to override the `ENTRYPOINT` instruction you can run:

```CLI
docker run --entrypoint newsleep2.0 ubuntu-sleeper 10
```

Command at startup -> `newsleep2.0 10`

## Commands and arguments in the pod-definition.yaml

How do `CMD` and `ENTRYPOINT` translate into a pod-definition.yaml file.

```YAML
apiVersion: v1
kind: Pod
metadata:
  name: ubuntu-sleeper-pod
spec:
  containers:
    - name: ubuntu-sleeper
      image: ubuntu-sleeper
      command: ["newsleep.2"] # -> corresponds to Entrypoint in dockerfile
      args: ["10"] # -> corresponds to CMD in dockerfile
```

To create the pod:

```CLI
kubectl create -f pod-definition.yaml
```

We override the `CMD` instruction in the dockerfile with the args.

```DOCKERFILE
FROM Ubuntu

ENTRYPOINT ["sleep"]

CMD ["5"]
```

To override the `ENTRYPOINT` use the command field in the yaml defintion.
