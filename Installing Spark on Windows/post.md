# Getting started with Apache Spark and Scala on Windows

## Introduction

This tutorial is intended for people who really need to run Apache Spark on windows. Usually it would be better to run it in a Linux VM or on Docker.

There are a few things that cause problems with Spark on windows. But using this way of installation I managed to minimize the impact.

## Getting the needed files

First, download the ***spark-1.6.0-bin-hadoop2.6.tgz*** file from the [Apache Spark Downloads][1] website. You should also install a program to open it, for example 7zip.

Next, download [hadoop2.6.0 for Windows][2]

This is all we will need.

## Putting everything together

Extract both files to their own folders.

Open the environment variables settings. Add your Spark *bin* folder to your **PATH** variable and create a new **SPARK_HOME** environment variable and set it to your Apache Spark base folder.

Next, create another environment variable named **HADOOP_HOME** and set it to your hadoop base folder.

### Preparing the configuration files

Open the conf folder in your Apache Spark base location. There should be a file called log4j.properties.template, change the name to log4j.properties.

For the sake of convenience you can open the file and change the value of ***rootCategory*** from **INFO** to **WARN**.

## Launching it

After trying different ways, this one seems to be the least error-prone:

To launch the master open the command line and type in:

***spark-class.cmd org.apache.spark.deploy.master.Master***

Next, open your browser, and navigate to: http://localhost:8080/

There you should see the address of your spark master, like *spark://192.168.1.27:7077*

Open another command line, and start one worker using:

***spark-class.cmd org.apache.spark.deploy.worker.Worker masteraddress***

That's it, spark is running.

To run the spark shell, use:

***spark-shell --master masteraddress***

To submit a packaged application jar, use:

***spark-submit --class mainclassname --master masteraddress packagedjarlocation***

***Have fun with it!***

[1]:https://spark.apache.org/downloads.html
[2]:https://www.barik.net/archive/2015/01/19/172716/
