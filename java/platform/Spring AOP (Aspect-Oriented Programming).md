---
id: Spring AOP (Aspect-Oriented Programming)
aliases:
  - Spring AOP (Aspect-Oriented Programming)
tags: []
---

# Spring AOP (Aspect-Oriented Programming)

Spring AOP is a framework implementation of aspect-oriented programming that allows you to _separate cross-cutting concerns from your business logic_.

Key Concepts in Spring AOP:

1. Aspects: Classes that implement cross-cutting concerns (logging, security, transactions)
2. Join Points: Points in program execution where aspects can be applied
3. Pointcuts: Expressions that define which join points to apply aspects to
4. Advice: Action taken at a specific join point (Before, After, Around)

```java
@Aspect
@Component
public class LoggingAspect {

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @Before("execution(* com.example.service.*.*(..))")
    public void logBefore(JoinPoint joinPoint) {
        logger.info("Before executing: " + joinPoint.getSignature().getName());
    }

    @AfterReturning(pointcut = "execution(* com.example.service.*.*(..))", returning = "result")
    public void logAfterReturning(JoinPoint joinPoint, Object result) {
        logger.info("Method " + joinPoint.getSignature().getName() + " returned: " + result);
    }

    @Around("@annotation(Loggable)")
    public Object logExecutionTime(ProceedingJoinPoint joinPoint) throws Throwable {
        long start = System.currentTimeMillis();
        Object proceed = joinPoint.proceed();
        long executionTime = System.currentTimeMillis() - start;
        logger.info(joinPoint.getSignature() + " executed in " + executionTime + "ms");
        return proceed;
    }
}

// Custom annotation
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface Loggable {
}
```

[[Equivalents of Spring AOP in NET]]
