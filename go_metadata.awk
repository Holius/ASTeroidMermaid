#! /bin/awk -f

BEGIN {
  IMPORT_RAN=0
}

# the FILENAME variable is not defined in the BEGIN block,
# so we check the first record
NR == 1 {
    printf "{\nFile: \"%s\",\n",FILENAME
}

$1 == "package" {
   printf "Package: \"%s\",\n",$2 
}

$1 == "import" {
   if ($2 == "(") {
      print "Imports: []string{"
      getline;
      while ($1 != ")") {
        if ($0 == "" || $1 == "//") {
        } else {
            if (NF == 1) {
              printf "%s,\n",$1;
            } else {
              printf "%s,\n",$2;
            } 
        } 
        getline;
      }
      printf "},\n"
   } else {
     printf "Imports: []string{%s,},\n",$NF
   }
   IMPORT_RAN=1
   printf "Functions: []string{\n"
}



# this script does not handle Generics
# func mergeMap[T any](base, override map[string]T) map[string]T {
$1 == "func" {
  if (IMPORT_RAN == 0) {
    printf "Functions: []string{\n"
   IMPORT_RAN=1
  }
  for (i = 1; i < NF + 1; i++) {
    #print $i
    # regex match for function declaration like "TestThis(t"
    if (match($i, /^[A-Za-z][a-zA-Z_0-9]+\(/) != 0) {
      found = index($i, "(")
      printf "\"%s\",\n",substr($i, 0, found - 1) 
      break
    }
  }
}

END {
      if (IMPORT_RAN == 1) {
        printf "},},\n"
      } else {
        printf "},\n"
      }
}