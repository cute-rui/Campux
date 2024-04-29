/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

import axios from 'axios'
import store from './store/index.js'

const app = createApp(App)
app.use(store)

const axiosInstance = axios.create({
    withCredentials: true,
    baseURL: store.state.base_url,
})
app.config.globalProperties.$axios = { ...axiosInstance }

// let config=require("../config.json");
fetch('/config.json').then(response => response.json()).then(config => {
    console.log(config)
    store.commit('setBaseURL', config.base_url)
})

registerPlugins(app)

app.mount('#app')

