# 1 goctl 
```
COMMANDS:
   bug         report a bug
   upgrade     upgrade goctl to latest version
   env         check or edit goctl environment
   migrate     migrate from tal-tech to zeromicro
   api         generate api related files    //
   docker      generate Dockerfile
   kube        generate kubernetes files
   rpc         generate rpc code             // 
   model       generate model code   // 根据表生成model
   template    template operation
   completion  generation completion script, it only works for unix-like OS
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
# 2 goctl api
```
goctl api go -api *.api -dir ../  --style=goZero
```

# 3 goctl docker
```
goctl docker -go user.go  // 基于入口文件user.go
```

# 4 goctl kube
```
goctl kube deploy -name user-api -namespace go-zero-demo -image user-api:v1.0 -o user-api.yaml -port 10001 -nodePort 30001
```
得到以下文件
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: user-api
  namespace: go-zero-demo
  labels:
    app: user-api
spec:
  replicas: 3
  revisionHistoryLimit: 5
  selector:
    matchLabels:
      app: user-api
  template:
    metadata:
      labels:
        app: user-api
    spec:
      containers:
      - name: user-api
        image: user-api:v1.0
        lifecycle:
          preStop:
            exec:
              command: ["sh","-c","sleep 5"]
        ports:
        - containerPort: 10001
        readinessProbe:
          tcpSocket:
            port: 10001
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 10001
          initialDelaySeconds: 15
          periodSeconds: 20
        resources:
          requests:
            cpu: 500m
            memory: 512Mi
          limits:
            cpu: 1000m
            memory: 1024Mi
        volumeMounts:
        - name: timezone
          mountPath: /etc/localtime
      volumes:
        - name: timezone
          hostPath:
            path: /usr/share/zoneinfo/Asia/Shanghai

---

apiVersion: v1
kind: Service
metadata:
  name: user-api-svc
  namespace: go-zero-demo
spec:
  ports:
    - nodePort: 30001
      port: 10001
      protocol: TCP
      targetPort: 10001
  type: NodePort
  selector:
    app: user-api

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: user-api-hpa-c
  namespace: go-zero-demo
  labels:
    app: user-api-hpa-c
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: user-api
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80

---

apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: user-api-hpa-m
  namespace: go-zero-demo
  labels:
    app: user-api-hpa-m
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: user-api
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: memory
      targetAverageUtilization: 80
```
// 启动服务
```
kubectl app -f user-api.yaml
```



