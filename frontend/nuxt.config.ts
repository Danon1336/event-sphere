export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@nuxtjs/leaflet"],
  css: ["~/assets/css/tailwind.css"],
  plugins: [{ src: "~/plugins/leaflet.client.js", mode: "client" }],
  ssr: false,
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080",
    },
  },
});
