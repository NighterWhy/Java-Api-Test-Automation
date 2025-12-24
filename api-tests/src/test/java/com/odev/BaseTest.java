package com.odev;

import io.restassured.RestAssured;
import io.restassured.filter.log.RequestLoggingFilter;
import io.restassured.filter.log.ResponseLoggingFilter;
import org.junit.BeforeClass;

public class BaseTest {
    @BeforeClass
    public static void setup() {
        RestAssured.baseURI = "http://localhost:8080";
        
        
        RestAssured.filters(new RequestLoggingFilter(), new ResponseLoggingFilter());
        
    }
}