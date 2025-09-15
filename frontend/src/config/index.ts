import {createRouter, createWebHashHistory} from 'vue-router'
import Connection from './../components/connection/Connection.vue'
import Log from "../components/log/Log.vue";
import Config from "../components/config/Config.vue"
import Info from "../components/info/Info.vue";


const routes = [
    {path: '/', component: Connection},
    {path: '/log', component: Log},
    {path: '/config', component: Config},
    {path: '/info', component: Info},
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router