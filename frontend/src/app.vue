<script lang=ts setup>
import { reactive } from 'vue'
import axios from 'axios'

interface Message {
    message: string;
}

interface Data {
    message: Message | null,
}
const data: Data = reactive({
    message: null,
})

async function fetch() {
    try {
        const response = await axios.get<Message>('/api/message')
        console.log(response)
        data.message = response.data
    } catch (err) {
        console.log(err)
    }
}

fetch()

</script>

<template>
    <h1 v-if="data.message">Message: {{ data.message.message }}</h1>
</template>

