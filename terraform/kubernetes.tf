resource "kubernetes_deployment" "nginx" {
  metadata {
    name = "scalable-nginx-challenge"
    labels = {
      App = "ScalableNginxChallenge"
    }
  }

  spec {
    replicas = 2
    selector {
      match_labels = {
        App = "ScalableNginxChallenge"
      }
    }
    template {
      metadata {
        labels = {
          App = "ScalableNginxChallenge"
        }
      }
      spec {
        container {
          image = "nginx:1.7.8"
          name  = "challenge"

          port {
            container_port = 80
          }

          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}
