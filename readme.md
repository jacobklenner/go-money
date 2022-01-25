MONEY

basic implementation of money type in golang

money is represented by a value, and precision.

e.g. 

money{
    value: 1000
    precision: 0
}

represents 1000 units of the specified currency

money{
    value: 1000
    precision: 2
}

represents 10.00 units of the specified currency

money can be operated on, as you would expect

e.g. addition

// 10.00
a := money{
    value: 1000
    precision: 2
}

// 5.0
b := money{
    value: 50
    precision: 1
}

// 10.00 + 5.0
c := a.Add(b)

print(c.Value) // 15
print(c.Precision) // 0

also supported is subtraction, multiplication and division.

division returns the largest whole number of times the numerator fits into the denominator

default behaviour is set to return the lowest precision possible

e.g. above response represents 15, not 15.0 or 15.00

open todo's
    > ability to configure default precision
    > extend to represent different currencies, and their units e.g. EUR currency, with cents and euro as units
    > add support for multiplication by a percentage (also to be represented by {value, precision} combination)
    > make it a bit more readable
