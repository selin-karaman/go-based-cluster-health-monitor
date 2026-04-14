# ☸️ Go-Based Cluster Health Monitor 
This project demonstrates a professional Cloud-Native tool implementation using Go (Golang) to monitor Kubernetes cluster health. It bridges the gap between infrastructure management and software development by providing a lightweight, automated inspection tool that interacts directly with the Kubernetes API.

## 🛠️ Tech Stack
* Language: Go (Golang)
* SDK: Kubernetes Client-Go
* Containerization: Docker (Multi-stage builds)
* Orchestration: Kubernetes (Minikube)
* Security: RBAC (ServiceAccounts, ClusterRoles)

## 🏗️ Architecture & Workflow
* Hybrid Configuration: The application automatically detects its environment. It utilizes InClusterConfig when running as a Pod or falls back to local kubeconfig for external execution.

* API Interaction: Connects to the Kubernetes API Server via a secure TLS tunnel to fetch real-time data of all pods across all namespaces.

* RBAC Authorization: Operates under a strictly defined ServiceAccount, using a ClusterRole with minimal "list/get" permissions to adhere to the principle of least privilege.

* Optimized Packaging: Utilizes a multi-stage Dockerfile to compile the Go source into a static binary, resulting in a highly optimized image (~80MB).

## 🗒️ Tool Visualization
### 1. Cluster Health Report

The tool generates a structured table in the terminal/logs, displaying the Namespace, Pod Name, Status, and a calculated Health indicator for every resource in the cluster.

### 2. In-Cluster Execution Logs

Final execution logs from within the Kubernetes pod confirming successful authentication via ServiceAccount and a "PERFECT" cluster status report.

## 🚀 How to Run
### 1. Clone the Repo:
```
git clone https://github.com/selin-karaman/go-based-cluster-health-monitor.git
```

### 2. Build and Load Image (Minikube):

```
docker build -t kubecheck:v2 .
minikube image load kubecheck:v2
```

### 3. Deploy to Kubernetes:
```
# Apply RBAC and Pod manifests
kubectl apply -f k8s-deploy.yaml

# Check the report
kubectl logs kubecheck-monitor
```

## 🧠 Key Learning Outcomes
Through the development of this Go-centric Kubernetes monitoring infrastructure, I achieved the following technical milestones:

* Kubernetes Client-Go Mastery: Developed a deep understanding of the Kubernetes API machinery, specifically using the client-go library to programmatically manage cluster resources.

* RBAC Security Implementation: Mastered the creation and binding of ServiceAccounts, ClusterRoles, and RoleBindings to secure containerized applications within a cluster.

* Cloud-Native Optimization: Achieved a ~90% reduction in image size compared to traditional runtimes by utilizing Go's static binary compilation and Docker multi-stage builds.

* Hybrid Logic Development: Implemented a robust "Detection Logic" that allows the same binary to function seamlessly both as a local CLI tool and a containerized Kubernetes agent.



