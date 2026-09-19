# Rebuild Guide

This document records how to recreate the current Student Career
Platform deployment after the AWS/EKS infrastructure is destroyed.

## Architecture

``` text
Terraform
  ↓
AWS VPC + EKS + managed node group
  ↓
Kubernetes / Helm platform components
  ↓
Application manifests + CI/CD
  ↓
Monitoring + centralized logging
```

> Terraform currently provisions the AWS infrastructure. Several
> Kubernetes platform components and EKS Pod Identity associations are
> configured manually; this guide records the current rebuild process.

## 1. Provision AWS infrastructure

``` bash
cd terraform
terraform init
terraform validate
terraform plan
terraform apply
```

Confirm the apply with `yes`.

Terraform provisions the VPC, public/private subnets, NAT Gateway, route
tables, EKS cluster, managed node group, and configured EKS add-ons.

## 2. Configure kubectl

``` bash
aws eks update-kubeconfig   --region ap-south-1   --name student-career-eks   --profile terraform

kubectl get nodes
```

Wait until the nodes are `Ready`.

## 3. Recreate Pod Identity associations

These are currently outside Terraform.

### EBS CSI Driver

``` bash
aws eks create-pod-identity-association   --cluster-name student-career-eks   --role-arn arn:aws:iam::<ACCOUNT_ID>:role/AmazonEKS_EBS_CSI_DriverRole   --namespace kube-system   --service-account ebs-csi-controller-sa   --profile terraform   --region ap-south-1
```

### AWS Load Balancer Controller

``` bash
aws eks create-pod-identity-association   --cluster-name student-career-eks   --role-arn arn:aws:iam::<ACCOUNT_ID>:role/AWSLoadBalancerControllerRole   --namespace kube-system   --service-account aws-load-balancer-controller   --profile terraform   --region ap-south-1
```

### Cluster Autoscaler

``` bash
aws eks create-pod-identity-association   --cluster-name student-career-eks   --role-arn arn:aws:iam::<ACCOUNT_ID>:role/ClusterAutoscalerRole   --namespace kube-system   --service-account cluster-autoscaler   --profile terraform   --region ap-south-1
```

Replace `<ACCOUNT_ID>` with the AWS account ID.

## 4. Install AWS Load Balancer Controller

``` bash
kubectl apply -f k8s/aws-load-balancer-controller-service-account.yaml

helm repo add eks https://aws.github.io/eks-charts
helm repo update
```

Find the new VPC ID:

``` bash
aws ec2 describe-vpcs   --filters "Name=tag:Name,Values=student-career-eks-vpc"   --region ap-south-1   --profile terraform
```

Install:

``` bash
helm upgrade --install aws-load-balancer-controller   eks/aws-load-balancer-controller   -n kube-system   --set clusterName=student-career-eks   --set region=ap-south-1   --set vpcId=<NEW_VPC_ID>   --set serviceAccount.create=false   --set serviceAccount.name=aws-load-balancer-controller   --version 1.14.0
```

Verify:

``` bash
kubectl get pods -n kube-system | grep aws-load-balancer-controller
```

## 5. Create namespace and secrets

``` bash
kubectl apply -f k8s/namespace.yaml

kubectl create secret generic career-secrets   -n student-career   --from-literal=DB_PASSWORD='<DB_PASSWORD>'   --from-literal=JWT_SECRET='<JWT_SECRET>'
```

Do not commit real secret values.

## 6. Deploy PostgreSQL

``` bash
kubectl apply -f k8s/postgres-pvc.yaml
kubectl apply -f k8s/postgres.yaml

kubectl get pvc -n student-career
kubectl get pods -n student-career
```

The PVC should become `Bound` and PostgreSQL should become `Running`.

## 7. Deploy the application

``` bash
kubectl apply -f k8s/backend.yaml
kubectl apply -f k8s/frontend.yaml

kubectl get pods -n student-career
```

## 8. Configure HPA

``` bash
kubectl apply -f k8s/hpa.yaml
kubectl get hpa -n student-career
```

Current configuration:

``` text
minReplicas: 1
maxReplicas: 2
CPU target: 60%
```

## 9. Install Cluster Autoscaler

``` bash
helm repo add autoscaler https://kubernetes.github.io/autoscaler
helm repo update

helm upgrade --install cluster-autoscaler   autoscaler/cluster-autoscaler   --namespace kube-system   --version 9.59.0   --set autoDiscovery.clusterName=student-career-eks   --set awsRegion=ap-south-1   --set cloudProvider=aws   --set rbac.serviceAccount.create=false   --set rbac.serviceAccount.name=cluster-autoscaler   --set 'extraArgs.expander=least-waste'   --set 'extraArgs.balance-similar-node-groups=true'
```

Verify:

``` bash
kubectl get pods -n kube-system | grep autoscaler
```

## 10. Deploy Ingress

``` bash
kubectl apply -f k8s/ingress.yaml
kubectl get ingress -n student-career
```

Get the ALB address:

``` bash
kubectl get ingress student-career-ingress -n student-career
```

Test:

``` bash
curl http://<ALB-DNS>/api/v1/health
```

Expected:

``` json
{
  "message": "Student Career Platform API is running",
  "status": "healthy"
}
```

## 11. Install monitoring

``` bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install kube-prometheus-stack   prometheus-community/kube-prometheus-stack   --version 90.1.1   --namespace monitoring   --create-namespace   -f k8s/monitoring/values.yaml

kubectl apply -f k8s/monitoring/alerts.yaml
```

Verify:

``` bash
kubectl get pods -n monitoring
```

## 12. Install Loki

``` bash
helm repo add grafana-community https://grafana-community.github.io/helm-charts
helm repo update

helm upgrade --install loki   grafana-community/loki   --version 18.13.1   --namespace monitoring   --values k8s/monitoring/loki-values.yaml
```

Verify:

``` bash
kubectl get pods -n monitoring | grep loki
```

## 13. Install Grafana Alloy

``` bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm upgrade --install alloy   grafana/alloy   --version 1.12.1   --namespace monitoring   --values k8s/monitoring/alloy-values.yaml
```

Verify:

``` bash
kubectl get pods -n monitoring | grep alloy
```

Logging flow:

``` text
Kubernetes Pods → Grafana Alloy → Loki → Grafana
```

## 14. Access Grafana

``` bash
kubectl port-forward -n monitoring   svc/kube-prometheus-stack-grafana 3000:80
```

Open:

``` text
http://localhost:3000
```

Loki datasource:

``` text
http://loki.monitoring.svc.cluster.local:3100
```

Verify metrics, dashboards, logs, and alerts.

## 15. CI/CD

After the AWS/EKS platform exists, GitHub Actions handles application
delivery:

``` text
Git Push
   ↓
Backend Tests
   ↓
Frontend Tests + Build
   ↓
SonarQube
   ↓
Docker Build
   ↓
Trivy
   ↓
AWS OIDC
   ↓
ECR
   ↓
EKS Deployment
```

The deployment workflow updates the backend and frontend images using
their ECR image tagged with the Git commit SHA.

Normal application update:

``` bash
git add .
git commit -m "Update application"
git push origin main
```

## 16. Final verification

``` bash
kubectl get nodes
kubectl get pods -A
kubectl get pods -n student-career
kubectl get ingress -n student-career
kubectl get hpa -n student-career
kubectl get pvc -n student-career
kubectl get pods -n monitoring
```

Then:

``` bash
curl http://<ALB-DNS>/api/v1/health
```

Expected final state:

``` text
EKS nodes          → Ready
Backend             → Running
Frontend            → Running
PostgreSQL          → Running
PostgreSQL PVC      → Bound
ALB/Ingress         → Available
HPA                 → Active
Prometheus          → Running
Grafana             → Running
Alertmanager        → Running
Loki                → Running
Alloy               → Running
API health endpoint → Healthy
GitHub Actions      → Successful
```

## 17. Current design decisions

### Terraform scope

Terraform manages the AWS infrastructure. Pod Identity associations are
currently created manually with AWS CLI.

For a production version, these associations could be managed through
Terraform with `aws_eks_pod_identity_association`.

### Secrets

The project uses native Kubernetes Secrets. AWS Secrets Manager +
Secrets Store CSI Driver / ASCP was evaluated but not retained because
of additional cluster resource overhead.

### Terraform state

Terraform state is currently local. A team/production setup should
migrate state to a remote backend with an appropriate locking approach.

### Application deployment

GitHub Actions builds and pushes application images to ECR. The
deployment workflow updates EKS with `kubectl set image`.

### Helm

Helm is used for platform components including AWS Load Balancer
Controller, Cluster Autoscaler, kube-prometheus-stack, Loki, and Grafana
Alloy.

## 18. Rebuild architecture

``` text
                         GitHub
                            │
                            ▼
                    GitHub Actions
                            │
              ┌─────────────┴─────────────┐
              │                           │
             ECR                    SonarQube + Trivy
              │
              ▼
        ┌───────────────┐
        │    AWS EKS    │
        │               │
        │  Frontend     │
        │  Backend      │
        │  PostgreSQL   │
        └───────┬───────┘
                │
                ▼
               ALB
                │
                ▼
             Internet

Observability:
Pods → Alloy → Loki → Grafana

Metrics:
Pods/Nodes → Prometheus → Grafana

Alerts:
Prometheus → Alertmanager → Email
```
