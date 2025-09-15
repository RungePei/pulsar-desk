import {Component, createApp} from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import {createPinia} from 'pinia'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './config/index.js'
import 'element-plus/theme-chalk/dark/css-vars.css'

const app = createApp(App)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component as Component)
}

app.use(ElementPlus)
    .use(createPinia())
    .use(router)
    .mount('#app')


