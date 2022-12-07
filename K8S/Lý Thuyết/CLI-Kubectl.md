# Command Line Kubectl options

### kubectl get

- List all pods in ps output format

    ` kubectl get pods `

- List all pods in ps output format with more information (such as node name)

    ` kubectl get pods -o wide `

- List a single replication controller with specified NAME in ps output format

    ` kubectl get replicationcontroller web `

### kubectl run

- Start a nginx pod

    ` kubectl run nginx --image=nginx `

- Start a nginx pod and let the container expose port 5701

    ` kubectl run nginx --image=nginx/nginx --port=5701 
    
- Start a busybox pod and keep it in the foreground, don't restart it if it exits

    ` kubectl run -i -t busybox --image=busybox --restart=Never `

- Start the nginx pod using the default command, but use custom arguments (arg1 .. argN) for that command

    ` kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN> `

- Start the nginx pod using a different command and custom arguments

    ` kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN> `