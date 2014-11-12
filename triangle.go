package main

/*
  Triangle Project Code.

  Triangle analyzes the lengths of the sides of a triangle
  (represented by a, b and c) and returns the type of triangle.

  It returns:
    'equilateral'  if all sides are equal
    'isosceles'    if exactly 2 sides are equal
    'scalene'      if no sides are equal

  The tests for this method can be found in
    about_triangle_project.go
  and
    about_triangle_project_2.go
*/

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func Triangle(a, b, c int) (result string, err error) {
	var e errorString
	
	if ((a + b) <= c) ||
		((a + c) <= b) ||
		((b + c) <= a) {
		return "", e
	}
	
	if (a == 0) || (b == 0) || (c == 0) {
		return "", e
	}

	if (a < 0) || (b < 0) || (c < 0) {
		return "", e
	}
	
	if (a == b) && (a == c) {
		return "equilateral", nil
	}
	
	if (a == b) || (a == c) || (b == c) {
		return "isosceles", nil
	}
	return "scalene", nil
}
