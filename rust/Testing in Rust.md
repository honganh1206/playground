Tests are not necessarily run in the same order they are declared, as *Rust is designed for writing concurrent code - code that can be run across multiple threads*, so the tests might be run on parallel.o

To run on a single thread, use `cargo test -- --test-threads=1`