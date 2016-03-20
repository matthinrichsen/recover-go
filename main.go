package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void
print_bin(char *in, char *out)
{
	FILE *input, *output;
   	char buff[1024];
	input = fopen(in, "r");
	output = fopen(out, "w");
	if( input == NULL || output == NULL )
	{
		return;
	}

	fgets(buff, 1024, input);
	while (!feof(input)) {
		fputs(buff, output);
		fgets(buff, 1024, input);
	}
  	fclose(input);
  	fclose(output);
	return;
}
*/
import "C"
import "flag"

var in = flag.String("input_file", ``, ``)
var out = flag.String("out", ``, ``)

func main() {
	flag.Parse()
	filename := C.CString(*in)
	outfile := C.CString(*out)
	C.print_bin(filename, outfile)
}
