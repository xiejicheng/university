#include<stdio.h>
#include<winsock2.h>
#pragma comment (lib, "ws2_32.lib")  //加载ws2_32.dll

int main()
{
    //初始化 DLL
    WSADATA wsaData;
    WSAStartup(MAKEWORD(2, 2), &wsaData);

    //创建套接字
    SOCKET servSock = socket(PF_INET, SOCK_STRAM, TPPROTO_TCP);

    //绑定套接字
    sockaddr_in sockAddr;
    memset(&sockAddr, 0, sizeof(sockAddr)); //每个字节都用 0 填充
    sockAddr.sin_family = PF_INET; //使用 IPv4 地址
    sockAddr.sin_addr.s_addr = inet_addr("127.0.0.1"); //具体的 IP 地址
    sockAddr.sin_port = htons(1234); //端口
    bind(servSock, (SOCKADDR *)&sockAddr, sizeof(SOCKADDR));

    //进入监听状态
    listen(servSock, 20);

    //接收客户端请求
    SOCKADDR clntAddr;
    int nSize = sizeof(SOCKADDR);
    SOCKET clntSock = accept(servSock, (SOCKADDR *)&clntAddr, &nSize);

    //向客户端发送数据
    char *str = "Hello World!";
    send(clntSock, str, strlen(str) + sizeof(char), NULL);

    //关闭套接字
    closesocket(clntSock);
    closesocket(servSock);

    //终止 DLL 使用
    WSACleanup();

    getchar();
    return 0;
}