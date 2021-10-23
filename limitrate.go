package googlesearch

import (
    "errors"
    "golang.org/x/time/rate"
)

// rate limiting request otherwise google started to block you IP 

errBlocked := errors.New("google block")

//Use golang rate function to limit the requets send by the client in the given time frame
var RateLimit - rate.NewLimiter(rate.Inf, 0)
