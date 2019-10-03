## Kubernetes

K8s 创建一个 Cluster, K8s 是一个容器编排工具。 Docker 自身也有一个容器编排工具，Docker 中的 Dockerfile 是针对单容器的，而 Docker-compose 是针对多容器的，而 docker swarm 是对容器进行编排的。而 K8s 就是一个容器编排工具。

K8s, Kubernetes 源自希腊语，其意思是船长、舵手之义。由于微服务架构，现在业内的一种趋势是 DevOps，即开发与运维结合。DevOps 是一种文化，一种趋势。

- CI: 持续集成
- CD: 持续交付。
- CD:持续布署。

K8s，2014 年发布的。基特点有：

1. 自动装箱，自我修复，水平扩展，服务发现和负载均衡，自动发布和回滚。
2. 密钥和配置管理（配置不在本地获取，而是通过配置中心服务器）
3. 存储编排，批量处理执行。

mysql : master / nodes 主从模式，由 Master 来指派任务让 node 执行，而 Master 作冗余来保证高可用。而 redis 是 node/node 模式。

K8s 也是 master / node 模式，高可用。只用  Master 才能接收请求，然后通过 Scheduler 在 node 之间进行调度分配任务。K8s 只是一个 API Server, 而其它客户端能过 api 调用的方式来请求，如 `kubectl`. 而且 K8s 的 API 是 RESTful 的 API，即所有的内容都是资源，可以用 HTTP 的 methods 来获取。

K8s, 由一个 API Server, 一个调度器 Scheduler, 还有一个高可用探测器进行探测容器是否健康运行，而 K8s 有自愈能力，所以还有一个高可探测器的管理器。而  Master 自身会做冗余，以保证高可用。

K8s 调试的最小单位是 pod。pod 是一个逻辑组件，是对容器的封装， 一个 pod 内可以有多个容器，它们在同一个 pod 内可以共享资源。

ELK: 一般的模式是主容器运行，sidecar 辅助主容器做一些其它功能，如 log 容器，主要做对主容器的日志收集，还有其它的 sidecar 容器。

node 是用来执行 master 指派的功能的，因为重启后 id 会变化，不能用 id 作为 node 的标识，所以我们用 标签 Label 来做标识。而 Label 也有选择品，Label Selector, 简称为 Selector。

Node 上的 kubelet 用来在 node 上接收请求和执行任务。

### K8s

- K8s 是一个集群， cluster.
- master/ node
  - master: API Server, Scheduler, controller-manager
  - node: kubelet, docker, kube-proxy

- pod(逻辑组件)： Label, Label Selector.
  - Label: key=value, (metadata, 源数据信息)
  - Label selector.

- Pod :

  自主式 POD, 控制器管理 POD, POD 是有生命周期的。

  - ReplicationController
  - ReplicaSet
  - Deployment(无状态)
  - StatefulSet(有状态)
  - DamonSet
  - Job, CtonJob

  上面不同的 Pod 控制器运行不同的 Pod 资源。

HPA: Horizontal Pod AutoScaler

Pod 挂掉了，会重启一个，怎么访问到那？

客户端不是直接调用 pod 的，而是通过服务层来调用。服务层类似一个总线，pod 在总线上进行注册说明，这们就可以做到服务发现。

Pod 有生命周期，可能随时被替换，所以 Pod 与客户端有一个 Service 中间层。Service 除了发现 Pod 还可以做负载，service 靠 label selector （metadata）来发现 pod, 不同主机名，也不用 IP。Service 是一个代理。

Pod 用 serive 来代理服务，而 Pod 的重启替换，则有一个控制器来管理。

K8s 是一个集群，机房。结点服务器只有一个边界。

pod 内是一个网络，service 是一个网络，而节点本身也是一个网络。

- 同一个 pod 内的多个容器通信用 `lo`, 本地通信。

- 各个 pod 之间的通信用 Overlay Network，叠加网络。Pod 和 pod pp 之间是可以直接通信的。
- pod 与 services 之间通信用 (kube-proxy)

多个 master 的信息用共享存储的 etcd 来存储。etcd是一个开源的、分布式的键值对数据存储系统，提供共享配置、服务的注册和发现。etcd与zookeeper相比算是轻量级系统，两者的一致性协议也一样，etcd的raft比zookeeper的paxos简单。

K8s 系统配置，传统的配置会很复杂。 另一各是将 K8s 的各组件作为 Pod 来才能，如 kubelet, dokcer。这些基础的 pod 称之为 static pod.

Etcd/ zookeeper.

### Kubernetes Pods

A pod is a Kubernetes abstraction that represents a group of one or more application containers(such as Docker or rkt), and some shared resources for those containers. Those resources include:

- Shared storage, as Volumes
- Networking, as a unique cluster IP address
- Information about how to run each container, such as the container image version or specific ports to use.

### Nodes

A Pod always runs on a **Node**. A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master.1

Every Kubernetes Node runs at least:

- Kubelet, a process responsible for communication between the Kubernetes Master and the Node; it manages the Pods and the containers running on a machine.
- A container runtime (like Docker, rkt) responsible for pulling the container image from a registry, unpacking the container, and running the application.

### Overview of Kubernetes Services

Pods 是有生命周期的。

### Tutorial on Mac

#### Installing the Kubernetes Dashboard

```bash
kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml
```

从 `yaml`配置文件中可以看出，我们把 Dashboard 应用程序作为一个 Pod 布署在了 `kube-system` 命名空间内了。

通过下面的命令可以查看命令空间中的 Pods.

```bash
kubectl get
```

#### Create sample user

创建 Kubernetes  Dashboard 的[管理帐号](https://github.com/kubernetes/dashboard/wiki/Creating-sample-user)。

```bash
## 创建 dashboard-adminuser 文件
cat >> dashboard-adminuser.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
# ctrl + d

## 启动
kubectl apply -f dashboard-adminuser.yaml

## 查看 Token, Bearer Token
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}')

```

访问 [`http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/`](http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/)

### Troubleshooting and Debugging Commands 

监测和调试命令。

```bash
Troubleshooting adn Debugging Commands
  describe       Show details of a specific resource or group of resources
  logs           Print the logs for a container in a pod
  attach         Attach to a running container
  exec           Execute a command in a container
  port-forward   Forward one or more local ports to a pod
  proxy          Run a proxy to the Kubernetes API server
  cp             Copy files and directories to and from containers.
  auth           Inspect authorization	
```

#### Advanced Commands

```
Advanced Commands:
  apply          Apply a configuration to a resource by filename or stdin
  patch          Update field(s) of a resource using strategic merge patch
  replace        Replace a resource by filename or stdin
  convert        Convert config files between different API versions
```

#### Deploy Commands

```
Deploy Commands:
  rollout        Manage the rollout of a resource
  rolling-update Perform a rolling update of the given ReplicationController
  scale          Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job
  autoscale      Auto-scale a Deployment, ReplicaSet, or ReplicationController
```

### 删除资源 Resources

在 Kubernetes 中什么才是 Resources 资源那？

```
> kubectl get service
NAME          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
hello-nginx   NodePort    10.100.53.228   <none>        80:30650/TCP   2h
kubernetes    ClusterIP   10.96.0.1       <none>        443/TCP        3h
mongo         ClusterIP   None            <none>        27017/TCP      1h
> kubectl delete
error: You must provide one or more resources by argument or filename.
Example resource specifications include:
   '-f rsrc.yaml'
   '--filename=rsrc.json'
   '<resource> <name>'
   '<resource>'
kubernetes-tutorial ⍉ ➜ kubectl delete service mongo
service "mongo" deleted
kubernetes-tutorial ➜ kubectl delete service hello-nginx
service "hello-nginx" deleted
kubernetes-tutorial ➜ kubectl get service
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   3h
```

### Running MongoDB on Kubernates

| IP            | Root密码  | MySQL Server | Web Server | PHP版本 |
| ------------- | --------- | ------------ | ---------- | ------- |
| 10.12.236.117 | apppb@dev | 无           | Apache     | 7.0.8   |      |
| 10.12.234.96  | apppb@dev | 5.6.25       | Apache     | 5.4.43  |      |
| 10.12.236.22  | apppb@dev |              | Nginx      | 7.0.9   |      |
| 10.123.8.23   | apppb@dev | 无           | Nginx      | 7.0.9   |



