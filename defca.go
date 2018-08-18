//
// Deferring function calls
//

// Go supports the notion of deferring a function call. Placing the keyword defer before a
// function call has the interesting effect of pushing the function unto an internal stack,
// delaying its execution right before the enclosing function returns.
