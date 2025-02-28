// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,

  typescript: {
    shim: false
  },

  build: {
    transpile: ["vuetify"],
  },

  vite: {
    define: {
      "process.env.DEBUG": false,
    },
  },

  nitro: {
    serveStatic: true,
  },

  devServerHandlers: [],

  hooks: {
  },

  devtools: { enabled: true },

  modules: [
    "@sidebase/nuxt-auth",
  ],

  auth: {
    baseURL: process.env.BACKEND_BASEURL,
    provider: {
      type: 'local',
      token: {
        type: "Bearer",
        signInResponseTokenPointer: "/data/token",
        maxAgeInSeconds: 3600,
        cookieName: "auth.token",
        headerName: "Authorization",
        cookieDomain: 'localhost',
        sameSiteAttribute: 'lax',  
        secureCookieAttribute: false,  
        httpOnlyCookieAttribute: false
      },
      endpoints: {
        signIn: { path: '/api/v1/auth/login', method: 'post' },
        signOut: { path: '/api/v1/auth/logout', method: 'post' },
        signUp: { path: '/api/v1/auth/register', method: 'post' },
        getSession: { path: '/api/v1/auth/session', method: 'get' },
      }
    },
    globalAppMiddleware: {
      isEnabled: true
    }
  },

  compatibilityDate: '2025-02-26'
})