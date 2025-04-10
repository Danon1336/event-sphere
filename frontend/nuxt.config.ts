// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@nuxtjs/leaflet"],
  css: ["~/assets/css/tailwind.css"],
  plugins: [
    { src: "~/plugins/leaflet.client.js", mode: "client" },
    { src: "~/plugins/axios.client.js", mode: "client" },
  ],
});
