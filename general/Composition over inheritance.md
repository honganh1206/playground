Composition over inheritance is a design principle where *classes achieve polymorphic behavior and reuse code by containing instances of other classes (composition) rather than by inheriting from them (inheritance)*.

It emphasizes "has-a" relationships over "is-a" relationships to promote flexibility, reduce coupling, and improve testability.

Imagine an `Animal` class. You want some animals to fly.

```js
Animal
  - walk()
  - eat()
  |
  +-- Bird (inherits Animal)
  |     - fly()
  |
  +-- Ostrich (inherits Bird) // Problem: Ostrich cannot fly!
  |     - fly() { // Must override to do nothing or throw error }
  |
  +-- Penguin (inherits Bird) // Problem: Penguin cannot fly!
        - fly() { // Must override to do nothing or throw error }
```
This hierarchy becomes problematic because not all `Bird`s can fly. You end up with a rigid structure where methods are overridden to do nothing, or you're forced to create complex, less intuitive hierarchies (e.g., `FlyingBird`, `FlightlessBird`).

**Solution**: Extract the "flying" behavior into a separate interface and concrete implementations.

```java
// Interface for flying behavior
interface FlyBehavior {
    void fly();
}

// Concrete flying behaviors
class CanFly implements FlyBehavior {
    public void fly() { System.out.println("Flying high!"); }
}

class CannotFly implements FlyBehavior {
    public void fly() { System.out.println("Cannot fly."); }
}

class RocketPoweredFly implements FlyBehavior {
    public void fly() { System.out.println("Flying with a rocket!"); }
}

// Animal class composes a FlyBehavior
class Animal {
    // Composition: Animal "has-a" FlyBehavior
    protected FlyBehavior flyBehavior;

    public void setFlyBehavior(FlyBehavior fb) {
        this.flyBehavior = fb;
    }

    public void performFly() {
        flyBehavior.fly(); // Delegate the behavior
    }
    // other animal methods like walk(), eat()
}

// Specific animals
class Sparrow extends Animal {
    public Sparrow() {
        this.flyBehavior = new CanFly(); // Composes a flying behavior
    }
}

class Ostrich extends Animal {
    public Ostrich() {
        this.flyBehavior = new CannotFly(); // Composes a non-flying behavior
    }
}

class RobotBird extends Animal {
    public RobotBird() {
        this.flyBehavior = new RocketPoweredFly(); // Can even change behavior dynamically
    }
}
```

**Benefits:**
*   **Flexibility:** Easily swap flying behaviors at runtime or during construction.
*   **Reduced Coupling:** `Animal` doesn't need to know *how* to fly, only that it *can* delegate to a `FlyBehavior`.
*   **Testability:** Each `FlyBehavior` can be tested independently.
*   **Avoids "Liskov Substitution Principle" violations:** An `Ostrich` still *is-a* `Bird`, but you're not forced to implement a meaningless `fly()` method.
