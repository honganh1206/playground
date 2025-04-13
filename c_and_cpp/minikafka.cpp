#include "minikafka.h"
#include <fstream>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>

// Write messages to the message broker
// All messages are stored in a file
void producer(const char *filePath) {
  std::fstream os(filePath, ios_base::app | ios_base::out);
  int count = 0;
  // Disable buffering for the file stream
  // So writing can happen immediately without waiting to fill a buffer
  // Here we write to page cache
  os.rdbuf()->pubsetbuf(0, 0);
  while (os.is_open()) {
    // Append one char to a file every second
    os.seekp(0, os.end);
    os << char('a' + count++);
    count %= 26;
    sleep(1);
  }
}

// Send request to get messages from the file every second
void consumer (const char *filePath) {
  int fd = open(filePath, O_RDONLY);
  if (fd < 0){
     perror("open");
     return;
  }
  off_t offset = 0; // Set the current position (in bytes) within an open file
  int len = 64;
  char *buf = (char *)malloc(len+1);
  setvbuf(stdout, 0, _IONBF, 0); // Change the buffering mode of the given file stream
  while (fd >0) {
    if (lseek(fd, offset, SEEK_SET) == -1) { // Move the read pointer to offset
      break;
    }
    int numRead = read(fd, buf, len); // Read bytes to buffer
    if (numRead > 0) {
      buf[numRead] =0; // Add a null terminator
      printf("%s", buf);
    }
    memset((void*)buf, 0, len+1); // Clear the buffer to remove old data
    offset += numRead; // Advance the offset
    sleep(1);
  }
  free(buf);
}

int main(int argc, char *argv[]) {
  const char* filePath = "/tmp/deadbeef.log";
  auto f1 = async(launch::async, producer, filePath); // Run the producer async
  auto f2 = async(launch::async, consumer, filePath);
  f2.get(); // Get one message per second
  return 0;
}
