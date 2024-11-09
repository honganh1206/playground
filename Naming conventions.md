# Naming conventions

Go follows simple and consistent naming conventions that emphasize clarity and readability. Here's a summary of the key conventions:

### 1. **Package Names:**
   - Use short, all-lowercase names without underscores or mixed caps.
   - Package names should be descriptive but concise (e.g., `fmt`, `http`).
   - The package name is usually the same as the directory name.

### 2. **Variable and Function Names:**
   - Use **camelCase** for naming variables and functions (e.g., `userName`, `getData`).
   - Keep names descriptive but concise, avoiding overly long names.

### 3. **Constants:**
   - Use **PascalCase** for naming exported constants (e.g., `MaxValue`).
   - Use **ALL_CAPS** with underscores for unexported constants (e.g., `max_value`), but this is rare in Go.

### 4. **Type Names:**
   - Use **PascalCase** for naming types (e.g., `User`, `Server`).
   - Type names should be nouns that describe the entity (e.g., `Buffer`, `Reader`).

### 5. **Interfaces:**
   - Name interfaces based on the behavior they describe, typically using the **-er** suffix (e.g., `Reader`, `Writer`).
   - If an interface has a single method, name it after that method (e.g., `Reader` for `Read()`).

### 6. **Exported vs. Unexported Names:**
   - **Exported (public) names** should start with an uppercase letter (e.g., `ExportedFunction`).
   - **Unexported (private) names** should start with a lowercase letter (e.g., `internalFunction`).

### 7. **Acronyms:**
   - When using acronyms in names, keep them consistent in casing (e.g., `HTTPRequest`, `XMLParser`).
   - Stick to either all uppercase or all lowercase for the acronym part.

### 8. **Receiver Names in Methods:**
   - Use short and meaningful names for method receivers, often one or two letters (e.g., `func (s Server) Start()`).

### 9. **Error Variables:**
   - Name error variables starting with `err` (e.g., `err`, `errNotFound`).
   - For sentinel errors (predefined errors), use **Err** as a prefix (e.g., `ErrInvalidInput`).

### 10. **File Names:**
   - Use all lowercase with underscores to separate words if necessary (e.g., `user_service.go`, `http_server.go`).

### 11. **Test Files:**
   - Test file names should end with `_test.go` (e.g., `user_service_test.go`).
   - Test function names should start with `Test` followed by the name of the function being tested (e.g., `TestUserService`).

### 12. **Avoid Redundant Naming:**
   - Avoid repeating the package name in identifiers (e.g., prefer `user.Get()` over `user.UserGet()`).

These conventions aim to make Go code more idiomatic, readable, and maintainable across different projects and teams.