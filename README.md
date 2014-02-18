# Payroll Calculator

The payroll scheduler calculates the pay date for a salary that is paid from January through to December of any given year. The user can either supply the payment date as the day of the month, i.e. always on the 30th day; or a payment frequency, i.e. every 2 weeks. Both of the payment schedules are able to correctly accept and negotiate a list of public holidays supplied as a file containing json holiday dates.

## Example usage:
```
payroll_scheduler --year 2013 --day 30
payroll_scheduler --pay_frequency "2 week" --starting "1/1/2013"
payroll_scheduler --pay_frequency "4 week" --starting "1/1/2013" --public_holidays "./holidays.json"
```

To get the help with all options:
```
payroll_scheduler --help
```


## Tests:

To run the test:
````
go test
````

## Requirements (all):
* The payment date and year must be supplied to the scheduler from the command line.
* All payment dates must fall within the year being processed.
* All payment dates must be made during a standard working week; e.g. Monday through Friday.
* Any payment date that falls on a weekend should be made on the preceding Friday.
* Any payment date that falls on a public holiday should be made on the preceding Friday, unless the public holiday is a Friday.
* If the public holiday is a Friday then the payment should be made on the first preceding day that is not a public holiday or weekend.
* Each payment date must be printed to the command line in the following format; Monday, January 1, 2012 and include a trailing new line.

## Requirements (DayScheduler)
* There must be one payment only per calendar month. No calendar month should have less or more than one payment.
* If the payment date falls on an invalid day, such as a weekend or public holiday, and the preceding available day is outside of the current month, then the payment should be made on the first available day after the current date.
* If the payment date does not exists, such as 2/31/2013 (or any month with fewer than 31 days), then the first prior valid day of the month should be used. This can occur if the scheduler is required to pay on the 31st day of every month.%
