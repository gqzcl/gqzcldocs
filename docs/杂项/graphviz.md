## 设置点和线的形状与颜色

digraph是有向图，graph是无向图。要注意，->和--都表示图中的一条边，但是前者用于有向图中，而后者用于无向图中，不能混用。

代码示例

```text
diaGraph G{
    main -> parse -> execute
    main -> init
    main -> cleanup
    execute -> make_string
    execute -> printf
    init -> make_string
    main -> printf
    execute -> compare
}
```

然后在命令行(cmd)中运行此文件，

```
dot -T输出格式 程序文件名.dot -o 输出文件名.输出格式
```

我们可以手动设置图的属性，即设置图中边的属性与图中点的属性。

我们可以在每条边后面添加中括号，在其中设置边的相关属性，也可以使用edge设置边的默认值。点的属性设置方法相同，node表示点的默认值。

```text
diGraph G{
    size = "4, 4"  /*设置图片尺寸为4 * 4（英寸）*/
    main[shape = box] /*设置点main形状为矩形，默认为椭圆形*/
    main -> parse[weight = 8]  /*设置main到parse的边的重要程度，默认为1*/
    parse -> execute
    main -> init[style = dotted] /*设置main到init的边的样式为点，默认为实线*/
    main -> cleanup
    execute -> {make_string printf} /*一次连接两条边，以隔开目标点*/
    init -> make_string
    edge[color = red]  /*将此语句后的边的颜色设置为红色*/
    main -> printf[style = bold, label = "100 times"]
    node[shape = box, style = filled, color = ".7 .3 1.0"] /*设置此语句后的点的默认属性， 其中color的值采用RGB标准*/
    make_string[label = "make a\nstring"]
    execute -> compare
    compare -> return
}
```

此外，还可以使用dir设置每条边的箭头方向，有forward(default),back,both,nonc，分别表示前向，反向，双向和无。如下所示：

```text
digraph G{
    A -> B[dir = both]
    B -> C[dir = nonc]
    C -> D[dir = back]
    D -> A[dir = forward]
}
```

点的属性除了record和Mrecord这两种之外，其他的形状都是多边形，而我们可以对多边形进行设置。sides用于设置边数，peripheries用于设置多边形外框的层数，regular = true可以使你的多边形是一个规则的多边形，orientation = * 可以让你的多边形旋转*角度，skew后面跟一个（-1.0~1.0）的小数，能使你的图形斜切一个角度，distortion则可以令你的图形产生透视效果。

```text
digraph G{
    a -> b -> c
    b ->d 
    a[shape = polygon, sides = 5, peripheries = 3, color = lightblue, style = filled]
    c[shape = polygon, sides = 4, skew = 0.4, label = "hello world"]
    d[shape = invtriangle]
    e[shape = polygon, sizes = 4, distortion = .7]
}
digraph G{
    A -> B
    A[orientation = 15, regular = true, shape = polygon, sides = 8, peripheries = 4, color = red, style = filled]
    B[shape = polygon, sides = 4, skew = 0.5, color = blue]
}
```

record与Mrecord的区别就是Mrecord的角是圆的，而record是由横竖的矩形组成的图形。

```text
digraph G{
   node[shape = record]
   struct1[label = "<f0> left | <f1> middle | <f2> right"]
   struct2[label = "<f0> one | <f1> two"]
   struct3[label = "hello \nworld | {b | {c | <here> d | e} | f} | g | h"]
   struct1 -> struct2
   struct1 -> struct3
}
```

当你的线与线的label较多时，可以设置线的属性decorate=true使得每条线与其属性之间产生连线。你还可以为每条线设置属性headlabel，taillabel，给每条线的起始点和终点设置属性，它们的颜色由labelfontcolor决定，而label的颜色由fontcolor来决定。

```text
graph G{
   label = "我爱你" /*为图设置标签*/
   labelloc = b    /*图标签的位置在底部，也可以设置为r到顶部*/
   labeljust = l   /*图标签的位置在左边，也可以设置为r到右边*/

   edge[decorate = true]
   C -- D[label = "s1"]
   C -- E[label = "s2"]
   C -- F[label = "s3"]
   D -- E[label = "s4"]
   D -- F[label = "s5"]

   edge[decorate = false, labelfontcolor = blue, fontcolor = red]
   C1 -- D1[headlabel = "d1", taillabel = "c1", label = "c1 - d1"]
}
```

在dot语言中，我们可以使用html语言制作表格。在label后添加<>，而不是""就能添加html语言。

```text
digraph html{
   abc[shape = none, margin = 0, label = <
   <TABLE BORDER = "0" CELLBORDER = "1" CELLSPACING = "0" CELLPADDING = "4">
   <TR><TD ROWSPAN = "3"><FONT COLOR = "red">hello</FONT><BR/>world</TD>
   <TD COLSPAN = "3">b</TD>
   <TD ROWSPAN = "3" BGCCOLOR = "lightgrey">g</TD>
   <TD ROWSPAN = "3">h</TD>
   </TR>
   <TR><TD>c</TD>
   <TD PORT = "here">d</TD>
   <TD>e</TD>
   </TR>
   <TR><TD COLSPAN = "3">f</TD>
   </TR>
   </TABLE>
   >]
}
```

## 设置点，线的位置以及子图的概念

图中的线默认是从上往下的，我们可以通过在文件的最上层设置rankdir=LR将其改为从左往右。同理，rankdir可以设置为TB(默认),BT,RL。

当图中时间表之类的东西时，我们可能需要点排列成一行（列），这时我们就需要rank，在花括号中设置rank=same，然后将需要并排的点一次输入。

```text
digraph G{
   rankdir = LR
   {
    node[shape = plaintext]
    1995 -> 1996 -> 1997 -> 1998 -> 1999 -> 2000 -> 2001
   }
   {
    node[shape = box, style = filled]
    WAR3 -> Xhero -> Footman -> DOTA
    WAR3 -> Battleship
   }
   {
    {rank = same 1996 WAR3}
    {rank = same 1998 Xhero Battleship}
    {rank = same 1999 Footman}
    {rank = same 2001 DOTA}
   }
}
```

设立一条边时，我们可以制定这条边从起点的哪个位置出发，到终点的哪个位置结束。控制符有"n","ne","e","se","s","sw","w"和"nw"，具体效果见下：

```text
digraph G{
   node[shape = box]
   c:n -> d[label = n]
   c1:ne -> d1[label = ne]
   c2:e -> d2[label = e]
   c3:se -> d3[label = se]
   c4:s -> d4[label = s]
   c5:sw -> d5[label = sw]
   c6:w -> d6[label = w]
   c7:nw -> d7[label = nw]
}
```

我们也可以在record中给点定义一些port，以确定record中哪个部分用于连线。

```text
digraph G{
   label = "Binary Search Tree"
   node[shape = record]
   A[label = "<f0> | <f1> A | <f2>"]
   B[label = "<f0> | <f1> B | <f2>"]
   C[label = "<f0> | <f1> C | <f2>"]
   D[label = "<f0> | <f1> D | <f2>"]
   E[label = "<f0> | <f1> E | <f2>"]
   A: f0: sw -> B: f1
   A: f2: se -> C: f1
   B: f0: sw -> D: f1
   B: f2: se -> E: f1
}
```

我们可以借此构建一个Hash表。

```text
digraph G{
    nodesep = .05
    rankdir = LR
    node[shape = record, width = .1, height = .1]

    node0[label = "<f0> | <f1> | <f2> | <f3> | <f4> | <f5> | <f6>",height = 2.5]

    node[width = 1.5]

    node1[label = "{<n> n14 | 719 | <p>}"]
    node2[label = "{<n> a1 | 805 | <p>}"]
    node3[label = "{<n> i9 | 718 | <p>}"]
    node4[label = "{<n> e5 | 989 | <p>}"]
    node5[label = "{<n> t20 | 959 | <p>}"]
    node6[label = "{<n> o15 | 794 | <p>}"]
    node7[label = "{<n> s19 | 659 | <p>}"]

    node0: f0 -> node1: n
    node0: f1 -> node2: n
    node0: f2 -> node3: n
    node0: f5 -> node4: n
    node0: f6 -> node5: n
    node2: p -> node6: n
    node4: p -> node7: n
}
```

画一个子图就是subgraph cluster#，必须有clutser前缀。这里"#"代表任意字符，当然为了便于区分，最好将"#"设置为非英文字母。

```text
digraph G{
    subgraph cluster1{
       node[style = filled, color = white]
       style = filled
       color = lightgrey
       a0 -> a1 -> a2 -> a3
       label = "process1"
    }

    subgraph cluster2{
        node[style = filled]
        color = blue
        b0 -> b1 -> b2 -> b3
        label = "process2"
    }
    start -> a0
    start -> b0
    a1 -> b3
    b2 -> a3
    a3 -> a0
    a3 -> end
    b3 -> end
    start[shape = Mdiamond]
    end[shape = Msquare]
}
```

如果你想将一条边连到一个子图的边界上，先输入compound=true，然后就能用lhead和ltail设置连接的子图。

## 附录

- 点，边和图的属性
  [http://www.graphviz.org/doc/info/attrs.html](https://link.zhihu.com/?target=http%3A//www.graphviz.org/doc/info/attrs.html)
- 点的形状
  [http://www.graphviz.org/doc/info/shapes.html](https://link.zhihu.com/?target=http%3A//www.graphviz.org/doc/info/shapes.html)
- 颜色设置
  [http://www.graphviz.org/doc/info/colors.html](https://link.zhihu.com/?target=http%3A//www.graphviz.org/doc/info/colors.html)

## DOT语法

### 基本语法

graph（无向图）或者digraph（无向图）表示图，然后｛｝中的内容是对图的描述，注释风格和C类似（“//”用于单行注释，/**/用于多行注释）。如一个无向图：

```
graph graph1 {      //无向图graph1  
    a -- b          //节点a和b之间连线  
}  
```

### 节点

DOT中，节点可以不用声明直接使用。每个节点首次出现的名称做为该节点的唯一标识。

属性设置：对节点可以设置的常见通用属性有shape、label、style、color、fillcolor、rank等，对于不同的形状，可能还有不同的属性可以设置，如对于多边形可以设置边数等。节点属性设置时，node用于设置默认属性（对设置位置之后的点有效），在点后面用[]设置单独一个点的属性。

```
graph node_settings {  
    node [shape = "box", style = "filled", color = "black", fillcolor = "green"]   //设置节点的默认形状，类型，颜色，填充颜色  
    a [shape = "ellipse", color = "red", label = "this is a"]     //设置节点a的颜色，注意默认节点类型是filled，所以这里的color只是设置a的边框颜色  
    b [label = "two\nlines"]   //label支持'\n'换行  
    a -- b  
    a -- c     //如果不需要设置c的属性，可以不用声明c而直接使用  
    node [shape = "circle"]  
    d [label = "cicle"]  
    c -- d      //d使用之前最后出现的node设置，形状为circle  
    {rank = same a, d}     //设置a和d在同一层  
}     
```

### 边

边有有向边和无向边两种，无向边用于无向图，有向边用于有向图，不可混用。

属性设置：边的常见设置有style、color、weight、label、labelfontcolor、headlabel、taillabel、decorate等，对于有向边，还可以设置边的起点位置等（用n、e、s、w和相邻字母的组合表示位置）。和节点类似的，边属性设置时，用edge[]设置默认属性，在边之后用[]设置单独一条边的属性。

```
digraph edge_settings {  
    edge [color = "green", decorate = false]        //设置边的默认属性  
    node [shape = "polygon", sides = 4, color = "blue"]  
    a -> b [style = "dotted", color = "red", label = "a to b"]  //设置style、color、label  
    b: se -> c: w [headlabel = "end",  taillabel = "start"]     //设置边从b的“东南方”出发，从c的“西方”结束，设置有向边起点和重点的label  
    edge [style = "bond", decorate = true]      //设置之后的边加粗并且标签和连线之间有线标注  
    {c, f} -> {d, e} [label = "multi-lines"]    //可以用这种方式同时画多条边  
}    
```

### 图

从前面的例子中已经可以看出，DOT语言可以描述无向图和有向图两种图，graph标识无向图，digraph标识有向图。

属性设置：在一个图的开头写入属性即可设置图形的属性，常用的图形属性有size、label、labelloc、labeljust、bgcolor、rankdir等。

子图，可以进行和“父图”类似的设置，唯一注意的是子图必须以"cluster"做为名称的开始。
下面是实现的官网首页上的图：

```
digraph graph_settings {  
    start [shape = "Mdiamond"]  
    end [shape = "Msquare"]  
      
    subgraph cluster_sub1 {  
        label = "process #1"  
        labelloc = "t"  
        bgcolor = "gray55"  
        node [style = "filled", color = "white"]  
        a0 -> a1 -> a2 -> a3 -> a0  
    }  
    subgraph cluster_sub2 {  
        label = "process #2"  
        labelloc = "t"  
        color = "blue"  
        node [style = "filled", color = "black", fillcolor = "gray55"]  
        b0 -> b1 -> b2 -> b3  
    }  
      
    start -> {a0, b0}  
    a1 -> b3  
    b2 -> a3  
    {a3, b3} -> end  
}  
```

 