---
id: Spring Boot (Rapid Application Development)
aliases:
  - Spring Boot (Rapid Application Development)
tags: []
---

# Spring Boot (Rapid Application Development)

Makes creating stand-alone, production-grade applications simple:

```java
// Main application class
@SpringBootApplication
public class MyApplication {
    public static void main(String[] args) {
        SpringApplication.run(MyApplication.class, args);
    }
}

// Properties configuration (application.properties)
// server.port=8080
// spring.datasource.url=jdbc:mysql://localhost:3306/mydb
// spring.jpa.hibernate.ddl-auto=update

// REST Controller
@RestController
@RequestMapping("/api/users")
public class UserRestController {

    private final UserService userService;

    public UserRestController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        User user = userService.findUserById(id);
        return ResponseEntity.ok(user);
    }
}
```
