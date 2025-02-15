---
tags:
  - "#study"
cssclasses:
  - center-images
---
There might be *a small percentage of real world use cases* when we *only need the 1st element* from the lookup, but we are unsure of *how many elements would match our condition*. Furthermore, we might just *take one of the matching element and ignore the rest* => Cost inefficient as we have to traverse through the whole array

We usually do single element lookup with its **unique identifier**, so using alternatives like `Single()` or `SingleOrDefault()` makes much more sense

[[One case in which we should replace FirstOrDefault() with First() or Single()]]

