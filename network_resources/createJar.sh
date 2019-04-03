#!/bin/bash
#
cd ..

cd java/

mvn install

cd target/

cp blockchain-healthcare-0.0.1-SNAPSHOT-jar-with-dependencies.jar blockchain-healthcare.jar


cp blockchain-healthcare.jar ../../network_resources

cd ../../network_resources
