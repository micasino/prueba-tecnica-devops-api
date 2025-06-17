# Build & Push Challenge Instructions

## Objective

You are provided with a Go microservice project located in the `devops/` directory. Your task is to demonstrate your ability to containerize, build, and push this application using modern DevOps practices.

## Tasks

### 1. Dockerize the Application
- Create a `Dockerfile` in the `devops/` directory to build the Go microservice.
- The Dockerfile should:
  - Use a multi-stage build for efficiency.
  - Produce a minimal, production-ready image.
  - Expose the appropriate port for the service.
  - Follow best practices for Go application containerization.

### 2. Automate Build & Push with GitHub Actions
- Create a GitHub Actions workflow (YAML file) in `.github/workflows/` that:
  - Triggers on pushes to the `main` branch.
  - Builds the Docker image from the `devops/` directory.
  - Tags the image with both the commit SHA and the current date (to avoid overwriting previous images).
  - Pushes the image to the following Google Artifact Registry:
    ```
    southamerica-east5-docker.pkg.dev/micasino-devops-test/docker-gar/golang-test
    ```
  - Authenticates securely to Google Cloud (you may use Workload Identity Federation or a Service Account key—document your choice).
- **Note:** The workflow should not overwrite existing images; each build must be uniquely tagged.

### 3. Propose and Implement a CI/CD Deployment Solution
- In addition to a written proposal, you are required to provide some functional code that demonstrates your approach to deploying the built Docker image.
- Your proposal should include:
  - The tools and services you would use.
  - How you would handle rollbacks, secrets management, and zero-downtime deployments.
  - Any diagrams or workflow descriptions that help illustrate your approach.
  - **Functional code**: Provide at least one example of code or configuration that would be used in your proposed CI/CD pipeline.
- This section is open-ended: we are interested in your reasoning, creativity, and understanding of modern DevOps practices.

---

## Deliverables
- A `Dockerfile` in `devops/`.
- A GitHub Actions workflow file in `.github/workflows/`.
- A written and code-based CI/CD deployment proposal (may be in a separate markdown file, or as part of your pull request).
