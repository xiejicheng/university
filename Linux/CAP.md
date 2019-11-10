## CAP 原理

### CAP 定义

**CAP 原理**：指的是 **C**（Consistency）一致性， **A** （Availability）可用性， **P** （Partition tolerance）分区容错性 在一个完整的计算机系统中三种特性不能同时得到完全满足。



#### Consistency（（强）一致性）

指的是在同一时间点里，所有数据状态是否是一致。

#### Availability（（高）可用性）

系统拥有非常高的非故障运行时间百分率，就是计算机性能非常好，能满足用户需求。

#### Partition tolerance（分区容错性）

指的是在整个系统拥有两台或以上数量的计算机单元提供服务，当其中部分分区节点故障时，整个系统依然可以向用户提供可靠的服务。



CAP 定理的存在是想说明 CAP 三种特性不能同时完全满足，因此存在以下三种舍去情况：

#### 1.选择CA，放弃 P

通常是由一台非常强大的计算机对外提供服务来保证A（高可用性），所有的数据操作都在同一台机器上。在数据库层面上通过本地事务来轻松满足C（（强）一致性）。

缺点：由于目前科学技术上的限制，无法横向扩容的系统是无法应付海量请求，大并发的场景的，同时也容易导致单点故障出现，因此这里的A（高可用性）实际上是大打折扣的。

#### 2.选择 CP，放弃 A

意味着系统在提供服务时，一旦出现分区故障，为了确保数据的强一致性，所涉及到的服务全部都会停止响应，因此A（高可用性）也就不复存在啦。暂时还没有想到什么场景下会采取这种方案。

缺点：个人认为低可用性本身就是系统的一大缺点。


#### **3.选择 AP，放弃 C **

这种架构可以说是如今互联网时代最流行的架构。通过分布式和集群进行横向的系统性能拓展，尽可能的舍弃对数据强一致性的需求，同时在遇到不可避免的分布式事务场景时，大牛们也已经提出了各种各样的方案来满足最终的数据一致性（（弱）一致性），就不展开说明了。

缺点：不适合对数据一致性C（（强）一致性）有极高要求的场景，以及不追求A（高可用性）的场景（分区总是会带来额外的麻烦）。

其实，选择AP而放弃C的设计理念并不仅仅适用于整个系统的整体架构，例如在系统内部，数据库集群（可以理解为一个子系统）主从分离的同步延迟造成从库数据与主库数据出现数据暂时不一致的情况等等。



#### 结尾

CAP中的三个元素也并不是完全的互斥，例如选择AP，其实也不是彻底的放弃了C（一致性），而是进行了一定程度的让步，同时C（一致性）也能细分为读一致性，写一致性等。因此不能机粗暴的将CAP理论理解为完全互斥，三选二，而是在这三个维度上面进行不同程度的取舍。个人认为CAP理论其实是告诉人们天下没有免费的午餐，也不存在完美无缺的系统设计，需要反复的斟酌，权衡，才能设计出合理的，符合需求的系统架构。