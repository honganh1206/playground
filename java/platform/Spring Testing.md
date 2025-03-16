---
id: Spring Testing
aliases:
  - Spring Testing
tags: []
---

# Spring Testing

Comprehensive testing support:

```java
@SpringBootTest
class UserServiceIntegrationTest {

    @Autowired
    private UserService userService;

    @Autowired
    private UserRepository userRepository;

    @BeforeEach
    void setUp() {
        userRepository.deleteAll();
    }

    @Test
    void shouldFindUserById() {
        // Given
        User savedUser = userRepository.save(new User("testuser", "test@example.com"));

        // When
        User foundUser = userService.findUserById(savedUser.getId());

        // Then
        assertNotNull(foundUser);
        assertEquals("testuser", foundUser.getUsername());
    }
}

@WebMvcTest(UserController.class)
class UserControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockBean
    private UserService userService;

    @Test
    void shouldReturnUserDetails() throws Exception {
        // Given
        User user = new User("testuser", "test@example.com");
        user.setId(1L);
        when(userService.findUserById(1L)).thenReturn(user);

        // When / Then
        mockMvc.perform(get("/users/1"))
            .andExpect(status().isOk())
            .andExpect(view().name("user-details"))
            .andExpect(model().attributeExists("user"))
            .andExpect(model().attribute("user", hasProperty("username", is("testuser"))));
    }
}
```
