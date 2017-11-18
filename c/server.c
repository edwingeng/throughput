#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <time.h>

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
    const int BUF_SIZE = 1024 * 16;
    int sock, conn, c, read_size;
    struct sockaddr_in server, client;
    unsigned char buf[BUF_SIZE];
    int offset = 0, off;
    int opt;
    unsigned int sz;
    clock_t start;
    float dt;
    int throughput;
    int n = 0;
    void* ptr;

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        perror("failed to create socket\n");
        return 1;
    }

    opt = 1;
    if (setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(int)) < 0) {
        perror("failed to set SO_REUSEADDR");
    }

    server.sin_family = AF_INET;
    server.sin_addr.s_addr = INADDR_ANY;
    server.sin_port = htons(8888);

    if (bind(sock, (struct sockaddr *)&server, sizeof(server)) < 0) {
        perror("bind failed\n");
        return 1;
    }

    listen(sock, 3);
    printNow();
    printf("listening on :8888\n");

    c = sizeof(struct sockaddr_in);
    conn = accept(sock, (struct sockaddr *)&client, (socklen_t*)&c);
    if (conn < 0) {
        perror("accept failed\n");
        return 1;
    }

    while ((read_size = recv(conn, buf + offset, BUF_SIZE - offset, 0)) > 0) {
        offset += read_size;
        off = 0;
        while (offset - off >= 4) {
            sz = (buf[off + 0] << 24) | (buf[off + 1] << 16) | (buf[off + 2] << 8) | buf[off + 3];
            if (sz > BUF_SIZE) {
                perror("payload is too big\n");
                return 1;
            }
            if (off + 4 + sz > offset) {
                break;
            }

            n++;
            ptr = malloc(sz);
            memcpy(ptr, buf + off + 4, sz);
            free(ptr);
            off += 4 + sz;
        }
        if (off < offset && off > 0) {
            memcpy(buf + off, buf, offset - off);
        }
        if (off > 0) {
            offset -= off;
        }
    }

	dt = (float)(clock() - start) / 1000000;
	throughput = (int)((float)n / dt);
	printNow();
	printf("n: %d, time: %.3f, throughput: %d\n", n, dt, throughput);
    fflush(stdout);
    return 0;
}
