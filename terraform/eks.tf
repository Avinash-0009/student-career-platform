module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 21.25"

  name               = var.cluster_name
  kubernetes_version = "1.35"

  endpoint_public_access = true

  enable_cluster_creator_admin_permissions = true

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  addons = {
    coredns = {}

    kube-proxy = {}

    eks-pod-identity-agent = {}

    vpc-cni = {
      before_compute = true
    }
  }

  eks_managed_node_groups = {
    main = {
      name = "student-career-nodes"

      instance_types = ["t3.small"]

      min_size     = 1
      max_size     = 3
      desired_size = 2

      capacity_type = "ON_DEMAND"

      subnet_ids = module.vpc.private_subnets
      tags = {
        "k8s.io/cluster-autoscaler/enabled"            = "true"
        "k8s.io/cluster-autoscaler/student-career-eks" = "owned"
      }
    }
  }

  tags = {
    Environment = var.environment
    Project     = "student-career-platform"
    Terraform   = "true"
  }
}