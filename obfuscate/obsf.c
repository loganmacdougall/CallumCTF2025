#include <stdlib.h>
#include <stdio.h>


char m[100000];
char k[] = "Solve this riddle and find the code";

void x()
{
  int mm = 0xffff;

  for (int i = 0; m[i] != '\0'; i++)
  {
    if (k[i % mm] == '\0')
    {
      mm = i;
    }

    m[i] ^= k[i % mm];
    if (m[i] == '\0')
    {
      m[i] ^= k[i % mm];
    }
  }
}
void r(char* n) {
  char e;
  int i;
  FILE *f = fopen(n, "r");
  for (int i = 0; (e = fgetc(f)) != EOF; i++) {
    m[i] = e;
  }
  m[i] = 0;
  fclose(f);
}

int main(int c, char** v)
{
  int i;

  r("c_to_c2.c");
  x();

  printf("\n\n%d (should be zero)\n\n", m[sizeof(m)-1]);

  FILE *f = fopen("bytes.txt", "w");
  fprintf(f, "{");
  for (i = 0; m[i] != '\0'; i++)
  {
    if (i != 0)
      fprintf(f, ",");
    fprintf(f, "%d", m[i]);
  }
  fprintf(f, "};\n");
  fclose(f);

  x();
  printf("%s\n", m);
}