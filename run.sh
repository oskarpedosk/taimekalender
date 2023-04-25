#!/bin/bash

cd back-end
go run ./cmd &
cd ..
cd front-end
npm start