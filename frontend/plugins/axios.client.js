import axios from "axios";

export default defineNuxtPlugin((nuxtApp) => {
  const instance = axios.create({
    baseURL: "http://localhost:8080", // Ваш базовый URL
  });

  return {
    provide: {
      axios: instance,
    },
  };
});
