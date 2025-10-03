shader_type canvas_item;

void vertex() {
	float xo = cos(TIME * 0.5 + VERTEX.y)*2.0;
	float yo = abs(sin(TIME * 0.5 + VERTEX.x)) * 0.5;
	
	VERTEX += vec2(xo, yo);
	
}

void fragment() {
	
}

//void light() {
//	// Called for every pixel for every light affecting the CanvasItem.
//	// Uncomment to replace the default light processing function with this one.
//}
