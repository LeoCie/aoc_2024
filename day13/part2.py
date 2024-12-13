import re
import math
from pulp import LpProblem, LpMinimize, LpVariable, lpSum, value, PulpSolverError
from constants import inputS

def parseInputToMachine(input):
	machines = []
	for machineString in input.split("\n\n"):
		machine = {"aButton": {}, "bButton": {}, "xPrize": 0, "yPrize": 0}
		for i,row in enumerate(machineString.split("\n")):
			if i == 0 or i == 1:
				button = {"xMove": 0, "yMove": 0}
				split = row.split("+")
				button["xMove"] = int(split[1][:2])
				button["yMove"] = int(split[2][:2])
				if i == 0:
					machine["aButton"] = button
				else:
					machine["bButton"] = button
			else:
				split = row.split("=")
				machine["xPrize"] = 10000000000000 + int(re.findall('\d+', split[1])[0])
				machine["yPrize"] = 10000000000000 + int(re.findall('\d+', split[2])[0])
		machines.append(machine)
	return machines
	
def solve_integer_problem(aButtonXMove, aButtonYMove, bButtonXMove, bButtonYMove, xPrize, yPrize):
    prob = LpProblem("Minimize_3a_plus_b", LpMinimize)
	
    a = LpVariable("a", lowBound=0, cat="Integer")
    b = LpVariable("b", lowBound=0, cat="Integer")

    prob += 3 * a + b, "Objective"

    prob += aButtonXMove * a + bButtonXMove * b == xPrize, "Constraint_1"
    prob += aButtonYMove * a + bButtonYMove * b == yPrize, "Constraint_2"

    try:
        prob.solve()
        if prob.status == 1:
            a_val = value(a)
            b_val = value(b)
            minimized_value = value(prob.objective)
            return a_val, b_val, minimized_value
        else:
            return 0,0,"not found"
    except PulpSolverError as e:
            return 0,0,"not found"

totalcost = 0
for machine in parseInputToMachine(inputS):
	a,b,cost = solve_integer_problem(machine["aButton"]["xMove"], machine["aButton"]["yMove"], machine["bButton"]["xMove"], machine["bButton"]["yMove"], machine["xPrize"], machine["yPrize"])
	if (cost == "not found"):
		continue
	totalcost = totalcost + (a*3 + b)
print(totalcost)