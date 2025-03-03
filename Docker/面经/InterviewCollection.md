# InterviewCollection

## 什么是Docker？

Docker是一个开源的应用容器引擎，允许开发者打包他们的应用及其依赖项到一个轻量级、可移植的容器中，然后发布到任何流行的Linux机器上，也可以实现虚拟化。容器完全使用沙箱机制，相互之间不会有任何接口（类似Linux的chroot）。

## Docker和虚拟机（VM）之间的区别是什么？

虚拟机创建一个完整的操作系统实例，而Docker容器共享主机的操作系统内核，因此容器更加轻量级且启动更快。虚拟机通常占用更多的资源，因为它们需要自己的操作系统副本。

## Docker 的虚拟化机制和底层原理



## Docker容器有几种状态？

Docker容器有以下几种状态：

- 运行（Running）
- 已暂停（Paused）
- 已停止（Stopped）
- 重新启动（Restarting）
- 已退出（Exited）

## Dockerfile中最常见的指令有哪些？

Dockerfile中常见的指令包括：

- FROM：指定基础镜像。
- MAINTAINER：指定镜像的作者。
- RUN：执行命令来设置镜像。
- CMD：指定容器启动时默认执行的命令。
- EXPOSE：声明容器将监听的端口。

## COPY和ADD指令的区别是什么？

COPY和ADD指令都用于将文件从本地文件系统复制到镜像中。ADD指令还支持远程URLs和tar归档的自动提取，而COPY仅支持简单的文件复制。

## 什么是Docker镜像？

Docker镜像是构建容器的基础。它是一个只读模板，包含启动容器所需的所有文件和依赖项。

## Docker网络模式有哪些？

Docker提供了四种网络模式：bridge、host、none和container。

## 如何在Docker容器间进行通信？

容器间通信可以通过Docker网络（如桥接网络）实现，允许容器通过服务名或IP地址相互通信。

## Docker容器如何共享文件系统？

容器可以通过数据卷（Volumes）、bind mounts或tmpfs mounts与宿主机或其他容器共享文件系统。

## Docker Compose的作用是什么？

Docker Compose是一个工具，用于定义和运行多容器Docker应用，允许用户在一个YAML文件中描述整个应用的环境。

## Docker Swarm和Kubernetes的区别是什么？

Docker Swarm和Kubernetes都是容器编排工具，但Kubernetes具有更强大的自动化、可伸缩性和可移植性，而Swarm更简单，适用于较小的集群。

## Docker容器和LXC（Linux Containers）有什么区别？

Docker容器基于Go语言编写，具有更好的可移植性和标准化，而LXC是一个较早的容器技术，使用C语言编写。

## Docker如何实现资源隔离？

Docker利用Linux的命名空间（namespaces）和控制组（cgroups）技术来实现资源隔离。命名空间提供了进程、网络、文件系统等方面的隔离，而cgroups则限制、记录和隔离进程组使用的物理资源（CPU、内存、磁盘I/O等）。

## Docker如何处理跨主机容器通信？

在单个主机上，Docker使用桥接网络模型来实现容器间的通信。对于跨主机通信，可以使用Docker Overlay网络，或者集成外部SDN解决方案如Weave、Calico、Flannel等。

## 解释Docker的存储驱动机制。

Docker使用不同的存储驱动来管理容器层的文件系统。常见的存储驱动有AUFS、OverlayFS、Btrfs、ZFS等，每种驱动都有其特点和适用场景，如性能、空间效率和数据一致性等。

## Docker如何处理数据持久化？

Docker通过数据卷（Volumes）和绑定挂载（Bind Mounts）来实现数据持久化。数据卷由Docker管理，独立于容器存在，即使容器被删除数据也不会丢失；而绑定挂载则是将宿主机上的目录或文件直接挂载到容器中。

## Docker的网络模式如何影响容器的网络性能？

不同的网络模式会影响容器的网络性能。例如，“host”模式下容器使用宿主机的网络堆栈，性能最好但没有隔离；“bridge”模式下容器通过虚拟网桥通信，提供了较好的隔离性但可能有额外的网络延迟。

## 如何解决Docker容器中的DNS解析问题？

Docker容器中的DNS解析问题可以通过配置容器内的resolv.conf文件，或者使用Docker DNS插件和自定义网络来解决，确保容器可以正确解析域名。

## Docker的Layered File System（分层文件系统）如何工作？

Docker镜像是分层的，每一层代表一个Dockerfile指令的结果。当容器运行时，这些层堆叠在一起形成完整的文件系统。这种分层结构使得镜像构建和更新高效，并能节省存储空间。

## Docker容器的生命周期管理命令有哪些？

Docker容器的生命周期管理命令包括`docker run`（启动容器）、`docker stop`（停止容器）、`docker start`（重启停止的容器）、`docker restart`（重启容器）、`docker rm`（移除容器）等。

## Docker的安全性考虑有哪些？

Docker的安全性考虑包括使用可信的镜像来源、最小权限原则、定期更新和打补丁、使用Docker Secrets管理敏感信息、实施容器防火墙规则等。

## 如何优化Docker镜像的大小？

减小Docker镜像的大小可以通过使用小型的基础镜像、多阶段构建、清理不必要的依赖、压缩和合并文件、以及避免使用全局变量等方式来实现。

## docker 如何限制容器的CPU、内存等硬件资源

在 Docker 中，可以通过运行时参数限制容器的 CPU、内存等硬件资源。

内存限制：

* `--memory`或 `-m` 设置容器可用的最大物理内存。
* --memory-swap 设置物理内存+交换分区 Swap 总限制 与 --memory 配合使用

CPU 限制

* `--cpus` 指定容器可使用的 CPU 核心数（支持小数，如 `1.5` 表示 1.5 个核心）。
* `--cpuset-cpus`：绑定容器到特定的 CPU 核心（如 `0-2` 或 `0,1`）。

## Dockerfile 是什么，是如何构建容器的



## 如何减小镜像

