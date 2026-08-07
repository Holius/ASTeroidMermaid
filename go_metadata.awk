#! /bin/awk -f

BEGIN {
  print "package filedata"
  print "var GoFileDataExpected = []GoMetadata{"
}

BEGINFILE {
  IMPORT_RAN=0
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



# script now handles this case
# func mergeMap[T any](base, override map[string]T) map[string]T {
$1 ~ "^func$" {
  if (IMPORT_RAN == 0) {
    print "Imports: []string{},"
    printf "Functions: []string{\n"
   IMPORT_RAN=1
  }

   # check for generic
   found = index($2, "[")
   if (found != 0) {
     # currently, nothing is printed to distinguish this function as "generic"
     printf "\"%s\",\n",substr($2, 0, found - 1) 
     next
   }

  
  funcPos=2  
  # check for method on struct
  if (substr($2, 1, 1) == "(") {
    funcPos=4
    found = index($3, ")")
    struct=substr($3, 0, found - 1) 
    gsub(/\*/, "", struct)
    printf "\"%s.", struct
  }

  found = index($funcPos, "(")
  # output differently if method on strut
  if (funcPos == 4) {
    printf "%s\",\n",substr($funcPos, 0, found - 1) 
  } else {
    printf "\"%s\",\n",substr($funcPos, 0, found - 1) 
  }
}

ENDFILE {
    if (IMPORT_RAN == 1) {
      printf "},},\n"
    } else {
      print "Imports: []string{},"
      print "Functions: []string{},"
      printf "},\n"
    }
}

END {
    printf "}\n"
}