# m∞

Moo is an engine for machine learning workflows that runs on your local machine or in a Kubernetes cluster. M stands for machine learning on which we are working and ∞ for nothing but being 🆒.

## Architecture

![moo.jpg](https://s2.loli.net/2022/07/11/WbBHm3RuVX8Ik7d.jpg)

Here are some termiologies you need to know before diving in.

### Builder

The builder is used to turn sources into bundles which are then delivered to the runtime that runs each bundle in one or several pods. The sources could be a series of files retrieved from a remote Git repository. A bundle is an OCI image when Moo is running in a Kubernetes cluster, while in your local machine it stays unchanged but with dependencies installed in an isolated environment.

### Pod

In a Kubernetes cluster, a pod is an instance that has several containers running in it. Here for Moo, we reuse the concept but only when running in a k8s or k3s cluster. When in a local machine, a pod refers to an isolated environment created by the Moo builder. In the environment runs the actual process, an inferencing service for instance. The way to isolate environments depends on which driver the Moo runtime uses. We'll explain that later.

### Runtime

Runtime is the magic behind the Moo Engine, it maintains all the pods. For example, the runtime will restart those dead pods as long as it does not exceed the maximum retries of the pod.

In a Kubernetes cluster, the runtime simply calls the rest API given by Kubernetes to create or delete pods or list them by conditions. While in your local machine, Moo uses some drivers implemented [here](./runtime/local/driver/) to make those approaches and each driver needs to make an isolated environment for a pod. The simplest implementation Moo uses now is the [conda driver](./runtime/local/driver/conda/). ~~We will try to make a better implementation based directly on the cgroup and the namespace technique of Linux.~~ You can use the [Podman driver](./runtime/local/driver/podman/) instead.

### Router

The router as we know it maintains some routes to specific resources. Here Moo router maps each pod to its addresses, so we know where to call the stuff running in a pod. You probably realized we need some load balancing techniques here since a pod may have more than just one address. We're gonna talk about this later.

### Gateway

Since we build this for sharing machine learning model services, we need an entry point for users to access. This is where an API gateway, or gateway for short, comes into play. Here's what the Moo gateway does when receiving an external request: First it looks up in the router the available addresses of the pod being accessed and then selects one to which it then forwards the request. Yep, it's how the reversed proxy works.

### Server

Like the Kubernetes api-server, the Moo server handles all the internal requests that could be sent from a system administrator using the Moo CLI or from the *Moo Cloud* services which we will be talking about in the second phase in the near future. 

What's amazing is that this is how the Moo gateway knows pod addresses from the Moo router.

### Client

The Moo Client is just an encapsulation of some HTTP calls to the Moo server that can be easily used to implement things like the Moo CLI or  some SDKs.

### Preset

Preset is not shown in the architecture graph but it's important. It is a design pattern learned from [micro](https://github.com/micro/micro), the microservice platform of the future, to initialize the components we have talked about above when starting up the Moo engine. By now, we have three preset defined as follows.

- **local** for running Moo in your own local machine
- **kubernetes** for running Moo in a k8s/k3s cluster
- **test** for testing purposes with some pseudo implementations

Each of which can be used in the relevant scenarios.

## Roadmap

- Moo Engine (now)
- Moo Cloud Platform
- Moo Edge Network

## License

Moo is Apache 2.0 licensed.

