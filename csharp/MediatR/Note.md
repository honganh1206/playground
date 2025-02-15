# MediatR

- Definition: An **“in-process”** Mediator implementation, that helps us build CQRS systems
- All communication between the user interface and the data store happens via MediatR.

## How it works
- MediatR Requests are **simple request-response style messages** where a single request is **synchronously**  (synchronous from the request point of view, not C# internal async/await) handled by a single handler
- Two types of request
	- Return a value (read/queries)
	- NOT return a value (write/commands)

## Notifications with MediatR

```ad-warning
The interface `INotification` does not return a value, as they work on the fire-and-forget principle like publishers

```cs
public sealed record CompanyDeletedNotification(Guid Id, bool TrackChanges) :
INotification;
```

## Behaviors with MediatR
Behaviors are very similar to ASP.NET Core middleware in that they:
- Accept a request
- Perform some action
- (Optionally) pass along the request.


```ad-info
Some developers have a preference for **using fluent validation over data annotation attributes**. In that case, **behaviors are the perfect place** to execute that validation logic.

```

### Validate DTOs from the client with Fluent Validation
The FluentValidation library allows us to easily define very rich custom validation for our classes => Use validations for Commands

### Decorators with MediatR PipelineBehavior
- CQRS represents **a request-response pipeline** => Introduce additional behaviors around each request through the pipeline (similar to middlewares)