Cargo places downloaded sources in `.cargo` in home dir and build artifacts in `target/debug/deps` so *different programs use different versions of crates*.

However, this leads to `target` directory being quite large (Check with `du -shc .`)