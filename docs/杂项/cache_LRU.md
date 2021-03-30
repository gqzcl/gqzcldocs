LRU的英文全称是Least Recently Used，也即最不经常使用。我们看着好像挺迷糊的，其实这个含义要结合缓存一起使用。对于工程而言，**缓存是非常非常重要的机制**，尤其是在当下的互联网应用环境当中，起到的作用非常重要。为了便于大家更好地理解，我们从缓存的机制开始说起。

## **缓存**

缓存的英文是cache，最早其实指的是用于CPU和主存数据交互的。早年这块存储被称为高速缓存，最近已经听不到这个词了，不知道是不是淘汰了。因为缓存的读写速度要高于CPU低于主存，所以是用来过渡数据用的。CPU从缓存当中读取数据，主存的数据也会先加载到缓存当中来，之后再进入CPU。

后来缓存有了更多的应用和意为，**在后端服务当中一般用来快速响应请求**。其实这里的思想和记忆化搜索有点像，我们把可能要用到的数据先存起来，后期如果要用到的话，就可以直接从内存当中读取而不是再去计算一遍了。原理也是一样的，有了缓存我们可以把要返回给用户的数据储存在内存中，当同样的请求过来的时候，我们就可以直接从内存当中读取结果，而不是再走一次链路获取数据了。

举一个很简单的例子，比如说我们打开淘宝首页会看到很多商品的推荐。其实推荐商品的流程是非常复杂的，首先要根据一定的策略去商品库当中召回商品。比如根据用户之前的行为召回和历史行为相关的商品，召回了商品之后还要进行清洗，过滤掉一些明确不感兴趣或者是非法、违规的商品。过滤了之后需要使用机器学习或者是深度学习的模型来进行点击率预测，从而将发生点击可能性最高的商品排在前面。

到这里还没结束，推荐商品列表其实也是分展位的，有些位置的商品是运营配好的，有些位置固定展示的是广告。广告往往也有自己的一条链路，还有些位置有一些其他的逻辑。这些商品的数据都拿到了之后，还要获取图片以及其他一些零零散散的信息，最后才能展示出来。因此大家可以试一下**打开淘宝首页要比打开百度首页慢得多**，这并不是淘宝技术差，而是因为这中间的链路实在是太长了。

![img](https://pic2.zhimg.com/80/v2-66146e7e4f1fe3429d742bf089875f21_720w.jpg)

我们很容易发现，对于一些经常打开淘宝的用户来说，其实没有必要每一次都把完整的链路走一遍。我们**大可以把之前展示的结果存储在内存里，下一次再遇到同样请求的时候，直接从内存当中读取并且返回就可以了**。这样可以节省大量的时间以及计算资源，比如在大促的时候，就可以把计算资源节省下来用在更加需要的地方。

缓存虽然好用，但是也不是万能的，因为内存是很贵的，我们不可能把所有数据都存在内存里。内存里只能放一些我们认为比较高价值的数据，在这种情况下，计算科学家们想出了种种策略来调度缓存，保持缓存当中数据的高价值。LRU就是其中一种比较常用的策略。

## **LRU含义**

我们前面也说了，LRU的意思是最长不经常使用，也可以理解成最久没有使用。在这种策略下我们**用最近一次使用的时间来衡量一块内存的价值**，越久之前使用的价值也就越低，最近刚刚使用过的，后面接着会用到的概率也就越大，那么自然也就价值越高。

当然只有这个限制是不够的，我们前面也说了，由于内存是非常金贵的，导致我们可以存储在缓存当中的数据是有限的。比如说我们固定只能存储1w条，当内存满了之后，缓存每插入一条新数据，都要抛弃一条最长没有使用的旧数据。这样我们就保证了缓存当中的数据的价值都比较高，并且内存不会超过限制。

我们把上面的内容整理一下，可以得到几点要求：

1. 保证缓存的读写效率，比如读写的复杂度都是![[公式]](https://www.zhihu.com/equation?tex=O%281%29)
2. 当一条缓存当中的数据被读取，将它最近使用的时间更新
3. 当插入一条新数据的时候，弹出更新时间最远的数据

## **LRU原理**

我们仔细想一下这个问题会发现好像没有那么简单，显然我们不能通过数组来实现这个缓存。因为数组的查询速度是很慢的，不可能做到![[公式]](https://www.zhihu.com/equation?tex=O%281%29)。其次我们用HashMap好像也不行，因为虽然查询的速度可以做到![[公式]](https://www.zhihu.com/equation?tex=O%281%29)，但是我们没办法做到更新最近使用的时间，并且快速找出最远更新的数据。

如果是在面试当中被问到想到这里的时候，可能很多人都已经束手无策了。但是先别着急，我们冷静下来想想会发现问题其实并没有那么模糊。首先**HashMap是一定要用的**，因为只有HashMap才可以做到![[公式]](https://www.zhihu.com/equation?tex=O%281%29)时间内的读写，其他的数据结构几乎都不可行。但是只有HashMap解决不了更新以及淘汰的问题，必须要配合其他数据结构进行。这个数据结构**需要能够做到快速地插入和删除**，其实我这么一说已经很明显了，只有一个数据结构可以做到，就是链表。

链表有一个问题是我们想要查询链表当中的某一个节点需要![[公式]](https://www.zhihu.com/equation?tex=O%28n%29)的时间，这也是我们无法接受的。但这个问题并非无法解决，实际上解决也很简单，我们只需要把链表当中的节点作为HashMap中的value进行储存即可，最后得到的系统架构如下：

![img](https://pic4.zhimg.com/80/v2-98b6c3d744a74c9b639ef6734bc5eb23_720w.jpg)

对于缓存来说其实只有两种功能，第一种功能就是**查找**，第二种是**更新**。

### **查找**

查找会分为两种情况，第一种是没查到，这种没什么好说的，直接返回空即可。如果能查到节点，在我们返回结果的同时，我们**需要将查到的节点在链表当中移动位置**。我们假设越靠近链表头部表示数据越旧，越靠近链表尾部数据越新，那么当我们查找结束之后，我们需要把节点移动到链表的尾部。

移动可以通过两个步骤来完成，第一个步骤是在链表上删除该节点，第二个步骤是插入到尾部：

![img](https://pic1.zhimg.com/80/v2-b9b7be93c4621ac314f8330ad6637360_720w.jpg)

### **更新**

更新也同样分为两种情况，第一种情况就是更新的key已经在HashMap当中了，那么我们只需要更新它对应的Value，并且将它移动到链表尾部。对应的操作和上面的查找是一样的，只不过多了一个更新HashMap的步骤，这没什么好说的，大家应该都能想明白。

第二种情况就是要更新的值在链表当中不存在，这也会有两种情况，第一个情况是缓存当中的数量还没有达到限制，那么我们直接添加在链表结尾即可。第二种情况是链表现在已经满了，我们需要移除掉一个元素才可以放入新的数据。这个时候我们需要删除链表的第一个元素，不仅仅是链表当中删除就可以了，**还需要在HashMap当中也删除对应的值**，否则还是会占据一份内存。

因为我们要进行链表到HashMap的反向删除操作，所以链表当中的节点上也需要记录下Key值，否则的话删除就没办法了。删除之后再加入新的节点，这个都很简单：

![img](https://pic1.zhimg.com/80/v2-b2156e017d304d6a71477bf321bba0e8_720w.jpg)

我们理顺了整个过程之后来看代码：

```python
class Node:
    def __init__(self, key, val, prev=None, succ=None):
        self.key = key
        self.val = val
        # 前驱
        self.prev = prev
        # 后继
        self.succ = succ

    def __repr__(self):
        return str(self.val)


class LinkedList:
    def __init__(self):
        self.head = Node(None, 'header')
        self.tail = Node(None, 'tail')
        self.head.succ = self.tail
        self.tail.prev = self.head

    def append(self, node):
        # 将node节点添加在链表尾部
        prev = self.tail.prev
        node.prev = prev
        node.succ = prev.succ
        prev.succ = node
        node.succ.prev = node

    def delete(self, node):
        # 删除节点
        prev = node.prev
        succ = node.succ
        succ.prev, prev.succ = prev, succ

    def get_head(self):
        # 返回第一个节点
        return self.head.succ


class LRU:
    def __init__(self, cap=100):
        # cap即capacity，容量
        self.cap = cap
        self.cache = {}
        self.linked_list = LinkedList()


    def get(self, key):
        if key not in self.cache:
            return None

        self.put_recently(key)
        return self.cache[key]

    def put_recently(self, key):
        # 把节点更新到链表尾部
        node = self.cache[key]
        self.linked_list.delete(node)
        self.linked_list.append(node)

    def put(self, key, value):
        # 能查到的话先删除原数据再更新
        if key in self.cache:
            self.linked_list.delete(self.cache[key])
            self.cache[key] = Node(key, value)
            self.linked_list.append(self.cache[key])
            return 

        if len(self.cache) >= self.cap:
            # 如果容量已经满了，删除最旧的节点
            node = self.linked_list.get_head()
            self.linked_list.delete(node)
            del self.cache[node.key]

        u = Node(key, value)
        self.linked_list.append(u)
        self.cache[key] = u
```

在这种实现当中我们没有用除了dict之外的其他任何工具，连LinkedList都是自己实现的。实际上在Python语言当中有一个数据结构叫做**OrderedDict**，它是一个字典，但是它当中的元素都是有序的。我们利用OrderedDict来实现LRU就非常非常简单，代码量也要少得多。

我们来看代码：

```python
class LRU(OrderedDict):

    def __init__(self, cap=128, /, *args, **kwds):
        self.cap = cap
        super().__init__(*args, **kwds)

    def __getitem__(self, key):
        # 查询函数
        value = super().__getitem__(key)
        # 把节点移动到末尾
        self.move_to_end(key)
        return value

    def __setitem__(self, key, value):
        # 更新函数
        super().__setitem__(key, value)
        if len(self) > self.cap:
            oldest = next(iter(self))
            del self[oldest]
```

在上面一种实现当中由于只用了一个数据结构，所以整个代码要简洁许多。使用起来也更加方便，直接像是dict一样使用方括号进行查询以及更新就可以了。不过在其他语言当中可能没有OrderedDict这种数据结构，这就需要我们自己来编码实现了。

好了，今天的文章就到这里。

原文：https://link.zhihu.com/?target=https%3A//mp.weixin.qq.com/s%3F__biz%3DMzUyMTM5OTM2NA%3D%3D%26mid%3D2247487607%26idx%3D1%26sn%3D5a3056380420d17947421337f6a88fef%26chksm%3Df9daed5ccead644a5bde48d6a3cf6d5a2121e6a087a932fc79b8b304aa84178e81bd02caf87e%26token%3D1931832547%26lang%3Dzh_CN%23rd