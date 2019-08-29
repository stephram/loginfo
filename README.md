# loginfo - process a web log file

Assumes a working go1.12.x or later environment.

Build on OS X
```
./build.sh
```

Run on OS X with any of the following
```
./main
./main -v
./main -v -f <logfile-name>
```

The output should be something like the following.

```
$ ./main -v
{"level":"info","msg":"parsing './programming-task-example-data.log'","time":"2019-08-30T09:42:51+10:00"}
177.71.128.21, /intranet-analytics/
168.41.191.40, http://example.net/faq/
168.41.191.41, /this/page/does/not/exist/
168.41.191.40, http://example.net/blog/category/meta/
177.71.128.21, /blog/2018/08/survey-your-opinion-matters/
168.41.191.9, /docs/manage-users/
168.41.191.40, /blog/category/community/
168.41.191.34, /faq/
177.71.128.21, /docs/manage-websites/
50.112.00.28, /faq/how-to-install/
50.112.00.11, /asset.js
72.44.32.11, /to-an-error
72.44.32.10, /
168.41.191.9, /docs/
168.41.191.43, /moved-permanently
168.41.191.43, /temp-redirect
168.41.191.40, /docs/manage-websites/
168.41.191.34, /faq/how-to/
72.44.32.10, /translations/
79.125.00.21, /newsletter/
50.112.00.11, /hosting/
72.44.32.10, /download/counter/
50.112.00.11, /asset.css

IP address usage
2 : 168.41.191.9
1 : 72.44.32.11
2 : 168.41.191.43
3 : 72.44.32.10
1 : 79.125.00.21
3 : 177.71.128.21
4 : 168.41.191.40
1 : 168.41.191.41
2 : 168.41.191.34
1 : 50.112.00.28
3 : 50.112.00.11

URL usage
1 : http://example.net/faq/
2 : /docs/manage-websites/
1 : /moved-permanently
1 : /temp-redirect
1 : /newsletter/
1 : /asset.css
1 : /hosting/
1 : /download/counter/
1 : /blog/category/community/
1 : /faq/
1 : /faq/how-to-install/
1 : /asset.js
1 : /to-an-error
1 : /faq/how-to/
1 : /intranet-analytics/
1 : /this/page/does/not/exist/
1 : http://example.net/blog/category/meta/
1 : /blog/2018/08/survey-your-opinion-matters/
1 : /
1 : /translations/
1 : /docs/manage-users/
1 : /docs/


{"level":"info","msg":"Processed 23 log entries","time":"2019-08-30T09:42:51+10:00"}
unique IPs        : 11
most visited URLs : [/docs/manage-websites/ /blog/category/community/ /faq/]
most active IPs   : [168.41.191.40 50.112.00.11 72.44.32.10]
```
