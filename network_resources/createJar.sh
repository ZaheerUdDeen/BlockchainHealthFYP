#!/bin/bash
#
cd ..

cd java/

mvn install

cd target/

cp emergency-healthcare-0.0.1-SNAPSHOT-jar-with-dependencies.jar emergency-healthcare.jar


cp emergency-healthcare.jar ../../network_resources

cd ../../network_resources
