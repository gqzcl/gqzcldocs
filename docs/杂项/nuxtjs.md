Successfully created project miao

  To get started:

        cd miao
        yarn dev

  To build & start for production:

        cd miao
        yarn build
        yarn start

  To test:

        cd miao
        yarn test

  For TypeScript users. 

		See : https://typescript.nuxtjs.org/cookbook/components



## The engine "node" is incompatible with this module.

方案一：

升级node

#### 安装node版本管理工具'n'

```sh
$ sudo npm install n -g
```

#### 使用版本管理工具安装指定node或者升级到最新node版本

```sh
$ sudo n stable  （安装node最新版本）
```

```sh
$ sudo n 8.9.4 （安装node指定版本8.9.4）
```

方案二：

忽略错误重新安装

yarn config set ignore-engines true



