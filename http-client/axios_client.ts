import { MessengerApi } from "./openapi/axios";

async function main() {
  try {
    const api = new MessengerApi(undefined, "http://localhost:8080");
    const res = await api.messengerSayHello("test", 3);
    console.log(res.data.result?.message);
    console.log(res.data);
  } catch (error) {
    console.error(error);
  }
  return;
}

main();
