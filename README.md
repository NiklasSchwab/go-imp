<br />
<div align="center">
    <a href="https://www.h-ka.de/">
        <img src="https://upload.wikimedia.org/wikipedia/commons/1/13/HKA_Logo_Logoleiste_RGB.png" alt="Logo" width="50%">
    </a>
    <h1 align="center">A Go IMP</h1>
    <p align="center">
        <h3>Model-Based Software Development</h3>
        <br />
        Developed by Niklas Schwab
        <br />
        Supervised by Prof. Dr. Martin Sulzmann 
    </p>
</div>
<br />

<!-- ABOUT THE PROJECT -->
## About The Project
This repository contains the final project for the Model-Based Software Development lecture at Karlsruhe University of Applied Sciences. It implements a simple IMP in Go and provides helper functions to create abstract syntax trees (ASTs), which represent expressions, statements, or a combination of both as programs which then can be pretty printed, evaluated and type checked. For project details, please refer to the original project requirements of the lecture.

This project is based on the given code by Prof. Sulzmann.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- STRUCTURE & FILES -->
## Project Structure & Files
In order to avoid any complication with dependencies, this project only makes use of a single package. Nevertheless, the code is spread through multiple files:

| File           | Description                                                              |
|----------------|--------------------------------------------------------------------------|
| main.go        | Declares maps for Value- and Type-States and calls the example functions |
| types.go       | Contains functionality to handle typing                                  |
| values.go      | Contains functionality to handle values of certain types                 |
| expressions.go | Contains all code regarding expressions                                  |
| statements.go  | Contains all code regarding statements                                   |
| ast.go         | Contains helper functions to generate and "run" ASTs                     |
| examples.go    | Contains examples as functions, each defining and running "code" as ASTs |


<p align="right">(<a href="#top">back to top</a>)</p>
