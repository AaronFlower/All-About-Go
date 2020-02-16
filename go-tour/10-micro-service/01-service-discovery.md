## 服务发现 Service Discovery

当我们在向微服务迁移时，服务发现（Service Discovery) 是我们首先要做的事情。那么什么是服务发现那？怎么实现服务发现，以及有那些选型那？

### 1. 服务发现的历史

在网络不那么发达，我们电脑都是本地局域相连时，我们通过修改 `hosts` 文件来手动添加服务。

当进入互联网时代，服务越来越多，我们通过 DNS 来发现服务。当互联网上出现一个服务时，DNS 服务器会帮我们解析出这个服务。

而在微服务的场景下，新的 Host, Port, 服务 （添加，终止) 变化更加频繁，我们应该使用那种策略来做服务发现那？这里有三种服务发现的方法。

### 2. 当前三种服务发现的方法

1. 基于 DNS 的服务发现

假设公司组织都有了自己的 DNS 服务，那么利用这种方法就可以很方便的实现。基于 DNS 的服务发现系统有 Mesos-DNS 和 Spotify 的 SRV.

2. 基于 Key-Value 存储

第二种方法是基于现有的一致性 Key-Value 存储系统来实现，如: Apache Zookeeper, Consul, etcd。这些都是极复杂的分布式系统。

3. 其它特定的解决方法

除了上述的方法，还有一些特定的解决方案，如 Netflix Eureka.

### 3. 服务发现客户端 （Service discovery clients)

中心化的服务发现策略是必须的（如：Eureka, Consul, DNS), 此外我们的每一个微服务还需要一个客户端来与服务发现做通信。服务发现客户端需要具有两个核心的功能：服务注册(service registration） 和 服务解析(service resolution). 当一个服务启动后，通过服务注册来通知其它服务该启动服务已经可用。而当一个服务可用后，其它服务通过服务解析来定位该网络中的服务。

#### 3.1 服务注册(Service Reistration)

为了服务注册或注销，服务注册客户端需要初始化一个心跳系统 (heartbeat system). 心跳是间隔地通过其它服务该服务还在运行。一般心跳都是异步发送以减少对服务性能的影响。如果系统没有实现心跳策略，那么服务器可以通过轮询的方式来进行查询 。

#### 3.2 服务解析(Service resolution)

返回一个服务的物理网络地址的过程是服务解析。一个典型的服务发现客户端在实现服务解析时一般要实现以下几个特性：caching, failover and load balancing.

- caching: 避免每次服务都进行查询；
- failover, load balancing: 保证高服务的高可用。(round robin algorithm)


### 4. 服务发现的实现

#### 4.1 基于 DNS 实现

#### 4.2 Key-Value store and sidecar

#### 4.3 Specified service discovery and library/sidecar


### References

1. [Service Discovery of microservices](https://www.datawire.io/guide/traffic/service-discovery-microservices/)
