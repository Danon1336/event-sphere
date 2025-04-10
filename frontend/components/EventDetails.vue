<!-- components/EventDetails.vue -->
<template>
    <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-[1000]" @click.self="$emit('close')">
      <div :class="['p-6 rounded-2xl shadow-lg max-w-md w-full', isDarkMode ? 'dark-theme-bg' : 'light-theme-bg']">
        <h2 v-if="event" class="text-2xl font-semibold text-accent mb-4">{{ event.title }}</h2>
        <p v-else class="text-red-500">Событие не выбрано</p>
        <p v-if="event" class="text-opacity-80 mb-2">Координаты: Lat {{ event.lat.toFixed(2) }}, Lng: {{ event.lng.toFixed(2) }}</p>
        <p v-if="event" class="text-opacity-60 mb-2">Категория: {{ event.category || 'Не указана' }}</p>
        <p v-if="event" class="text-opacity-80 mb-4">Описание: {{ event.description || 'Описание отсутствует' }}</p>
        
        <button
          v-if="event"
          :class="['p-3 rounded-xl w-full flex items-center justify-center gap-2', isFavorite ? 'cancel-btn' : 'action-btn']"
          @click="toggleFavorite"
        >
          <svg v-if="isFavorite" class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24">
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
          </svg>
          <svg v-else class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
          </svg>
          {{ isFavorite ? 'Удалить из избранного' : 'Добавить в избранное' }}
        </button>
        
        <button class="mt-4 p-3 rounded-xl w-full cancel-btn" @click="$emit('close')">Закрыть</button>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const props = defineProps({
    event: Object,
    isDarkMode: Boolean
  })
  
  const emit = defineEmits(['close', 'toggleFavorite'])
  
  const favorites = ref(JSON.parse(localStorage.getItem('favorites')) || [])
  const isFavorite = ref(false)
  
  watch(() => props.event, (newEvent) => {
    console.log('EventDetails получил событие:', newEvent)
    if (newEvent) {
      isFavorite.value = favorites.value.some(fav => fav.id === newEvent.id)
    }
  }, { immediate: true })
  
  const toggleFavorite = () => {
    if (!props.event) return
    const wasFavorite = isFavorite.value
    if (wasFavorite) {
      favorites.value = favorites.value.filter(fav => fav.id !== props.event.id)
    } else {
      favorites.value.push(props.event)
    }
    isFavorite.value = !isFavorite.value
    localStorage.setItem('favorites', JSON.stringify(favorites.value))
    emit('toggleFavorite', { favorites: favorites.value, isAdded: !wasFavorite })
  }
  </script>
  
  <style scoped>
  .dark-theme-bg {
    background-color: #2A2A40;
    color: #E4E4F0;
  }
  
  .light-theme-bg {
    background-color: #E4E4F0;
    color: #1A1A2E;
  }
  
  .action-btn {
    background-color: #6C63FF;
    color: #FFFFFF;
    transition: all 0.3s ease;
  }
  
  .cancel-btn {
    background-color: #3A3A50;
    color: #E4E4F0;
    transition: all 0.3s ease;
  }
  
  .light-theme-bg .cancel-btn {
    background-color: #F0F0F5;
    color: #1A1A2E;
  }
  
  .action-btn:hover, .cancel-btn:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }
  
  .text-accent {
    color: #6C63FF;
  }
  
  .text-opacity-80 {
    opacity: 0.8;
  }
  
  .text-opacity-60 {
    opacity: 0.6;
  }
  </style>