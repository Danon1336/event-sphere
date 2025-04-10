<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />
    
    <div class="container mx-auto p-4" style="padding-top: 80px">
      <!-- Organizer Profile Header -->
      <div class="flex flex-col md:flex-row gap-6 mb-6">
        <!-- Profile Card -->
        <div class="p-6 rounded-2xl user-info shadow-lg flex-1 flex items-center gap-6">
          <div class="relative">
            <div class="h-20 w-20 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
              <span class="text-2xl font-bold text-white">ТС</span>
            </div>
            <div class="absolute -bottom-1 -right-1 h-6 w-6 rounded-full bg-green-500 border-2 border-white"></div>
          </div>
          <div class="flex-1">
            <h2 class="text-2xl font-bold">ООО "ТехноСфера"</h2>
            <p class="text-opacity-80 mb-2">ИТ инновации и разработка</p>
            <div class="flex gap-2">
              <span class="px-3 py-1 rounded-full text-sm bg-opacity极客时间20 bg-blue-500">Verified Organizer</span>
              <span class="px-3 py-1 rounded-full text-sm bg-opacity-20 bg-green-500">Active</span>
            </div>
          </div>
        </div>

        <!-- Stats Cards -->
        <div class="grid grid极客时间ols-2 md:grid-cols-4 gap-3 flex-1">
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Предстоящие</p>
            <p class="text-2xl font-bold">5</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Текущие</p>
            <p class="text-2xl font-bold">2</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-green-500 to-teal-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Завершённые</p>
            <p class="text-2xl font-bold">12</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-yellow-500 to-orange-500 rounded-full"></div>
          </div>
          <div class="p-4 rounded-xl stat-card shadow-md">
            <p class="text-opacity-60 text-sm">Участников</p>
            <p class="text-2xl font-bold">156</p>
            <div class="h-1 mt-2 bg-gradient-to-r from-pink-500 to-red-500 rounded-full"></div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Events Section -->
        <div class="lg:col-span-2">
          <div class="p-6 rounded-2xl events-section shadow-lg">
            <div class="极客时间flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold">Мои мероприятия</h2>
              <div class="relative">
                <input 
                  v-model="searchQuery"
                  type="text" 
                  placeholder="Поиск мероприятий..."
                  class="p-2 pl-10 rounded-lg bg-opacity-20 focus:bg-opacity-30 transition-colors"
                  :class="isDarkMode ? 'bg-white/10 text-white' : 'bg-black/10 text-black'"
                >
                <svg class="h-5 w-5 absolute left-3 top-2.5 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 极客时间0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
            </div>
            
            <!-- Event Filters -->
            <div class="mb-6 flex flex-wrap gap-3">
              <button 
                v-for="(tab, index) in tabs" 
                :key="index"
                @click="activeTab = tab.id"
                :class="['px-4 py-2 rounded-lg transition-colors', activeTab === tab.id ? 'bg-blue-500 text-white' : 'bg-opacity-20 hover:bg-opacity-30']"
              >
                {{ tab.label }}
              </button>
            </div>

            <!-- Events List -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div 
                v-for="event in filteredEvents" 
                :key="event.id"
                class="p-4 rounded-xl event-card shadow-md cursor-pointer hover:shadow-lg transition-all duration-300"
                @click="goToEvent(event.id)"
              >
                <div class="flex justify-between items-start mb-2">
                  <h3 class="text-lg font-bold">{{ event.title }}</h3>
                  <span class="px-2 py-1 rounded-full text-xs" 
                    :class="{
                      'bg-blue-100 text-blue-800': event.status === 'upcoming',
                      'bg-green-100 text-green-800': event.status === 'current',
                      'bg-gray-100 text-gray-800': event.status === 'past'
                    }"
                  >
                    {{ event.status === 'upcoming' ? 'Предстоящее' : 
                       event.status === 'current' ? 'Текущее' : 'Завершённое' }}
                  </span>
                </div>
                <p class="text-sm text-opacity-80 mb-2">{{ formatDate(event.date) }}</p>
                <p class="text-sm text-opacity-60">{{ event.description ? event.description.substring(0, 100) + '...' : '' }}</p>
                <div class="mt-3 flex flex-wrap gap-2">
                  <span 
                    v-for="(tag, index) in event.tags.split(',')" 
                    :key="index"
                    class="px-2 py-1 rounded-full text-xs bg-opacity-20"
                    :class="isDarkMode ? 'bg-white' : 'bg-black'"
                  >
                    {{ tag.trim() }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Column: Add New Event -->
        <div class="col-span-1 p-4 rounded-2xl add-event shadow-lg">
        <h2 class="text-xl font-semibold mb-4">Добавить мероприятие</h2>
        <form @submit.prevent="addEvent">
          <div class="mb-4 relative">
            <label class="block mb极客时间1 text-opacity-80">Название</label>
            <input v-model="newEvent.title" type="text" class="w-full p-3 pl-10 rounded-xl input-field" placeholder="Название" />
            <svg class="h-5 w-5 absolute left-3 top-10 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
            </svg>
          </div>
          <div class="mb-4 relative">
            <label class="block mb-1 text-opacity-80">Дата (ДД.ММ.ГГГГ)</label>
            <input v-model="newEvent.date" type="text" class="w-full p-3 pl-10 rounded-xl input-field" placeholder="ДД.ММ.ГГГГ" />
            <svg class="h-5 w-5 absolute left-3 top-10 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill极客时间="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2极客时间v12a2 2 极客时间0 002 2z" />
            </svg>
          </div>
          <div class="mb-4 relative">
            <label class="block mb-1 text-opacity-80">Описание</label>
            <textarea v-model="newEvent.description" class="w-full p-3 pl-10 rounded-xl input-field" placeholder="Описание" rows="3"></textarea>
            <svg class="h-5 w-5 absolute left-3 top-10 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
            </svg>
          </div>
          <div class="mb-4 relative">
            <label class="block mb-1 text-opacity-80">Теги</label>
            <input v-model="newEvent.tags" type="text" class="w-full p-3 pl-10 rounded-xl input-field" placeholder="Теги через запятую" />
            
            <svg class="h-5 w-5 absolute left-3 top-10 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
            </svg>
          </div>
          <div class="mb-4">
            <label class="block mb-1 text-opacity-80">Загрузка изображения</label>
            <label class="file-input w-full p-3 rounded-xl flex items-center gap-2 cursor-pointer">
              <svg class="h-5 w-5 text-opacity-60" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="极客时间0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
              </svg>
              <span>{{ newEvent.image ? newEvent.image.name : 'Выбрать файл' }}</span>
              <input type="file" @change="onFileChange" class="hidden" />
            </label>
          </div>
          <div class="map-placeholder rounded-xl">
            <div id="add-event-map" class="h-[300px] w-full rounded-xl"></div>
          </div>

          <div class="flex gap-3">
            <button type="submit" class="p-3 rounded-xl action-btn">Добавить</button>
            <button type="button" class="p-3 rounded-xl cancel-btn">Отменить</button>
          </div>
        </form>
        </div>
      </div>
    </div>
  </div>
  <!-- Theme Toggle Button -->
  <button @click="toggleTheme" class="theme-toggle fixed bottom-4 right-4 p-3 rounded-full shadow-lg">
      <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m极客时间0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
    </button>
</template>


<script>
import { onMounted, ref } from 'vue'

export default {
  data() {
    return {
      isDarkMode: true,
      activeTab: 'all',
      filter: 'participated',
      newEvent: {
        title: '',
        date: '',
        description: '',
        tags: '',
        image: null,
      },
      events: [
        { id: 1, title: 'ИТ Конференция 2025', date: '2025-04-10', tags: 'конференция, ИТ', creator: 'ООО "ТехноСфера"', isMine: true },
        { id: 2, title: 'Научный симпозиум', date: '2025-06-15', tags: 'наука, симпозиум', creator: 'Другой организатор', isMine: false },
      ],
    };
  },
  setup() {
    const map = ref(null)
    const marker = ref(null)
    const selectedLatLng = ref({ lat: 55.7963, lng: 49.1088 }) // Казань по умолчанию

    onMounted(async () => {
      if (process.client) {
        const L = await import('leaflet')
        await import('leaflet/dist/leaflet.css')
        
        map.value = L.map('add-event-map').setView([selectedLatLng.value.lat, selectedLatLng.value.lng], 13)

        L.tileLayer('https://{s}.tile.openstreetmap.org/{极客时间z}/{x}/{y}.png', {
          attribution: '&copy; OpenStreetMap contributors'
        }).addTo(map.value)

        marker.value = L.marker([selectedLatLng.value.lat, selectedLatLng.value.lng], {
          draggable: true
        }).addTo(map.value)

        map.value.on('click', (e) => {
          const { lat, lng } = e.latlng
          selectedLatLng.value = { lat, lng }
          marker.value.setLatLng(e.latlng)
        })
      }
    })

    return {
      map,
      marker,
      selectedLatLng
    }
  },
  computed: {
    filteredEvents() {
      if (this.activeTab === 'my') {
        return this.events.filter(event => event.isMine);
      }
      return this.events;
    },
  },
  methods: {
    toggleTheme() {
      this.isDarkMode = !this.isDarkMode;
    },
    addEvent() {
      const newEvent = {
        id: this.events.length + 1,
        title: this.newEvent.title,
        date: this.newEvent.date,
        tags: this.newEvent.tags,
        creator: 'ООО "ТехноСфера"',
        isMine: true,
      };
      this.events.push(newEvent);
      this.newEvent = { title: '', date: '', description: '', tags: '', image: null };
    },
    onFileChange(event) {
      this.newEvent.image = event.target.files[0];
    },
    goToEvent(id) {
      this.$router.push(`/event/${id}`);
    },
    formatDate(dateString) {
      const options = { year: 'numeric', month: 'long', day: 'numeric' };
      return new Date(dateString).toLocaleDateString('ru-RU', options);
    }
  }
};


</script>

<style scoped>

@import "leaflet/dist/leaflet.css";
/* Светлая тема */
.light-theme {
  background-color: #D5D5E0;
  color: #1A1A2E;
}

.light-theme .user-info,
.light-theme .add-event,
.light-theme .events-section {
  background-color: #E4E4F0;
}

.light-theme .event-card {
  background-color: #F0F0F5;
}

.light-theme .nav-btn,
.light-theme .action-btn,
.light-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.light-theme .logout-btn {
  background-color: #FF6584;
  color: #FFFFFF;
}

.light-theme .cancel-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

.light-theme .input-field,
.light-theme .file-input {
  background-color: #F0F0F5;
  border: 1px solid #E4E4F0;
  color: #1A1A2E;
}

.light-theme .tab-btn {
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

.dark-theme .user-info,
.dark-theme .add-event,
.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
}

.dark-theme .nav-btn,
.dark-theme .action-btn,
.dark-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .logout-btn {
  background-color: #FF6584;
  color: #FFFFFF;
}

.dark-theme .cancel-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

.dark-theme .input-field,
.dark-theme .file-input {
  background-color: #3A3A50;
  border: 1px solid #E4E4F0;
  color: #E4E4F0;
}

.dark-theme .tab-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

.dark-theme .tab-active {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.map-placeholder {
  height: 100%;
  background-color: #ccc;
  display: flex;
  margin-bottom: 1%;
  align-items: center;
  justify-content: center;
}

.logo {
  color: #6C63FF;
}

.text-opacity-80 {
  opacity: 0.8;
}

.text-opacity-60 {
  opacity: 0.6;
}

.theme-toggle {
  transition: all 0.3s ease;
}

.theme-toggle:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}


#add-event-map {
  height: 100%;
  width: 100%;
  border-radius: 12px;
  z-index: 0;
}
.light-theme .nav-bar {
  background: rgba(228, 228, 240, 0.9);
  border-bottom: 1px solid rgba(26, 26, 46, 0.1);
}

.light-theme .nav-btn {
  background: rgba(240, 240, 245, 0.8);
  color: #1A1A2E;
}

.light-theme .nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.light-theme .logout-btn {
  background: rgba(255, 101, 132, 0.8);
}

.light-theme .logout-btn:hover {
  background: #FF6584;
}

.light-theme .logo {
  color: #6C63FF;
}

.dark-theme .nav-bar {
  background: rgba(42, 42, 64, 0.9);
}

.dark-theme .nav-btn {
  background: rgba(58, 58, 80, 0.8);
  color: #E4E4F0;
}

.dark-theme .nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .logout-btn {
  background: rgba(255, 101, 132, 0.8);
}

.dark-theme .logout-btn:hover {
  background: #FF6584;
}

.dark-theme .logo {
  color: #6C63FF;
}
</style>
