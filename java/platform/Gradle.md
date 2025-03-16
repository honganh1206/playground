---
id: Gradle
aliases:
  - Gradle
tags: []
---

# Gradle

Gradle is a newer, more flexible build tool that uses a Groovy or Kotlin-based DSL instead of XML. It's becoming the preferred choice for many Java projects, especially Android development.

**Key Gradle Concepts:**

- **build.gradle**: Script file containing project configuration (Groovy or Kotlin)
- **Tasks**: Units of work that Gradle executes (more flexible than Maven's fixed lifecycle)
- **Dependency configurations**: Different scopes for dependencies (implementation, testImplementation, etc.)
- **Plugins**: Extend Gradle functionality for specific project types

**Simple Gradle build file example:**

```groovy
plugins {
    id 'java'
    id 'application'
}

group = 'com.example'
version = '1.0-SNAPSHOT'
sourceCompatibility = '17'

repositories {
    mavenCentral()
}

dependencies {
    implementation 'org.springframework:spring-core:5.3.20'
    testImplementation 'junit:junit:4.13.2'
}

application {
    mainClass = 'com.example.Main'
}

tasks.named('test') {
    useJUnitPlatform()
}
```
