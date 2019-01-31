#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <time.h>
#include <netinet/tcp.h>

void printNow() {
    time_t timer;
    char buf[26];
    struct tm* tm_info;

    time(&timer);
    tm_info = localtime(&timer);

    strftime(buf, 26, "%Y/%m/%d %H:%M:%S", tm_info);
    printf("%s ", buf);
}

int main(int argc, char *argv[]) {
    const int MSG_SIZE = 128;
    const int N = 1000000;

    int sock;
    struct sockaddr_in server;
    unsigned char buf[4096];
    int opt;
    int i;
    int offset;
    int sent;

    for (i = 0; i < MSG_SIZE; i++) {
        buf[i + 4] = 'A';
    }

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        perror("failed to create socket\n");
        return 1;
    }

    server.sin_addr.s_addr = inet_addr("127.0.0.1");
    server.sin_family = AF_INET;
    server.sin_port = htons(8888);

    if (connect(sock, (struct sockaddr *)&server, sizeof(server)) < 0) {
        perror("failed to connect\n");
        return 1;
    }

    opt = 0;
    for (i = 0; i < argc; i++) {
        if (strcmp(argv[i], "-nodelay") == 0) {
            opt = 1;
        }
    }
    if (setsockopt(sock, IPPROTO_TCP, TCP_NODELAY, (void*)&opt, sizeof(opt)) < 0) {
        perror("failed to set TCP_NODELAY\n");
    }

    printNow();
    char* nodelayState = (opt == 1) ? ", TCP_NODELAY" : "";
    printf("payload size: %d, n: %d%s\n", MSG_SIZE, N, nodelayState);

    for (i = 0; i < N; i++) {
        buf[0] = (MSG_SIZE >> 24) & 0xFF;
        buf[1] = (MSG_SIZE >> 16) & 0xFF;
        buf[2] = (MSG_SIZE >> 8) & 0xFF;
        buf[3] = (MSG_SIZE) & 0xFF;

        offset = 0;
        while (offset < 4) {
            sent = send(sock, buf, 4 - offset, 0);
            if (sent < 0) {
                perror("failed to send [1]\n");
                return 1;
            }
            offset += sent;
        }

        while (offset < MSG_SIZE + 4) {
            sent = send(sock, buf + offset, MSG_SIZE + 4 - offset, 0);
            if (sent < 0) {
                perror("failed to send [2]\n");
                return 1;
            }
            offset += sent;
        }
    }

    shutdown(sock, 0);
    return 0;
}
