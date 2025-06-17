# Infrastructure as Code (IaC) Challenge Instructions

## Objective

You are provided with a partially complete Terraform configuration in the `iac/gcp/` directory. Your task is to demonstrate your ability to design, fix, and extend infrastructure as code for a cloud-native application on Google Cloud Platform (GCP).

## Tasks

### 1. Complete and Fix the Provided Terraform Code
- Review the existing Terraform code in `iac/gcp/`.
- Fix any issues or missing configuration so that the following resources are created and functional:
  - A GKE (Google Kubernetes Engine) cluster.
  - A Google Cloud SQL instance for Postgres.
- Ensure the code follows best practices for security, modularity, and maintainability.

### 2. Extend the Infrastructure
- Add Terraform code to provision:
  - A Redis instance (using Memorystore or another managed solution).
  - A Google Cloud Storage bucket.
- Ensure these resources are properly integrated and follow best practices.

### 3. Propose and Implement an IaC Management Solution
- In addition to a written proposal, you are required to provide some functional code that demonstrates your approach to managing and evolving IaC in a real-world scenario.
- Your proposal should include:
  - The tools, workflows, and practices you would use for IaC management.
  - Any diagrams or workflow descriptions that help illustrate your approach.

---

## Deliverables
- A fixed and complete Terraform configuration in `iac/gcp/`.
- Additional Terraform code for Redis and a Storage bucket.
- A written and code-based IaC management proposal (may be in a separate markdown file, or as part of your pull request).

--- 