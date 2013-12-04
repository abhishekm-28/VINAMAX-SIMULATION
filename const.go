//This file contains all the constants

package vinamax 

import (
	"math"
)

const (
	//are these actually const? maybe as Particle property
	Dx   = 2e-9   // Size-x of Particle in m
	Dy   = 2e-9   // Size-y of Particle in m
	Dz   = 2e-9   // Size-z of Particle in m

	Gamma0 = 1.7595e11          // Gyromagnetic ratio of electron, in rad/Ts
	Mu0    = 4 * math.Pi * 1e-7 // Permeability of vacuum in Tm/A
	MuB    = 9.2740091523E-24   // Bohr magneton in J/T
	Kb     = 1.380650424E-23    // Boltzmann's constant in J/K
	Qe     = 1.60217646E-19     // Electron charge in C
)
