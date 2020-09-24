
IF %1.==. GOTO No1

  git add *
  git commit -m %1
  git push origin master

GOTO End1

:No1
  ECHO No param 1
GOTO End1


:End1
