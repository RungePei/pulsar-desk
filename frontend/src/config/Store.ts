import {defineStore} from "pinia";
import {ref} from 'vue';

export const idStore = defineStore('id', () => {
    const connId = ref(0)
    const topicId = ref(0)

    return {connId, topicId}
})