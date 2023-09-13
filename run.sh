#!/bin/bash

go build -o hotelBookings cmd/web/*.go && ./hotelBookings -dbname=HotelBookings -dbuser=rahuldhiman -cache=false -production=false
