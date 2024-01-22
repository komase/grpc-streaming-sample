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
      count: 1,
    };
    // If count is greater than or equal to 2, error occurs here
    const res = await api.messengerSayHello(params);
    console.log(res.result?.message);
  } catch (error) {
    console.error(error);
  }
  return;
}

main();
