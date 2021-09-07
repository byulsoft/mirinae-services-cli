# mirinae-services-cli

🗺 개발 환경에서 배포서버로 배포했을 시, 로깅 및 서비스 상태 조회의 불편사항을 느끼시는 것을 고려하여 만든 CLI 입니다. 🗺
   1. 기본 모듈 파악 및 AWS Jenkins EC2 와의 커넥션 상태 확인
   2. MobaXterm 의 커넥션 제한으로 불편함으로 인한 터널링 지원 및 터널링 상태 확인
   3. Kubernetes 클러스터 내부 서비스 및 자원 상태 확인 [배포 후 확인 용이성]
   4. 쿠버네티스 자원에 대한 로그 지원 [배포 후 디버깅 용이성]
   
   해당 CLI 는 go 언어로 만들었습니다

<br><br>

## 준비사항
mirinae-services-cli 는 `ssh` 가 있어야 동작합니다. mac 환경에서는 `ssh`, `ssh-client` 모듈이 전부 존재하고 있을 것입니다.

<br>

- 윈도우 운영체제에서 SSH 와 SSH Client 를 받는 법은 다음과 같습니다.
   - https://forbes.tistory.com/910
- 윈도우 환경에서는 Git Bash 를 기준으로 명령어를 구성 및 파싱하였으므로 Git Bash 를 사용해야 동작합니다.
   - https://git-scm.com/downloads

<br><br>


## 설치

### mirinae-services-cli for window
---
1. https://github.com/soominJung0413/mirinae-services-cli/releases 에 접속하여 .zip 파일을 내려받고 압축을 해제합니다.

![image](https://user-images.githubusercontent.com/67881815/132282632-e6517b74-0b89-4020-bf8d-bc52d4c59a71.png)

2. `mirinae-services-cli` 는 내부적으로 키를 찾는 데에 `MIRINAE_PATH` 를 사용합니다. (EC2 키는 보안 상 프로세스에 넣지 않았습니다.) 환경변수를 편집합니다.

![image](https://user-images.githubusercontent.com/67881815/132282937-568c81f7-bdc2-4d3e-92b3-70df621a6900.png)

3. 프로그램에서 git bash 에 맞게 경로를 파싱하므로, 윈도우 기준으로 `jenkins 서버의 .pem 키` 파일을 놓을 폴더를 `MIRINAE_PATH` 로 설정하면 됩니다.

![image](https://user-images.githubusercontent.com/67881815/132283017-b0d57f54-2834-4ef3-a063-0abd57ada1e8.png)

4. `git bash` 를 열어 해당 명령을 실행해봅니다.

```bash
printenv | grep MIRINAE
## MIRINAE_PATH=C:\Users\JungSooMin\Documents\AWS\byulsoft-pemkey
```

5. `MIRINAE_PATH` 에 가서 `jenkins 서버의 .pem 키` 를 놓고 이름을 `byul-jenkins.pem` 로 변경 한 후, 하위에 bin 폴더를 생성합니다.

![image](https://user-images.githubusercontent.com/67881815/132283250-09b5175c-495c-4942-90ac-3bc0a6495e12.png)

6. 시스템 환경변수 중 `PATH` 변수에 `%MIRINAE_PATH%\bin\` 를 추가합니다.

![image](https://user-images.githubusercontent.com/67881815/132283463-f4568741-80ff-4bd7-af58-b2389de46809.png)

7. `mirinae-services.zip` 파일을 `%MIRINAE_PATH%\bin\` 에 해당하는 경로로 옮기고 압축을 해제합니다.

![image](https://user-images.githubusercontent.com/67881815/132283729-4f7befeb-d36d-4c03-94d0-73cbd4c45e98.png)

8. `git bash` 를 열어 해당 명령을 실행 해봅니다.
```bash
mirinae-services -h
```
![image](https://user-images.githubusercontent.com/67881815/132283854-b3024020-c07c-455c-9848-8adee682580f.png)

9. `MIRINAE_PATH` 가 올바르게 등록되어 있고, 해당 경로에 `byul-jenkins.pem` 키가 있다면, 해당 명령은 성공합니다. `KeyPath` 가 파싱된 `MIRINAE_PATH` 에 `byul-jenkins.pem` 를 붙인 경로입니다.
```bash
mirinae-services require
```
![image](https://user-images.githubusercontent.com/67881815/132284022-ab24be4f-b4fc-450c-bc92-60b86c8cda34.png)


<br><br>


### mirinae-services-cli for mac, linux
---
1. go 로 빌드 된 파일인 mirinae-services.exe 를 받습니다. (크기 약 8M) => https://github.com/soominJung0413/mirinae-services-cli/releases
```bash
## .tar 파일 다운로드
curl -LO https://github.com/soominJung0413/mirinae-services-cli/releases/download/mirinae/mirinae-services.tar

tar -xvf mirinae-services.tar
```
3. `mirinae-services-cli` 는 내부적으로 키를 찾는 데에 `MIRINAE_PATH` 를 사용합니다. (EC2 키는 보안 상 프로세스에 넣지 않았습니다.) 다음의 커멘드를 실행합니다.
```bash
### zsh 를 사용한다면 .zshrc 를 엽니다.
vi ~/.bash_profile

### ~/.bash_profile
export MIRINAE_PATH=[원하는 폴더 경로]
```
3. 소스 파일을 쉘에 적용하고 터미널을 다시 키고 `MIRINAE_PATH` 적용 여부를 확인합니다.
```bash
### zsh 를 사용한다면 .zshrc 를 전달합니다.
source ~/.bash_profile

printenv | grep MIRINAE_PATH
## MIRINAE_PATH=/Users/jeongsumin/Documents/AWS/byul/byulsoft-pemkey/
```
4. `/usr/local/bin` 로 파일을 옮기고 터미널을 다시 열어 등록이 됬는지 확인합니다.
```bash
sudo mv mirinae-services /usr/local/bin/

mirinae-services -h
```
![image](https://user-images.githubusercontent.com/67881815/132218847-41a436b2-8981-4491-b815-a7b13bfc9827.png)

5. EC2 보안 정책에 따라서 키 파일의 mod 를 400 으로 변경하고, MIRINAE_PATH 에 복사합니다
```bash
chmod 400 [Jenkins 서버].pem && cp [Jenkins 서버].pem ${MIRINAE_PATH}/

ls -al ${MIRINAE_PATH}
```
6. 터미널을 다시 열고 실행 해봅니다
```bash
mirinae-services require
```
7. 해당 문구가 출력된다면, 설치 및 EC2 키  정상적으로 완료 된 것 입니다.
![image](https://user-images.githubusercontent.com/67881815/132219011-8a1a557b-47f4-495d-b667-73bc7e9705ae.png)

