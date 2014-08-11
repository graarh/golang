Current version limitations
===========================

The readme file below is the roadmap of this example.
I plan to make the complete useful software, not just
a few lines of sample code for expert system.

Currently, very early preview alpha version is available.
It can only small pack of features. In depth:

1. Rules format does not support many features
2. Relations hardcoded
3. Counters are not supported


Optimal cellphone billing plan
==============================

Disclaimer
----------

This software helps you to find the optimal cellphone billing plan. Upload your detailed bill for the few months, and the software calculates all your actions using different cellphone operators plans. As a result you will get the full table of the all plans calculated for you, sorted by expenses.

Limitations
-----------

1. Only calls are calculated, no sms/mss, internet, etc
2. No complex rouming rules, only country/regional bills

Input data
----------

Your bill can contain few different elements. 

* calls
* periodical charges

Calls should contain minimal required information

* in/out
* date and time
* duration in seconds
* phone number

Example:
``` csv
date and time, direction, phone number, duration
21/04/2014 22:31:05, out, +7 987 6543210, 55
```

Relations definition
--------------------

Conditions can check parameters from list:

1. Receiver (phone number information)
    * operator
    * region
    * additional number decomposition
    * additional time decomposition
2. Total duration of calls for period

Relations can affect parameters by adding predefined name value pair lines, or counters.

Counter collects all values of given field, if relation conditions is passed. And outputs name value pair to the output parameters.

Example:
``` yaml
-
    conditions:
        phone_country: [+7]            #phone country in this list
        phone_prefix: [951, 953, 800]  #phone prefix in this list
    add:
        region: local                  #add this pair to attributes
        operator: tele2
    counter:
        #collect 'duration' field of params into
        #counter 'local2local' if this relation
        #affects current call
        #adds local2local:number to output params
        local2local: duration
```

Rules definition
----------------

1. Price calculation structure
    * Price per second
    * Period in seconds
    * Affected interval of time
2. This price conditions list based on
    * operator
    * region
    * time parameter (day, hour, etc) in the range
    * duration in the range
    * additional number decomposition elements

Example:
``` yaml
-
    conditions:
        operator: [Tele2]   #operator value in list
        region: [Local]     #region value in list
        #<name> means that it is not list, but interval
        #nil means any value, e.g. [nil, 20] is from any to 20
        #counters can be used in relations conditions too
        <local2local>: [0,30] #first 30 minutes for free
    price:
        -
            price: 0
            period: 60  #time is rounded to minutes
            interval: [nil, nil] #affect all time duration
-
    conditiona:
        operator: [Tele2]
        #!before name means all except values in list
        !region: [Local]  
    price:
        -
            price: 1.00          #price per period
            period: 60           #period in seconds
            interval: [nil, 180] #first 3 minutes
        -
            price: 0.50
            period: 60
            interval: [181, nil] #all time from 3 minutes
```         

Calculations result
-------------------

In result of all calculations, you will get the data with all cellphone operators and their billing plans you have in system, with calculated total bill price for your entered calls. Sorted by operators. Operators are compared by their best bill. Inside operators sorted by bill.

Example:

|Operator   |Plan      |Amount|
|-----------|----------|------|
|Tele2      |Blue      |101.00|
|Tele2      |Green     |203.11|
|MTS        |One       |104.22|
|MTS        |Two       |405.44|

