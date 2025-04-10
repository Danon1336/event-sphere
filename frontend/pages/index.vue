<!-- pages/index.vue -->
<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />

    <!-- Уведомление о добавлении в избранное -->
    <div
      v-if="showSuccess"
      :class="['fixed top-[64px] right-6 text-[#FFFFFF] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-[150] transition-colors', 
               isDarkMode ? 'bg-[#5A52D9]' : 'bg-[#6C63FF]']"
    >
      ✅ Добавлено в избранное!
    </div>

    <!-- Уведомление об удалении из избранного -->
    <div
      v-if="showRemoveSuccess"
      :class="['fixed top-[64px] right-6 text-[#FFFFFF] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-[150] transition-colors',
               isDarkMode ? 'bg-[#E05574]' : 'bg-[#FF6584]']"
    >
      ❌ Удалено из избранного!
    </div>

    <!-- Фон с падающими звездами -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
      <div v-for="i in 30" :key="i" class="falling-star" :style="randomFallingStarStyle()" />
    </div>

    <div class="container mx-auto p-6 space-y-8 relative z-10 pt-20">
      <!-- Заголовок и краткое описание -->
      <div class="text-center space-y-2">
        <h1 class="text-3xl font-bold text-accent">Мероприятия на карте</h1>
        <p class="text-opacity-80">Найди интересные события рядом с тобой или в любой точке мира!</p>
      </div>

      <!-- Поиск и фильтры -->
      <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
        <SearchBar v-model="searchQuery" class="input-field flex-1" placeholder="Поиск мероприятий..." />
        <div class="flex gap-2">
          <select v-model="filterCategory" class="input-field p-2 rounded-xl">
            <option value="">Все категории</option>
            <option value="tech">Технологии</option>
            <option value="music">Музыка</option>
            <option value="art">Искусство</option>
          </select>
          <button class="action-btn p-2 rounded-xl" @click="applyFilters">Применить</button>
        </div>
      </div>

      <!-- Карта -->
      <div class="rounded-2xl shadow-lg overflow-hidden h-[500px]">
        <ClientOnly>
          <Map v-if="isClient" :events="filteredEvents" :selected-event-id="selectedEventId" @update:selected="onUpdateSelected" />
          <template #fallback>
            <div class="h-full flex items-center justify-center text-opacity-80">Загрузка карты...</div>
          </template>
        </ClientOnly>
      </div>

      <!-- Список мероприятий и дополнительная информация -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="col-span-2 space-y-4 events-section p-4 rounded-2xl shadow-lg">
          <h2 class="text-xl font-semibold text-accent">События</h2>
          <div
            v-for="event in filteredEvents"
            :key="event.id"
            :class="['event-card p-4 rounded-xl shadow-md cursor-pointer transition-all duration-300', { 'tab-active': selectedEventId === event.id }]"
            @click="showEventDetails(event)"
          >
            <h3 class="text-lg font-semibold">{{ event.title }}</h3>
            <p class="text-opacity-80">Lat: {{ event.lat.toFixed(2) }}, Lng: {{ event.lng.toFixed(2) }}</p>
            <p class="text-sm text-opacity-60">{{ event.category || 'Без категории' }}</p>
          </div>
        </div>

        <div class="col-span-1 space-y-4">
          <div class="p-4 rounded-2xl shadow-lg events-section">
            <h3 class="text-lg font-semibold text-accent">Популярное</h3>
            <ul class="space-y-2">
              <li v-for="(event, index) in popularEvents" :key="index" class="text-sm hover:text-accent cursor-pointer transition-colors">
                {{ event.title }}
              </li>
            </ul>
          </div>
          <div class="p-4 rounded-2xl shadow-lg events-section">
            <h3 class="text-lg font-semibold text-accent">Быстрые действия</h3>
            <div class="space-y-2">
              <button class="w-full action-btn p-2 rounded-xl flex items-center justify-center gap-2">
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
                Добавить событие
              </button>
              <button class="w-full cancel-btn p-2 rounded-xl flex items-center justify-center gap-2">
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8m4-8v8m4-8v8" />
                </svg>
                Поделиться картой
              </button>
            </div>
          </div>
        </div>
      </div>

      <EventDetails
        v-if="selectedEvent"
        :event="selectedEvent"
        :is-dark-mode="isDarkMode"
        @close="closeEventDetails"
        @toggleFavorite="handleToggleFavorite"
      />
    </div>

    <button @click="toggleTheme" class="theme-toggle fixed bottom-4 right-4 p-3 rounded-full shadow-lg z-[160]">
      <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Navbar from '~/components/Navbar.vue'
import SearchBar from '~/components/SearchBar.vue'
import Map from '~/components/Map.vue'
import EventDetails from '~/components/EventDetails.vue'

const isDarkMode = ref(true)
const events = ref([])
const selectedEventId = ref(null)
const searchQuery = ref('')
const filterCategory = ref('')
const selectedEvent = ref(null)
const isClient = ref(false)
const showSuccess = ref(false)
const showRemoveSuccess = ref(false)

onMounted(() => {
  isClient.value = true
  generateRandomEvents()
})

const generateRandomEvents = () => {
  const categories = ['tech', 'music', 'art']
  const newEvents = []
  for (let i = 0; i < 10; i++) {
    const lat = Math.random() * (82.0586232 - 41.1850968) + 41.1850968
    const lng = Math.random() * (180 - 19.6389) + 19.6389
    newEvents.push({
      id: i,
      title: `Event ${i + 1}`,
      lat,
      lng,
      category: categories[Math.floor(Math.random() * categories.length)],
      description: `Описание мероприятия ${i + 1}. Это тестовое событие с случайными данными.`
    })
  }
  events.value = newEvents
}

const popularEvents = computed(() => events.value.slice(0, 3))

const onUpdateSelected = (id) => {
  selectedEventId.value = id
}

const filteredEvents = computed(() => {
  let filtered = events.value
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(event => event.title.toLowerCase().includes(query))
  }
  if (filterCategory.value) {
    filtered = filtered.filter(event => event.category === filterCategory.value)
  }
  return filtered
})

const randomFallingStarStyle = () => {
  const size = Math.floor(Math.random() * 4) + 1
  const left = Math.random() * 100
  const top = Math.random() * -100
  const duration = 5 + Math.random() * 5
  return {
    width: `${size}px`,
    height: `${size}px`,
    left: `${left}%`,
    top: `${top}%`,
    animationDuration: `${duration}s`,
    animationDelay: `${Math.random() * 3}s`
  }
}

const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
}

const applyFilters = () => {
  console.log('Фильтры применены:', { searchQuery: searchQuery.value, filterCategory: filterCategory.value })
}

const showEventDetails = (event) => {
  console.log('Открытие события:', event)
  selectedEvent.value = event
}

const closeEventDetails = () => {
  console.log('Закрытие деталей события')
  selectedEvent.value = null
}

const handleToggleFavorite = ({ favorites, isAdded }) => {
  console.log('Обновление избранного:', favorites)
  if (isAdded) {
    showSuccess.value = true
    setTimeout(() => showSuccess.value = false, 3000)
  } else {
    showRemoveSuccess.value = true
    setTimeout(() => showRemoveSuccess.value = false, 3000)
  }
  playSuccessSound()
}

const playSuccessSound = () => {
  const audio = new Audio('/sounds/fav.mp3')
  audio.play().catch(error => console.error('Ошибка воспроизведения звука:', error))
}
</script>

<style scoped>
/* Светлая тема */
.light-theme {
  background-color: #D5D5E0;
  color: #1A1A2E;
}

.light-theme .events-section {
  background-color: #E4E4F0;
}

.light-theme .event-card {
  background-color: #F0F0F5;
}

.light-theme .input-field {
  background-color: #F0F0F5;
  border: 1px solid #E4E4F0;
  color: #1A1A2E;
}

.light-theme .theme-toggle,
.light-theme .action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.light-theme .cancel-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

.light-theme .tab-active {
  background-color: #6C63FF;
  color: #FFFFFF;
}

/* Темная тема */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
}

.dark-theme .input-field {
  background-color: #3A3A50;
  border: 1px solid #E4E4F0;
  color: #E4E4F0;
}

.dark-theme .theme-toggle,
.dark-theme .action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .cancel-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

.dark-theme .tab-active {
  background-color: #6C63FF;
  color: #FFFFFF;
}

/* Падающие звезды */
.falling-star {
  position: absolute;
  background-color: #6C63FF;
  border-radius: 50%;
  animation: fall linear infinite;
  filter: blur(1px);
}

@keyframes fall {
  0% {
    top: -10%;
    opacity: 1;
  }
  100% {
    top: 110%;
    opacity: 0;
  }
}

/* Общие стили */
.input-field {
  padding: 12px 16px;
  border-radius: 12px;
  width: 100%;
  transition: all 0.3s ease;
}

.input-field:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(108, 99, 255, 0.3);
}

.event-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.action-btn, .cancel-btn {
  transition: all 0.3s ease;
}

.action-btn:hover, .cancel-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.theme-toggle {
  transition: all 0.3s ease;
}

.theme-toggle:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.text-opacity-80 {
  opacity: 0.8;
}

.text-opacity-60 {
  opacity: 0.6;
}

.text-accent {
  color: #6C63FF;
}

.animate-fade-slide {
  animation: fadeSlide 0.5s ease-out;
}

@keyframes fadeSlide {
  0% { opacity: 0; transform: translateY(-20px); }
  100% { opacity: 1; transform: translateY(0); }
}
</style>