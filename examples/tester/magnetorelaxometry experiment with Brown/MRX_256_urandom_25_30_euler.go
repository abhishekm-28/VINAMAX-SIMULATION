//This example shows the agreement between the Dipole approximation Method implementation and the brute force implementation of the magnetostatic interaction. The same problem is also solved without taking this interaction into account so to show that it is of importance in this system.
package main

import (
	. "github.com/JLeliaert/vinamax"
	"fmt"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 2e-6 m	
	World(0,0,0,2e-6)
	//Adds a cube to the center of the world with side 1e-8 m
	test := Cube{S:2e-6}

	//additionally calculate Brownian rotation
	BrownianRotation = true
	//this requires a randomn number for the anisotropy dynamics
	Setrandomseed_anis(2)
	//set viscosity environement test (e.g. water 1mPas)
	test.Setviscosity(1e-3)

	//the particle have radius 25 nm and a hydrodynamic radius (core + coating together) of 30 nm (tauN = 0.03 s TauB = 10 µs)
	Particle_radius(25e-9)
	Particle_radius_h(30.e-9)

	//Adds 256 particles with radii defined above to the cube with viscosity 1 mPas
	test.Addparticles(256)
	fmt.Println("particles added")
	//Don't calculate the demagnetising field
	Demag=false

	//saturation magnetisation 860 000 A/m
	Msat (860e3)

	//Adaptivestep = true
	//timestep : 2ps
	Dt = 2e-12
	//initialise time at zero
	T = 0.
	Brown=true
	Temp = 300.
	//Set a randomseed for the thermal field
	Setrandomseed(2)
	//The Gilbert damping constant =0.1
	Alpha = 0.1
	//anisotropy constant=10 000 J/m**3
	Ku1 = 10000 

	//anisotropy axis is random for every particle 
	Anisotropy_random()
	
	//initialise the magnetisation along the z direction
	M_uniform(0,0,1)
	
	//curently anisodynamics only works with Euler
	Setsolver("euler")

	//write output every 0.1 µs 
	Output(0.1e-6)
	
	//fmt.Println("dt:   ", Dt)
	//run for 100 µs
	Run(100e-6)
}
