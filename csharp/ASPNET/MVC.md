# Models - Represent the data of the app
- Used with Entity Framework Core (EF Core) - An Object-Relational Mapping (ORM) framework
# Views - Display the app's interface
- View templates should not
	- Do business logic (for Models)
	- Interact with DB (for Controllers)
- View templates should only work with **Data from Controllers**
# Controllers
- Handle requests
- Retrieve model data
- Call view templates
## Create a controller
- Add endpoints
- The `Index` method is called by default

```ad-info
title: The point of MVC
The UI logic belongs in the view. Input logic belongs in the controller. Business logic belongs in the model.

```

# Entity Framework Core
- MVC provides the ability to pass strongly typed model objects to a view.

```ad-warning
title: Modifying data in an `HTTP GET` method is a security risk.
Modifying data in an `HTTP GET` method also violates HTTP best practices and the architectural [REST](http://rest.elkstein.org/) pattern, which specifies that GET requests shouldn't change the state of your application

```

```ad-tip
title: The DRY Principle
One of the design tenets of MVC is [DRY](https://wikipedia.org/wiki/Don%27t_repeat_yourself) ("Don't Repeat Yourself"). ASP.NET Core MVC encourages you to specify functionality or behavior only once, and then have it be reflected everywhere in an app. 

This reduces the amount of code you need to write and makes the code you do write less error prone, easier to test, and easier to maintain.

```

## Validation rules
- Having validation rules automatically enforced by ASP.NET Core helps **make your app more robust.** It also ensures that you can't forget to validate something and inadvertently let bad data into the database.

# Dependency Injection Service Lifetime
## Transient
- Definition
	- Whenever we want a new implementation, we create a new object
	- Every time a service is requested, a new object is created
- Pros
	- Safest
	- Never have to re-use existing object
## Scoped
- Definition
	- Depend on the HTTP
	- Whenever a service is first created, a new object is made and that same object is used whenever that service is requested
## Singleton
- Definition
	- One implementation is created for the lifetime of the application

## ViewBag - Wrapper for ViewData

- Transfer data from the Controller to View in a **dynamic** way => Ideal for temporary data which is not in the model
- Any number of properties and values can be assigned to ViewBag
- ViewBag's lifetime only lasts during the current Http request
- ViewBag's value will be null if redirection occurs

## ViewData
- Transfer data from the Controller to View -> Same approach as ViewBag
- Derived from ViewDataDictionary
- ViewData must be **type-cast** before use 
- ViewData's lifetime only lasts during the current Http request
- ViewData's value will be null if redirection occurs


```ad-note
title: ViewBag and ViewData
ViewBag internally inserts data into ViewData dictionary => The key of ViewData and property of ViewBag must **NOT** match.

```

## TempData
- Used to store data between two consecutive requests
- Internally use Session to store data => A short-lived session
- Must be type-cast before use
- Can only be used to store one-time messages e.g. errors, validations

## View Model / Strongly typed views
- Models specifically used for view

```ad-warning
title: Navigation properties when inserting a record
Never populate data into a navigation property when inserting a record in the Db

```
