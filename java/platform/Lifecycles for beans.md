---
id: Lifecycles for beans
aliases:
  - Lifecycles for beans
tags: []
---

# Lifecycles for beans

1. **Singleton** (Default): Only one instance of the bean is created per Spring IoC container. This is similar to the "Singleton" scope in other frameworks.

   ```java
   @Bean
   @Scope("singleton")
   public UserService userService() {
       return new UserServiceImpl();
   }

   // Or with annotation on the class
   @Component
   @Scope("singleton")
   public class UserServiceImpl implements UserService { }
   ```

2. **Prototype**: A new instance is created each time the bean is requested. This is similar to the "Transient" scope in ASP.NET.

   ```java
   @Bean
   @Scope("prototype")
   public UserService userService() {
       return new UserServiceImpl();
   }
   ```

3. **Request**: A single instance per HTTP request (only valid in web-aware Spring applications).

   ```java
   @Component
   @Scope(value = WebApplicationContext.SCOPE_REQUEST, proxyMode = ScopedProxyMode.TARGET_CLASS)
   public class RequestScopedBean { }
   ```

4. **Session**: A single instance per HTTP session (only in web-aware Spring applications).

   ```java
   @Component
   @Scope(value = WebApplicationContext.SCOPE_SESSION, proxyMode = ScopedProxyMode.TARGET_CLASS)
   public class SessionScopedBean { }
   ```

5. **Application**: A single instance per ServletContext (only in web-aware Spring applications).

6. **Websocket**: A single instance per WebSocket session.

Key differences from ASP.NET:

- Spring's default is "singleton" (one instance per Spring container)
- Spring doesn't use the term "transient" - its equivalent is "prototype"
- Spring offers web-specific scopes like "request" and "session" that are similar to ASP.NET's "scoped" concept

You can specify these scopes either through annotations or in XML configuration. When using web-specific scopes, you typically need to use proxy mode to handle the lifecycle correctly.
