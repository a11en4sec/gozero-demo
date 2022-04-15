# 1 别名，在config.go中定义
```yaml
UserRpcConf:
  # Etcd:
  #   Hosts:
  #     - 127.0.0.1:2379
  #   Key: user.rpc

#   Endpoints:
#     - 127.0.0.1:8080

    Target: k8s://go-zero-looklook/basic-rpc-svc:9001
```
- Target:           用来指定k8s内部定义的k8s协议
- k8s:              使用k8s来做服务注册和发现
- go-zero-looklook: 命名空间
- basic-rpc-svc:    yaml文件中定义的serviceName
- 9001:             serviceName的端口           

# 2 生成dockerfile
```
goctl docker -go user.go
```
# 3 打镜像
```
docker build -t gozerok8s-basic-rpc:v2 .
```

# 4 为生成登录镜像仓库的相关凭证

## 4.1 登录到镜像仓库
```
$ docker login 192.168.1.180:8077
$ Username: admin
$ Password:
Login Succeeded
```
## 4.2 对密码镜像base64
```
#查看上一步登陆harbor生成的凭证
$ cat /root/.docker/config.json  
{
	"auths": {
		"192.168.1.180:8077": {
			"auth": "YWRtaW46SGFyYm9yMTIzNDU="
		}
}
```
## 4.3 对密钥进行加密
```
$ cat /root/.docker/config.json  | base64 -w 0

ewoJImF1dGhzIjogewoJCSIxOTIuMTY4LjEuMTgwOjgwNzciOiB7CgkJCSJhdXRoIjogIllXUnRhVzQ2U0dGeVltOXlNVEl6TkRVPSIKCQl9Cgl9Cn0=
```
## 4.4 创建docker-secret.yaml
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: docker-login
type: kubernetes.io/dockerconfigjson
data:
  .dockerconfigjson: ewoJImF1dGhzIjogewoJCSIxOTIuMTY4LjEuMTgwOjgwNzciOiB7CgkJCSJhdXRoIjogIllXUnRhVzQ2U0dGeVltOXlNVEl6TkRVPSIKCQl9Cgl9Cn0=
```

```
$ kubectl create -f docker-secret.yaml -n go-zero-looklook
```

# 5 在命名空间中创建账号,并分配查看Endports的权限(RABC)
## auth.yaml
```yaml
#创建账号
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: go-zero-looklook
  name: find-endpoints

---
#创建角色对应操作
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: discov-endpoints
rules:
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["get","list","watch"]

---
#给账号绑定角色
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: find-endpoints-discov-endpoints
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: discov-endpoints
subjects:
- kind: ServiceAccount
  name: find-endpoints
  namespace: go-zero-looklook
```

## 在k8s中生成

```
kubectl apple -f auth.yaml
kubectl sa -n go-zero-looklook
```

# 6 部署rpc服务
```yaml
goctl kube deploy -replicas 3 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name basic-rpc -namespace go-zero-xxxx -image xxxx.xxx.com/xxxx-rpc:v1 -o basic-rpc.yaml -port 9001 --serviceAccout find_endpoints
```
```
kubuctl apply -f basic-api.yaml
```

# 7 部署api服务
```yaml
goctl kube deploy -nodePort 32010 -replicas 3 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name basic-api -namespace go-zero-xxxx -image xxxx.xxx.com/xxxx-api:v1 -o basic-api.yaml -port 8888 --serviceAccout find_endpoints
```

```
kubuctl apply -f basic-api.yaml
```

# 8 查看部署情况
```
kubect get svc -n your-ns
kubect get pot -n your-ns
```
