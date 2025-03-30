---
id: Packaging and Application Structure
aliases:
  - Packaging and Application Structure
tags: []
---

# Packaging and Application Structure

### JAR (Java Archive)

- Standard Java packaging format containing **compiled** classes, resources, and metadata
- Used for libraries and simple applications

### WAR (Web Application Archive)

- Extension of JAR format designed for web applications
- Contains web resources (HTML, CSS, JS) alongside Java classes
- Deployed to servlet containers like Tomcat or Jetty
- Less common in modern microservices architectures

### EAR (Enterprise Application Archive)

- Contains multiple modules including JARs, WARs, and deployment descriptors
- Used primarily in Java EE/Jakarta EE environments
- Deployed to application servers like WildFly or WebSphere

### Fat/Uber JARs

- Self-contained executable JARs with all dependencies included
- Tools: Maven Shade Plugin, Gradle Shadow Plugin
- Popular for microservices deployment
- Examples: Spring Boot executable JARs
