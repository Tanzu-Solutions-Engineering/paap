package com.example.pivotal.demo;


import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController {

    @RequestMapping(value="/", produces = "application/json")
    public StringResponse index() {
        StringResponse response = new StringResponse(System.getProperty("java.version"));
        return response;
    }

}
