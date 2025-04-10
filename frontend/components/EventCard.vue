<template>
    <div class="bg-black bg-opacity-70 p-4 rounded-lg shadow-lg hover:shadow-purple transition">
        <h3 class="text-lg font-orbitron">{{ event.title }}</h3>
        <p class="text-sm">{{ event.date }}</p>
        <div class="mt-2 flex gap-2">
        <span v-for="tag in event.tags" :key="tag" class="px-2 py-1 bg-purple bg-opacity-50 rounded-full text-xs">
            {{ tag }}
        </span>
        </div>
    </div>
</template>

<script setup>
defineProps({
    event: { type: Object, required: true }
});
</script>