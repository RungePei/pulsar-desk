import {defineConfig} from 'vite'
import Vue from '@vitejs/plugin-vue'
import Inspect from 'vite-plugin-inspect'


export default defineConfig({
    resolve: {
        alias: {
            '@': "src",
        },
    },
    plugins: [
        Vue(),
        Inspect(),
    ],
})