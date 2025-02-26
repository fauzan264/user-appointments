// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    "@sidebase/nuxt-auth",
  ],
  auth: {
    baseURL: process.env.BACKEND_BASEURL,
    provider: {
      type: 'local',
      endpoints: {
        signIn: { path: '/api/v1/auth/login', method: 'post'},
        signOut: { path: '/api/v1/auth/logout', method: 'post' },
        signUp: { path: '/api/v1/auth/register', method: 'post' },
        // getSession: { path: '/api/v1/user', method: 'get' },
      }
    },
    globalAppMiddleware: {
      isEnabled: true
    }
  }
})
