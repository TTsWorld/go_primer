#基于 centos:7 这个镜像来构建自定义镜像。这里需要注意，每个 Dockerfile 的第一行除了注释都必须以 FROM 开头
FROM centos:7
#
# 将 nginx.repo 文件复制到容器中
COPY nginx.repo /etc/yum.repos.d/nginx.repo
# 安装 nginx
RUN yum install -y nginx
# 暴露 80 端口
EXPOSE 80
# 设置环境变量
ENV HOST=mynginx
# 启动容器时运行的命令
CMD ["nginx","-g","daemon off;"]