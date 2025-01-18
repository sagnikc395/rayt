- shared_prt<type> -> pointer to some allocated type, with reference-counting semantics. Every time we assign its value to another shared pointer, the reference count is incremented. As shared pointers go out of scope, the reference count is decremented. Once count goes to 0, object i safely deleted.

