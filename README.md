# mirinae-services-cli

πΊ κ°λ° νκ²½μμ λ°°ν¬μλ²λ‘ λ°°ν¬νμ μ, λ‘κΉ λ° μλΉμ€ μν μ‘°νμ λΆνΈμ¬ν­μ λλΌμλ κ²μ κ³ λ €νμ¬ λ§λ  CLI μλλ€. πΊ
   1. κΈ°λ³Έ λͺ¨λ νμ λ° AWS Jenkins EC2 μμ μ»€λ₯μ μν νμΈ
   2. MobaXterm μ μ»€λ₯μ μ νμΌλ‘ λΆνΈν¨μΌλ‘ μΈν ν°λλ§ μ§μ λ° ν°λλ§ μν νμΈ
   3. Kubernetes ν΄λ¬μ€ν° λ΄λΆ μλΉμ€ λ° μμ μν νμΈ [λ°°ν¬ ν νμΈ μ©μ΄μ±]
   4. μΏ λ²λ€ν°μ€ μμμ λν λ‘κ·Έ μ§μ [λ°°ν¬ ν λλ²κΉ μ©μ΄μ±]
   
   ν΄λΉ CLI λ go μΈμ΄λ‘ λ§λ€μμ΅λλ€
   
   <br>
   
   μ¬μ©λ² : https://github.com/byulsoft/mirinae-services-cli/releases

<br><br>

## μ€λΉμ¬ν­
mirinae-services-cli λ `ssh` κ° μμ΄μΌ λμν©λλ€. mac νκ²½μμλ `ssh`, `ssh-client` λͺ¨λμ΄ μ λΆ μ‘΄μ¬νκ³  μμ κ²μλλ€.

<br>

- μλμ° μ΄μμ²΄μ μμ SSH μ SSH Client λ₯Ό λ°λ λ²μ λ€μκ³Ό κ°μ΅λλ€.
   - https://forbes.tistory.com/910
- μλμ° νκ²½μμλ Git Bash λ₯Ό κΈ°μ€μΌλ‘ λͺλ Ήμ΄λ₯Ό κ΅¬μ± λ° νμ±νμμΌλ―λ‘ Git Bash λ₯Ό μ¬μ©ν΄μΌ λμν©λλ€.
   - https://git-scm.com/downloads

<br><br>


## μ€μΉ

### mirinae-services-cli for window
---
1. https://github.com/byulsoft/mirinae-services-cli/releases μ μ μνμ¬ .zip νμΌμ λ΄λ €λ°κ³  μμΆμ ν΄μ ν©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132282632-e6517b74-0b89-4020-bf8d-bc52d4c59a71.png)

2. `mirinae-services-cli` λ λ΄λΆμ μΌλ‘ ν€λ₯Ό μ°Ύλ λ°μ `MIRINAE_PATH` λ₯Ό μ¬μ©ν©λλ€. (EC2 ν€λ λ³΄μ μ νλ‘μΈμ€μ λ£μ§ μμμ΅λλ€.) νκ²½λ³μλ₯Ό νΈμ§ν©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132282937-568c81f7-bdc2-4d3e-92b3-70df621a6900.png)

3. νλ‘κ·Έλ¨μμ git bash μ λ§κ² κ²½λ‘λ₯Ό νμ±νλ―λ‘, μλμ° κΈ°μ€μΌλ‘ `jenkins μλ²μ .pem ν€` νμΌμ λμ ν΄λλ₯Ό `MIRINAE_PATH` λ‘ μ€μ νλ©΄ λ©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132283017-b0d57f54-2834-4ef3-a063-0abd57ada1e8.png)

4. `git bash` λ₯Ό μ΄μ΄ ν΄λΉ λͺλ Ήμ μ€νν΄λ΄λλ€.

```bash
printenv | grep MIRINAE
## MIRINAE_PATH=C:\Users\JungSooMin\Documents\AWS\byulsoft-pemkey
```

5. `MIRINAE_PATH` μ κ°μ `jenkins μλ²μ .pem ν€` λ₯Ό λκ³  μ΄λ¦μ `byul-jenkins.pem` λ‘ λ³κ²½ ν ν, νμμ bin ν΄λλ₯Ό μμ±ν©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132283250-09b5175c-495c-4942-90ac-3bc0a6495e12.png)

6. μμ€ν νκ²½λ³μ μ€ `PATH` λ³μμ `%MIRINAE_PATH%\bin\` λ₯Ό μΆκ°ν©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132283463-f4568741-80ff-4bd7-af58-b2389de46809.png)

7. `mirinae-services.zip` νμΌμ `%MIRINAE_PATH%\bin\` μ ν΄λΉνλ κ²½λ‘λ‘ μ?κΈ°κ³  μμΆμ ν΄μ ν©λλ€.

![image](https://user-images.githubusercontent.com/67881815/132283729-4f7befeb-d36d-4c03-94d0-73cbd4c45e98.png)

8. `git bash` λ₯Ό μ΄μ΄ ν΄λΉ λͺλ Ήμ μ€ν ν΄λ΄λλ€.
```bash
mirinae-services -h
```
![image](https://user-images.githubusercontent.com/67881815/132283854-b3024020-c07c-455c-9848-8adee682580f.png)

9. `MIRINAE_PATH` κ° μ¬λ°λ₯΄κ² λ±λ‘λμ΄ μκ³ , ν΄λΉ κ²½λ‘μ `byul-jenkins.pem` ν€κ° μλ€λ©΄, ν΄λΉ λͺλ Ήμ μ±κ³΅ν©λλ€. `KeyPath` κ° νμ±λ `MIRINAE_PATH` μ `byul-jenkins.pem` λ₯Ό λΆμΈ κ²½λ‘μλλ€.
```bash
mirinae-services require
```
![image](https://user-images.githubusercontent.com/67881815/132284022-ab24be4f-b4fc-450c-bc92-60b86c8cda34.png)


<br><br>


### mirinae-services-cli for mac, linux
---
1. go λ‘ λΉλ λ νμΌμΈ mirinae-services.exe λ₯Ό λ°μ΅λλ€. (ν¬κΈ° μ½ 8M) => https://github.com/soominJung0413/mirinae-services-cli/releases
```bash
## .tar νμΌ λ€μ΄λ‘λ
curl -LO https://github.com/byulsoft/mirinae-services-cli/releases/download/mirinae/mirinae-services.tar

tar -xvf mirinae-services.tar
```
3. `mirinae-services-cli` λ λ΄λΆμ μΌλ‘ ν€λ₯Ό μ°Ύλ λ°μ `MIRINAE_PATH` λ₯Ό μ¬μ©ν©λλ€. (EC2 ν€λ λ³΄μ μ νλ‘μΈμ€μ λ£μ§ μμμ΅λλ€.) λ€μμ μ»€λ©λλ₯Ό μ€νν©λλ€.
```bash
### zsh λ₯Ό μ¬μ©νλ€λ©΄ .zshrc λ₯Ό μ½λλ€.
vi ~/.bash_profile

### ~/.bash_profile
export MIRINAE_PATH=[μνλ ν΄λ κ²½λ‘]
```
3. μμ€ νμΌμ μμ μ μ©νκ³  ν°λ―Έλμ λ€μ ν€κ³  `MIRINAE_PATH` μ μ© μ¬λΆλ₯Ό νμΈν©λλ€.
```bash
### zsh λ₯Ό μ¬μ©νλ€λ©΄ .zshrc λ₯Ό μ λ¬ν©λλ€.
source ~/.bash_profile

printenv | grep MIRINAE_PATH
## MIRINAE_PATH=/Users/jeongsumin/Documents/AWS/byul/byulsoft-pemkey/
```
4. `/usr/local/bin` λ‘ νμΌμ μ?κΈ°κ³  ν°λ―Έλμ λ€μ μ΄μ΄ λ±λ‘μ΄ λ¬λμ§ νμΈν©λλ€.
```bash
sudo mv mirinae-services /usr/local/bin/

mirinae-services -h
```
![image](https://user-images.githubusercontent.com/67881815/132218847-41a436b2-8981-4491-b815-a7b13bfc9827.png)

5. EC2 λ³΄μ μ μ±μ λ°λΌμ ν€ νμΌμ mod λ₯Ό 400 μΌλ‘ λ³κ²½νκ³ , MIRINAE_PATH μ λ³΅μ¬ν©λλ€
```bash
chmod 400 [Jenkins μλ²].pem && cp [Jenkins μλ²].pem ${MIRINAE_PATH}/

ls -al ${MIRINAE_PATH}
```
6. ν°λ―Έλμ λ€μ μ΄κ³  μ€ν ν΄λ΄λλ€
```bash
mirinae-services require
```
7. ν΄λΉ λ¬Έκ΅¬κ° μΆλ ₯λλ€λ©΄, μ€μΉ λ° EC2 ν€  μ μμ μΌλ‘ μλ£ λ κ² μλλ€.
![image](https://user-images.githubusercontent.com/67881815/132219011-8a1a557b-47f4-495d-b667-73bc7e9705ae.png)

