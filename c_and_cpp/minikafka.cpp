#include <fstream>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <sys/stat.h>
#include<future>
#include <iostream>
#include <cassert>

using namespace std;
using namespace std::chrono;

// Save the time when producer sends a msg and consumer receives a msg
high_resolution_clock::time_point p[128]={}, c[128]= {};
int pcount = 0; // Num of messages sent by producer
int ccount = 0; // Num of messages received by consumer

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
    os << (char)(pcount);
    if(++pcount == 128) {
      break;
    }
  }
}

// Read one byte at a time
void consumer (const char *filePath) {
  int fd = open(filePath, O_RDONLY);
  if (fd < 0){
     perror("open");
     return;
  }
  off_t offset = 0; // Keep track of how many bytes have been read
  int len = 1, numRead = 1; // Specify reading 1 byte at a time and storing num of bytes ACTUALLY read
  char *buf = (char *)malloc(len+1); // 2-byte buffer, one for the byte, one for the null terminator
  setvbuf(stdout, 0, _IONBF, 0); // Disable output buffering on stdout so any printf() will flush immediately
  while (fd >0) {
    if (lseek(fd, offset, SEEK_SET) == -1) continue; // Move the read pointer to offset
    int numRead = 0;
    // Keep polling the file until something new is appended
    // This eats CPU cycles if no new data is written
    // TODO: use inotify
    while (numRead == 0)
      numRead = read(fd, buf, len);
    if (numRead == 1) {
      ccount++; // Tracking num of chars read
      c[(int)buf[0]] = high_resolution_clock::now(); // Store the timestamp for the current character
      // Specific byte-based break condition
      if((int)(buf[0]) == 127) break;
    } else { break; }
    offset += numRead;
  }
  free(buf);
}

int main(int argc, char *argv[]) {
  const char* filePath = "/tmp/deadbeef.log";
  auto f1 = async(launch::async, producer, filePath); // Run the producer async
  auto f2 = async(launch::async, consumer, filePath);
  f1.get();
  f2.get(); // Get one message per second
  // Calculate the time between sending and receiving the message
  int sum = 0;
  for (int i = 0; i < 128; i++) {
    auto int_ns = chrono::duration_cast<chrono::nanoseconds>(c[i] - p[i]);
    cout << i << ": " << int_ns.count() << "ns" << endl;
    sum +=int_ns.count();
  }
  cout << "average delay: " << sum/128 << "ns" << endl;
  assert(pcount==ccount and pcount == 128);
  return 0;
}
