/* Stubs for GCC atomic builtins that ccgo doesn't handle natively. */
typedef unsigned int uint32_t;
static inline uint32_t __atomic_add_fetch(volatile uint32_t *p, uint32_t v, int m) { *p += v; return *p; }
static inline uint32_t __atomic_sub_fetch(volatile uint32_t *p, uint32_t v, int m) { *p -= v; return *p; }
