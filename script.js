import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';
export const options = {
   vus: 10,
   duration: '10s',
  };
const client = new grpc.Client();
client.load(['proto'], 'simple-grpc.proto');

export default () => {
  client.connect('localhost:50051', {
    plaintext: true
  });

  const response = client.invoke('simplegrpc.Services/GetAllUsers', {});

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
  });
  //console.log(JSON.stringify(response.message));
  client.close();
  sleep(1);
};