#include <stdio.h>

extern char state(char);
void print_state(unsigned char);

int main()
{
  for (int i = 0; i < 10; i++)
  {
    char s = state(i);
    print_state(s);
  }

  return 0;
}

void print_state(unsigned char s)
{
  for (int i = 0; i < 8; i++)
  {
    printf("%d", s & 128 ? 1 : 0);
    s <<= 1;
  }
  printf("\n");
}
