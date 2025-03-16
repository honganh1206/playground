---
id: Maven
aliases:
  - Maven
tags: []
---

# Maven

Maven is an older but widely used **build automation and dependency management tool** for Java projects.

It uses an XML-based configuration approach.

**Key Maven Concepts:**

- **POM (Project Object Model)**: XML file (`pom.xml`) that contains project configuration
- **Repositories**: Central locations where dependencies are stored and retrieved from
- **Lifecycle**: Predefined build phases (compile, test, package, install, deploy)
- **Coordinates**: Maven identifies artifacts using groupId:artifactId:version format

**Simple Maven POM example:**

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>my-app</artifactId>
    <version>1.0-SNAPSHOT</version>

    <dependencies>
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>4.13.2</version>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-core</artifactId>
            <version>5.3.20</version>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.10.1</version>
                <configuration>
                    <source>17</source>
                    <target>17</target>
                </configuration>
            </plugin>
        </plugins>
    </build>
</project>
```
