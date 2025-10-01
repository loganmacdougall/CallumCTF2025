#include <stdlib.h>
#include <stdio.h>


char m[] = "python3 -c \"import sys,codecs,base64;m=\'correct\'+chr(33);d=base64.decodebytes(codecs.decode(b\'x\\x9c\\x0bqw+\\x8d2\\x0e\\xcb\\x88\\xaa\\xf2KKt\\xb7,\\x8f0J/Ov\\xb3\\xcc\\x8a\\x0c75\\x880\\xf63\\x081v\\xb4\\xe5\\x02\\x00\\xe91\\x0bc\',\'zlib_codec\'));f=open(\'HXVNFIEFJEKS.txt\',\'w\');f.write(\'No argument provided\') if len(sys.argv)<2 else (f.write(m.title()) if sys.argv[1].encode(\'utf-8\')==d else f.write(f\'in{m}\'.title()))\"";
char k[] = "last one before the end";

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

char p[512];

int main(int c, char** v)
{
  x(m, k);

  printf("%s\n", m);

  printf("{");
  for (int i = 0; m[i] != '\0'; i++)
  {
    if (i != 0)
      printf(",");
    printf("%d", m[i]);
  }
  printf(",0}\n");

  // system(p)
}