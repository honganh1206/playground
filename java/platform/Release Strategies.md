---
id: Release Strategies
aliases:
  - Release Strategies
tags: []
---

# Release Strategies

### Feature Flags

- Libraries like FF4J specifically for Java applications
- Decouples deployment from feature release

### Canary Releases

- Gradual rollout to subset of users
- Often implemented via service mesh or load balancer configuration

### Blue-Green Deployment

- Two identical environments with only one live
- Instant rollback capability
