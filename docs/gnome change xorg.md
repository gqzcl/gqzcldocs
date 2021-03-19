```
$ sudo nano ~/.xinitrc
```
### 添加：
```
exec gnome-session
```
### 修改gdm

```bash
$ sudo nano /etc/gdm/custom.conf
```
### 修改：
```
# GDM configuration storage
[daemon]
WaylandEnable=false
DefaultSession=gnome-xorg.desktop

[security]
[xdmcp]
[chooser]

[debug]
#Enable=true
```

### 修改驱动：

```
$ mhwd -l #列出驱动
$ mhwd -li #列出安装的驱动
$ mhwd -i #安装驱动
$ mhwd -i pci 	驱动名	
```

### 查看dispalay：

```
$ inxi -G
```

### 查看FPS

```
$ glxgears
```
### 查看会话窗口
```
$ loginctl
```