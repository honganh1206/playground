---
id: Spring MVC (Web Framework)
aliases:
  - Spring MVC (Web Framework)
tags: []
---

# Spring MVC (Web Framework)

```java
// Controller
@Controller
@RequestMapping("/users")
public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/{id}")
    public String getUserDetails(@PathVariable Long id, Model model) {
        User user = userService.findUserById(id);
        model.addAttribute("user", user);
        return "user-details";  // View name
    }

    @PostMapping
    public String createUser(@ModelAttribute UserForm form, BindingResult result) {
        if (result.hasErrors()) {
            return "user-form";
        }
        userService.createUser(form);
        return "redirect:/users";
    }
}
```
