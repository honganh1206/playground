---
id: Dependency Injection in Spring
aliases:
  - Dependency Injection in Spring
tags: []
---

# Dependency Injection in Spring

The fundamental feature of Spring is its Inversion of Control (IoC) container.

```java
// Bean definition
public class UserService {
    private UserRepository userRepository;

    // Constructor injection
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User findUserById(Long id) {
        return userRepository.findById(id);
    }
}

// Configuration class
@Configuration
public class AppConfig {

    @Bean
    public UserRepository userRepository() {
        return new JpaUserRepository();
    }

    @Bean
    public UserService userService() {
        return new UserService(userRepository());
    }
}

// Using the beans
public class Application {
    public static void main(String[] args) {
        AnnotationConfigApplicationContext context =
            new AnnotationConfigApplicationContext(AppConfig.class);

        UserService userService = context.getBean(UserService.class);
        User user = userService.findUserById(1L);

        context.close();
    }
}
```

[[Lifecycles for beans]]
