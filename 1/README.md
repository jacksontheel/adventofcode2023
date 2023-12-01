# Part 1
This was fairly simple. I was able to use a regular expression to capture all that I needed to. The `FindStringSubmatchIndex` function did most of the heavy lifting.

Using a regular expression as simple as `\d` might clue me into the fact that regular expressions for this problem are absolutely overkill. For the advent calendar, I'm not writing for performance, this is code to work once, and forget about.

# Part 2
Initially I thought this would be a fairly simple addition, I just needed to change my regular expression a bit to
```
\d|one|two|three|four|five|six|seven|eight|nine
```
Then, I set up a simple function with a switch statement to transform a string like `one` into the int `1`. This was close, but giving me the wrong answer.

I reasoned pretty fast that it had to do with overlap. My regular expression was taking a string like `twone` and giving me just a single value in my response slice, when I wanted it to capture both `one` and `two`

Unfortunately, there just wasn't an option to capture overlaps in the regexp `FindAllString` function. So I had to build my own using the more primitive `FindStringSubmatchIndex` in a loop. The only tweak I had to make to my own implementation was to set the start index of my next search one further back if it had just consumed a string number like `two`. This way, after consuming the two in `twone`, the next search would be on the string `one`, instead of `ne`, which wouldn't capture me anything.