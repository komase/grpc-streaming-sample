import {
  Configuration,
  ConfigurationParameters,
  MessengerApi,
  MessengerSayHelloRequest,
} from "./openapi/fetch";

async function main() {
  try {
    const config: ConfigurationParameters = {
      basePath: "http://localhost:8080",
    };
    const api = new MessengerApi(new Configuration(config));
    const params: MessengerSayHelloRequest = {
      name: "test",
      count: 3,
    };
    const res = await api.messengerSayHello(params);
    console.log(res.result?.message);
    console.log(res);
  } catch (error) {
    console.error(error);
  }
  return;
}

main();
