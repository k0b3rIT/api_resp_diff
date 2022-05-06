# Create new endpoint for lookup query (lookup with build info) the old one should be the same for triager client
Or Include remote name to response???

# AQL exception
 http://10.125.28.52:8668/gittleman/tickets/hive?branch=HDP-2.6.5.361
 http://10.125.28.52:8668/gittleman/tickets/hive.txt?branch=HDP-2.6.5.361

# SLOW response UNSupported OperationException
```
Caused by: java.lang.UnsupportedOperationException
at java.util.AbstractMap.put(AbstractMap.java:209)
at com.hortonworks.sustaining.triager.service.GittlemanService.lambda$getBranchesContainingTicket$3(GittlemanService.java:58)
```
 http://10.125.28.52:8668/gittleman/lookup/HIVE-4167?showEmpty=true 

 http://10.125.28.52:8668/gittleman/lookup/HIVE-3400?showEmpty=true

# Commit appears multiple times ??

/gittleman/lookup/commithash/e20e5b06e3d8184376ab7a53835f3405433e9ee9
```
     {
         "branchInfo": {
-            "branchName": "cdh/CDH-7.1.7.20",
-            "component": "knox"
+            "component": "knox",
+            "remoteURI": "https://github.infra.cloudera.com/CDH/knox.git",
+            "branchName": "CDH-7.1.8.x"
         },
         "commitInfo": [
             {
                 "commitUri": "https://github.infra.cloudera.com/CDH/knox/commit/e20e5b06e3d8184376ab7a53835f3405433e9ee9",
                 "message": "KNOX-598: Concurrent JDBC clients via KNOX to Kerberized HiveServer2 causes HTTP 401 error (due to Kerberos Replay attack error)",
                 "fullMessage": "KNOX-598: Concurrent JDBC clients via KNOX to Kerberized HiveServer2 causes HTTP 401 error (due to Kerberos Replay attack error)\n",
                 "commitTime": 1442241792
+            },
+            {
+                "commitUri": "https://github.infra.cloudera.com/CDH/knox/commit/e20e5b06e3d8184376ab7a53835f3405433e9ee9",
+                "message": "KNOX-598: Concurrent JDBC clients via KNOX to Kerberized HiveServer2 causes HTTP 401 error (due to Kerberos Replay attack error)",
+                "fullMessage": "KNOX-598: Concurrent JDBC clients via KNOX to Kerberized HiveServer2 causes HTTP 401 error (due to Kerberos Replay attack error)\n",
+                "commitTime": 1442241792
             }
-        ],
-        "branchName": "cdh/CDH-7.1.7.20"
+        ]
     }
```

# Works with new model on sandbox but error on prod!!! WHAT????

# HTTP500 on prod endpoint, but ok, for sandbox 
http://10.125.29.146:8668/gittleman/lookup/commithash/c255a4fa8f7b8de2c206ff2d2f5daa7b3ee240bb

# Timeout on prod endpoint but ok for sandbix
http://10.125.29.146:8668/gittleman/search/hive/HIVE-3400
http://10.125.29.146:8668/gittleman/search/hive/HIVE-3839

# Request processing failed; nested exception is java.lang.NullPointerException
http://10.125.29.146:8668/gittleman/log/hive/HDP-2.6.5.361
http://10.125.29.146:8668/gittleman/log/hbase/Branch-2.2