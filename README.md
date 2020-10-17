# PortaudioExampleforGo
GO언어기반 Portaudio사용설명서  portaudio빌드부터 사용까지(windows10)



Build For Windows (Portaudio)
-------------
1. Set-up Cmake, mingw64 (newst or old) and mingw64 add path for bin folder (ex: D:\mingw\bin)
2. download Portaudio source code and unzip
3. start Cmd or powershell // working directory (cd) for portaudio source code directort
4. cmd or powershell input [cmake portaudiopath -G "MinGW Makefiles" -Bportaudiopath -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=librarypath] if not error next
5. input mingw32-make  if not error next
6. input mingw32-make install  if not error next
7. cmd or powershell input [go get github.com/gordonklaus/portaudio] You'll see an error. next
8. keyboard windows key + R  ->  input %gopath%/src\github.com\gordonklaus\portaudio
9. edit  portaudio.go  
/*
#cgo pkg-config: portaudio-2.0
#include <portaudio.h>
extern PaStreamCallback* paStreamCallback;
*/
to
/*
#cgo !windows pkg-config: portaudio-2.0
#cgo windows  CPPFLAGS:   -ID:/build/library/include    <- D:/build/library is portaudio installed path  (Path used for cmake cmake -DCMAKE_INSTALL_PREFIX=librarypath )
#cgo windows  LDFLAGS:    -LD:/build/library/lib -lportaudio
#include <portaudio.h>
extern PaStreamCallback* paStreamCallback;
*/
10. copy portaudio.dll (D:/build/library/lib inside [In your case, the path may be different.]) to main.go(Source Code Using Library) directory
11. run!

한글 
-------------
1. cmake, mingw64를 설치합니다. (최신버전이든 아니든 상관없습니다. 오류만없다면) mingw64 압축해제한 폴더의 경로를 PATH에 입력해줍니다 (시스템변수의 PATH)
2. portaudio 소스코드를 다운로드받고 압축해제합니다.
3. cmd나 파워쉘을 킨후 cd 명령어로 압축해제한 portaudio 폴더로 이동합니다.
4. cmd나 파워쉘에 위에 적힌 명령어를 적습니다. 경로는 바꿔야합니다.
5. cmd나 파워쉘에 mingw32-make를 입력합니다.
6. cmd나 파워쉘에 mingw32-install을 입력합니다.
7. cmd나 파워쉘에 위의 명령어를 적습니다. 오류가있으면 (당연히 있는게 정상 windows기준)
8. windows + R키로 실행을 실행시키고  위의 경로를 적어줍니다.
9. portaudio.go의 내용을 수정합니다.위에적힌거에서 아래꺼 to 다음꺼
10. cmake에서 입력한 폴더경로중 librarypath에 해당하는 경로를 들어가줍니다. (폴더로)
11. lib폴더안에있는 libportaudio.dll을 실행시킬 main.go 파일이있는 폴더로 복사해줍니다.
12. 실행해보세요
